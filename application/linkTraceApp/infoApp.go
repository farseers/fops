// @area /linkTrace/
package linkTraceApp

import (
	"fmt"
	"fops/application/linkTraceApp/response"
	"fops/domain/linkTrace"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/trace"
	"github.com/farseer-go/fs/trace/eumCallType"
	linkTraceCom "github.com/farseer-go/linkTrace"
	"github.com/farseer-go/linkTrace/eumTraceType"
	"github.com/farseer-go/mapper"
	"time"
)

// Info 链路追踪日志详情
// @get info/{traceId}
func Info(traceId int64, linkTraceRepository linkTrace.Repository) collections.List[response.LinkTraceVO] {
	l := linkTraceWarp{
		lst:       collections.NewList[response.LinkTraceVO](),
		rgbaIndex: -1,
		lstPO:     linkTraceRepository.ToEntity(traceId),
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
	lst       collections.List[response.LinkTraceVO]      // 要返回的集合
	rgbaIndex int                                         // 实现不同服务的调用，用颜色区分。这里通过服务入口调用时，使用索引（对应数组颜色）
	lstPO     collections.List[linkTraceCom.TraceContext] // 数据库读的集合
	startTs   int64                                       // 初始开始时间
}

// 服务调用入口
func (receiver *linkTraceWarp) addEntry(po linkTraceCom.TraceContext) {
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
func (receiver *linkTraceWarp) addDetail(po linkTraceCom.TraceContext) {
	for _, detail := range po.List {
		m := detail.(map[string]any)
		baseDetailPO := mapper.Single[trace.BaseTraceDetail](m)
		detailTrace := response.LinkTraceVO{
			Rgba: response.RgbaList[receiver.rgbaIndex], AppId: po.AppId, AppIp: po.AppIp, AppName: po.AppName,
			StartTs: baseDetailPO.StartTs - receiver.startTs,
			UseTs:   baseDetailPO.UseTs.Microseconds(), UseDesc: baseDetailPO.UseTs.String()}

		switch baseDetailPO.CallType {
		case eumCallType.Database:
			detailPO := mapper.Single[linkTraceCom.TraceDetailDatabase](m)
			if detailPO.TableName == "" && detailPO.Sql == "" {
				detailTrace.Caption = fmt.Sprintf("打开数据库 => %s %s", detailPO.DbName, detailPO.ConnectionString)
			} else {
				detailTrace.Caption = fmt.Sprintf("执行数据库 => %s %s 影响%v行", detailPO.DbName, detailPO.TableName, detailPO.RowsAffected)
			}
			detailTrace.Desc = detailPO.Sql
		case eumCallType.Http:
			detailPO := mapper.Single[linkTraceCom.TraceDetailHttp](m)
			detailTrace.Caption = fmt.Sprintf("调用http => %v %s %s", detailPO.StatusCode, detailPO.Method, detailPO.Url)
			lstHeader := collections.NewList[string]()
			for k, v := range detailPO.Headers.ToMap() {
				lstHeader.Add(fmt.Sprintf("%s=%v", k, v))
			}
			detailTrace.Desc = fmt.Sprintf("头部：%s 入参：%s 出参：%s", lstHeader.ToString(","), detailPO.RequestBody, detailPO.ResponseBody)
		case eumCallType.Grpc:
			detailPO := mapper.Single[linkTraceCom.TraceDetailGrpc](m)
			detailTrace.Caption = fmt.Sprintf("调用http => %v %s %s", detailPO.StatusCode, detailPO.Method, detailPO.Url)
			lstHeader := collections.NewList[string]()
			for k, v := range detailPO.Headers.ToMap() {
				lstHeader.Add(fmt.Sprintf("%s=%v", k, v))
			}
			detailTrace.Desc = fmt.Sprintf("头部：%s 入参：%s 出参：%s", lstHeader.ToString(","), detailPO.RequestBody, detailPO.ResponseBody)
		case eumCallType.Redis:
			detailPO := mapper.Single[linkTraceCom.TraceDetailRedis](m)
			detailTrace.Caption = fmt.Sprintf("执行Redis => %s %s %s", detailPO.MethodName, detailPO.Key, detailPO.Field)
			detailTrace.Desc = fmt.Sprintf("%s %s", detailPO.Key, detailPO.Field)
		case eumCallType.Mq:
			detailPO := mapper.Single[linkTraceCom.TraceDetailMq](m)
			detailTrace.Caption = fmt.Sprintf("发送MQ消息 => %s %s %s", detailPO.Server, detailPO.Exchange, detailPO.RoutingKey)
			detailTrace.Desc = fmt.Sprintf("%s %s %s", detailPO.Server, detailPO.Exchange, detailPO.RoutingKey)
		case eumCallType.Elasticsearch:
			detailPO := mapper.Single[linkTraceCom.TraceDetailEs](m)
			detailTrace.Caption = fmt.Sprintf("执行ES => %s %s", detailPO.IndexName, detailPO.AliasesName)
			detailTrace.Desc = fmt.Sprintf("%s %s", detailPO.IndexName, detailPO.AliasesName)
		case eumCallType.Etcd:
			detailPO := mapper.Single[linkTraceCom.TraceDetailEtcd](m)
			detailTrace.Caption = fmt.Sprintf("执行Etcd => %s %v", detailPO.Key, detailPO.LeaseID)
			detailTrace.Desc = fmt.Sprintf("%s %v", detailPO.Key, detailPO.LeaseID)
		case eumCallType.Hand:
			detailPO := mapper.Single[linkTraceCom.TraceDetailHand](m)
			detailTrace.Caption = fmt.Sprintf("手动埋点 => %s", detailPO.Name)
			detailTrace.Desc = fmt.Sprintf("%s", detailPO.Name)
		}
		receiver.lst.Add(detailTrace)

		if baseDetailPO.CallType == eumCallType.Http {
			// 查找串联的服务
			nextEntry := receiver.lstPO.Where(func(item linkTraceCom.TraceContext) bool {
				return item.ParentAppName == detailTrace.AppName
			}).First()
			receiver.addEntry(nextEntry)
		}
	}
}
