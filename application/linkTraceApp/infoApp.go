// @area /linkTrace/
package linkTraceApp

import (
	"fmt"
	"fops/application/linkTraceApp/response"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/trace/eumCallType"
	"github.com/farseer-go/linkTrace/driver/clickhouse"
	"github.com/farseer-go/linkTrace/eumTraceType"
	"github.com/farseer-go/mapper"
	"time"
)

// Info 链路追踪日志详情
// @get info/{traceId}
func Info(traceId int64) collections.List[response.LinkTraceVO] {
	l := linkTraceWarp{
		lst:       collections.NewList[response.LinkTraceVO](),
		rgbaIndex: -1,
		lstPO:     linkTrace_clickhouse.CHContext.TraceContext.Where("trace_id", traceId).Asc("start_ts").ToList(),
	}
	if l.lstPO.Count() == 0 {
		return l.lst
	}

	// 当A服务调用B服务时，前后均有可能包含数据库之类的操作。因此需要将lstPO重新组织。按实际的调用顺序重新排序
	// 前端就可以简单的遍历lst显示到页面即可
	entryPO := l.lstPO.First()
	l.startTs = entryPO.StartTs
	l.addEntry(entryPO)

	// 调用完成
	l.lst.Add(response.LinkTraceVO{
		Rgba: response.RgbaList[0], AppId: entryPO.AppId, AppIp: entryPO.AppIp, AppName: entryPO.AppName, UseTs: 0,
		StartTs: entryPO.EndTs - entryPO.StartTs,
		Caption: "调用完成",
		Desc:    fmt.Sprintf("耗时：%v ms", entryPO.UseTs.Milliseconds()),
	})

	return l.lst
}

type linkTraceWarp struct {
	lst       collections.List[response.LinkTraceVO]                // 要返回的集合
	rgbaIndex int                                                   // 实现不同服务的调用，用颜色区分。这里通过服务入口调用时，使用索引（对应数组颜色）
	lstPO     collections.List[linkTrace_clickhouse.TraceContextPO] // 数据库读的集合
	startTs   int64                                                 // 初始开始时间
}

// 服务调用入口
func (receiver *linkTraceWarp) addEntry(po linkTrace_clickhouse.TraceContextPO) {
	receiver.rgbaIndex++
	// 添加服务的入口
	entryTrace := response.LinkTraceVO{
		Rgba: response.RgbaList[receiver.rgbaIndex], AppId: po.AppId, AppIp: po.AppIp, AppName: po.AppName, UseTs: po.UseTs.Microseconds(), UseDesc: po.UseTs.String(),
		StartTs: (po.StartTs - receiver.startTs) / int64(time.Microsecond),
		Caption: fmt.Sprintf("%s", po.TraceType.ToString()),
	}
	switch po.TraceType {
	case eumTraceType.WebApi:
		entryTrace.Caption += fmt.Sprintf("%v %s => %s", po.WebStatusCode, po.WebMethod, po.WebPath)
		entryTrace.Desc = fmt.Sprintf("%s 客户端IP：%s", po.WebContentType, po.WebRequestIp)
	case eumTraceType.MqConsumer:
		entryTrace.Caption += fmt.Sprintf("%s %s %s", po.ConsumerServer, po.ConsumerQueueName, po.ConsumerRoutingKey)
	case eumTraceType.QueueConsumer:
		entryTrace.Caption += fmt.Sprintf("%s", po.ConsumerQueueName)
	case eumTraceType.FSchedule:
		entryTrace.Caption += fmt.Sprintf("任务组：%s(%v) 任务ID：%v", po.TaskName, po.TaskGroupId, po.TaskId)
	case eumTraceType.Task:
		entryTrace.Caption += fmt.Sprintf("%s", po.TaskName)
	case eumTraceType.WatchKey:
		entryTrace.Caption += fmt.Sprintf("%s", po.WatchKey)
	}
	receiver.lst.Add(entryTrace)
	receiver.addDetail(po)
	receiver.rgbaIndex--
}

// 服务所属的明细
func (receiver *linkTraceWarp) addDetail(po linkTrace_clickhouse.TraceContextPO) {
	for _, detail := range po.List {
		m := detail.(map[string]any)
		baseDetailPO := mapper.Single[linkTrace_clickhouse.BaseTraceDetailPO](m)
		detailTrace := response.LinkTraceVO{
			Rgba: response.RgbaList[receiver.rgbaIndex], AppId: po.AppId, AppIp: po.AppIp, AppName: po.AppName,
			StartTs: baseDetailPO.StartTs - receiver.startTs,
			UseTs:   baseDetailPO.UseTs.Microseconds(), UseDesc: baseDetailPO.UseTs.String()}

		switch eumCallType.Enum(parse.ToInt(detail.(map[string]any)["CallType"])) {
		case eumCallType.Database:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailDatabasePO](m)
			if detailPO.TableName == "" && detailPO.Sql == "" {
				detailTrace.Caption = fmt.Sprintf("打开数据库 => %s %s", detailPO.DbName, detailPO.ConnectionString)
			} else {
				detailTrace.Caption = fmt.Sprintf("执行数据库 => %s %s 影响%v行", detailPO.DbName, detailPO.TableName, detailPO.RowsAffected)
			}
			detailTrace.Desc = detailPO.Sql
			receiver.lst.Add(detailTrace)
		case eumCallType.Http:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailHttpPO](m)
			detailTrace.Caption = fmt.Sprintf("调用http => %v %s %s", detailPO.StatusCode, detailPO.Method, detailPO.Url)
			lstHeader := collections.NewList[string]()
			for k, v := range detailPO.Headers.ToMap() {
				lstHeader.Add(fmt.Sprintf("%s=%v", k, v))
			}
			detailTrace.Desc = fmt.Sprintf("头部：%s 入参：%s 出参：%s", lstHeader.ToString(","), detailPO.RequestBody, detailPO.ResponseBody)
			receiver.lst.Add(detailTrace)
			// 查找串联的服务
			nextEntry := receiver.lstPO.Where(func(item linkTrace_clickhouse.TraceContextPO) bool {
				return item.ParentAppName == detailTrace.AppName
			}).First()
			receiver.addEntry(nextEntry)
		case eumCallType.Grpc:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailGrpcPO](m)
			detailTrace.Caption = fmt.Sprintf("调用http => %v %s %s", detailPO.StatusCode, detailPO.Method, detailPO.Url)
			lstHeader := collections.NewList[string]()
			for k, v := range detailPO.Headers.ToMap() {
				lstHeader.Add(fmt.Sprintf("%s=%v", k, v))
			}
			detailTrace.Desc = fmt.Sprintf("头部：%s 入参：%s 出参：%s", lstHeader.ToString(","), detailPO.RequestBody, detailPO.ResponseBody)
			receiver.lst.Add(detailTrace)
		case eumCallType.Redis:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailRedisPO](m)
			detailTrace.Caption = fmt.Sprintf("执行Redis => %s %s", detailPO.Key, detailPO.Field)
			detailTrace.Desc = fmt.Sprintf("%s %s", detailPO.Key, detailPO.Field)
			receiver.lst.Add(detailTrace)
		case eumCallType.Mq:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailMqPO](m)
			detailTrace.Caption = fmt.Sprintf("发送MQ消息 => %s %s %s", detailPO.Server, detailPO.Exchange, detailPO.RoutingKey)
			detailTrace.Desc = fmt.Sprintf("%s %s %s", detailPO.Server, detailPO.Exchange, detailPO.RoutingKey)
			receiver.lst.Add(detailTrace)
		case eumCallType.Elasticsearch:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailEsPO](m)
			detailTrace.Caption = fmt.Sprintf("执行ES => %s %s", detailPO.IndexName, detailPO.AliasesName)
			detailTrace.Desc = fmt.Sprintf("%s %s", detailPO.IndexName, detailPO.AliasesName)
			receiver.lst.Add(detailTrace)
		case eumCallType.Etcd:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailEtcdPO](m)
			detailTrace.Caption = fmt.Sprintf("执行Etcd => %s %v", detailPO.Key, detailPO.LeaseID)
			detailTrace.Desc = fmt.Sprintf("%s %v", detailPO.Key, detailPO.LeaseID)
			receiver.lst.Add(detailTrace)
		case eumCallType.Hand:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailHandPO](m)
			detailTrace.Caption = fmt.Sprintf("手动埋点 => %s", detailPO.Name)
			detailTrace.Desc = fmt.Sprintf("%s", detailPO.Name)
			receiver.lst.Add(detailTrace)
		case eumCallType.KeyLocation:
			detailPO := mapper.Single[linkTrace_clickhouse.TraceDetailHandPO](m)
			detailTrace.Caption = fmt.Sprintf("关键埋点 => %s", detailPO.Name)
			detailTrace.Desc = fmt.Sprintf("%s", detailPO.Name)
			receiver.lst.Add(detailTrace)
		}
	}
}
