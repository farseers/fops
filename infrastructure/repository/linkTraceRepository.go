package repository

import (
	"fmt"
	"fops/domain/linkTrace"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/trace/eumCallType"
	linkTraceCom "github.com/farseer-go/linkTrace"
	"github.com/farseer-go/linkTrace/eumTraceType"
	"github.com/farseer-go/mapper"
	"time"
)

type linkTraceRepository struct {
}

func (receiver *linkTraceRepository) ToEntity(traceId int64) collections.List[linkTraceCom.TraceContext] {
	var lstPO collections.List[model.TraceContextPO]
	lst := collections.NewList[linkTraceCom.TraceContext]()
	if linkTrace.Config.Driver == "clickhouse" {
		lstPO = context.CHContext.TraceContext.Where("trace_id", traceId).Asc("start_ts").ToList()
	}

	lstPO.Foreach(func(item *model.TraceContextPO) {
		do := mapper.Single[linkTraceCom.TraceContext](item)
		do.SetTraceIdN()
		do.List = []any{}
		for _, detail := range item.List {
			switch eumCallType.Enum(parse.ToInt(detail.(map[string]any)["CallType"])) {
			case eumCallType.Database:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailDatabase](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Http:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailHttp](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Grpc:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailGrpc](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Redis:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailRedis](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Mq:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailMq](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Elasticsearch:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailEs](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Etcd:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailEtcd](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			case eumCallType.Hand:
				traceDetail := mapper.Single[linkTraceCom.TraceDetailHand](detail)
				traceDetail.SetTraceIdN()
				do.List = append(do.List, &traceDetail)
			}
		}
		lst.Add(do)
	})
	return lst
}

func (receiver *linkTraceRepository) ToWebApiList(appName, appIp, requestIp, searchUrl string, statusCode int, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,trace_type,start_ts,end_ts,use_ts,use_desc,create_at,web_domain,web_path,web_method,web_content_type,web_status_code,web_request_ip,web_request_body,web_response_body,web_headers").
			Where("trace_type = ? and parent_app_name = ''", eumTraceType.WebApi).
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Millisecond)).
			WhereIf(requestIp != "", "web_request_ip = ?", requestIp).
			WhereIf(searchUrl != "", "web_path like ?", "%"+searchUrl+"%").
			WhereIf(statusCode > 0, "web_status_code = ?", statusCode).
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		return receiver.setTraceIdN(lstPO)
	}
	return collections.NewPageList[linkTraceCom.TraceContext](collections.NewList[linkTraceCom.TraceContext](), 0)
}
func (receiver *linkTraceRepository) ToTaskList(appName, appIp, taskName string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,trace_type,start_ts,end_ts,use_ts,use_desc,create_at,task_name").
			Where("trace_type = ? and parent_app_name = ''", eumTraceType.Task).
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Millisecond)).
			WhereIf(taskName != "", "task_Name like ?", "%"+taskName+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		return receiver.setTraceIdN(lstPO)
	}
	return collections.NewPageList[linkTraceCom.TraceContext](collections.NewList[linkTraceCom.TraceContext](), 0)
}
func (receiver *linkTraceRepository) ToFScheduleList(appName, appIp, taskName string, taskGroupId, taskId, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,trace_type,start_ts,end_ts,use_ts,use_desc,create_at,task_name,task_group_id,task_id").
			Where("trace_type = ? and parent_app_name = ''", eumTraceType.FSchedule).
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Millisecond)).
			WhereIf(taskName != "", "task_Name like ?", "%"+taskName+"%").
			WhereIf(taskGroupId > 0, "task_group_id = ?", taskGroupId).
			WhereIf(taskId > 0, "task_id = ?", taskId).
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		return receiver.setTraceIdN(lstPO)
	}
	return collections.NewPageList[linkTraceCom.TraceContext](collections.NewList[linkTraceCom.TraceContext](), 0)
}
func (receiver *linkTraceRepository) ToConsumerList(appName, appIp, server, queueName, routingKey string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,trace_type,start_ts,end_ts,use_ts,use_desc,create_at,consumer_server,consumer_queue_name,consumer_routing_key").
			Where("(trace_type = ? or trace_type = ?) and parent_app_name = ''", eumTraceType.MqConsumer, eumTraceType.QueueConsumer).
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Millisecond)).
			WhereIf(server != "", "consumer_server like ?", "%"+server+"%").
			WhereIf(queueName != "", "consumer_queue_name like ?", "%"+queueName+"%").
			WhereIf(routingKey != "", "consumer_routing_key like ?", "%"+routingKey+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		return receiver.setTraceIdN(lstPO)
	}
	return collections.NewPageList[linkTraceCom.TraceContext](collections.NewList[linkTraceCom.TraceContext](), 0)
}

func (receiver *linkTraceRepository) ToSlowDbList(appName, appIp, dbName, tableName string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailDatabase] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailDatabase.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(dbName != "", "db_name like ?", "%"+dbName+"%").
			WhereIf(tableName != "", "table_name like ?", "%"+tableName+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)

		lst := mapper.ToPageList[linkTraceCom.TraceDetailDatabase](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailDatabase) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailDatabase](collections.NewList[linkTraceCom.TraceDetailDatabase](), 0)
}
func (receiver *linkTraceRepository) ToSlowEsList(appName, appIp, indexName, aliasesName string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailEs] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailEs.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(indexName != "", "index_name like ?", "%"+indexName+"%").
			WhereIf(aliasesName != "", "aliases_name like ?", "%"+aliasesName+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailEs](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailEs) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailEs](collections.NewList[linkTraceCom.TraceDetailEs](), 0)
}
func (receiver *linkTraceRepository) ToSlowEtcdList(appName, appIp, key string, leaseID, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailEtcd] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailEtcd.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(key != "", "key like ?", "%"+key+"%").
			WhereIf(leaseID > 0, "leaseID = ?", leaseID).
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailEtcd](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailEtcd) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailEtcd](collections.NewList[linkTraceCom.TraceDetailEtcd](), 0)
}
func (receiver *linkTraceRepository) ToSlowHandList(appName, appIp, name string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailHand] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailHand.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(name != "", "name like ?", "%"+name+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailHand](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailHand) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailHand](collections.NewList[linkTraceCom.TraceDetailHand](), 0)
}
func (receiver *linkTraceRepository) ToSlowHttpList(appName, appIp, method, url, requestBody, responseBody string, statusCode int, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailHttp] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailHttp.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(method != "", "method like ?", "%"+method+"%").
			WhereIf(url != "", "url like ?", "%"+url+"%").
			WhereIf(requestBody != "", "request_body like ?", "%"+requestBody+"%").
			WhereIf(responseBody != "", "response_body like ?", "%"+responseBody+"%").
			WhereIf(statusCode > 0, "status_code >= ?", statusCode).
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailHttp](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailHttp) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailHttp](collections.NewList[linkTraceCom.TraceDetailHttp](), 0)
}
func (receiver *linkTraceRepository) ToSlowMqList(appName, appIp, server, exchange, routingKey string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailMq] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailMq.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(server != "", "server like ?", "%"+server+"%").
			WhereIf(exchange != "", "url like ?", "%"+exchange+"%").
			WhereIf(routingKey != "", "url like ?", "%"+routingKey+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailMq](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailMq) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailMq](collections.NewList[linkTraceCom.TraceDetailMq](), 0)
}
func (receiver *linkTraceRepository) ToSlowRedisList(appName, appIp, key, field string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailRedis] {
	if linkTrace.Config.Driver == "clickhouse" {
		ts := context.CHContext.TraceDetailRedis.
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
			WhereIf(key != "", "key like ?", "%"+key+"%").
			WhereIf(field != "", "field like ?", "%"+field+"%").
			WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro())

		lstPO := ts.DescIfElse(startMin > 0, "use_ts", "start_ts").ToPageList(pageSize, pageIndex)
		lst := mapper.ToPageList[linkTraceCom.TraceDetailRedis](lstPO)
		lst.List.Foreach(func(item *linkTraceCom.TraceDetailRedis) {
			item.SetTraceIdN()
		})
		return lst
	}
	return collections.NewPageList[linkTraceCom.TraceDetailRedis](collections.NewList[linkTraceCom.TraceDetailRedis](), 0)
}

func (receiver *linkTraceRepository) Save(lstEO collections.List[linkTraceCom.TraceContext]) error {
	lst := mapper.ToList[model.TraceContextPO](lstEO)
	lst.Foreach(func(item *model.TraceContextPO) {
		item.UseDesc = item.UseTs.String()
		item.CreateAt = dateTime.NewUnixMicro(item.StartTs)

		for index, detail := range item.List {
			m := detail.(map[string]any)
			baseDetailPO := mapper.Single[model.BaseTraceDetailPO](m)
			m["UseDesc"] = baseDetailPO.UseTs.String()
			m["CreateAt"] = dateTime.NewUnixMicro(baseDetailPO.StartTs)
			item.List[index] = m
		}
	})

	if linkTrace.Config.Driver == "clickhouse" {
		// 写入上下文
		_, err := context.CHContext.TraceContext.InsertList(lst, 2000)
		flog.ErrorIfExists(err)
	} else {
		return fmt.Errorf("不支持的链路追踪驱动：%s", linkTrace.Config.Driver)
	}

	// 写入明细
	return receiver.saveDetail(lst)
}

// 写入明细
func (receiver *linkTraceRepository) saveDetail(lst collections.List[model.TraceContextPO]) error {
	lstTraceDetailDatabase := collections.NewList[model.TraceDetailDatabasePO]()
	lstTraceDetailEs := collections.NewList[model.TraceDetailEsPO]()
	lstTraceDetailEtcd := collections.NewList[model.TraceDetailEtcdPO]()
	lstTraceDetailHand := collections.NewList[model.TraceDetailHandPO]()
	lstTraceDetailHttp := collections.NewList[model.TraceDetailHttpPO]()
	lstTraceDetailGrpc := collections.NewList[model.TraceDetailGrpcPO]()
	lstTraceDetailMq := collections.NewList[model.TraceDetailMqPO]()
	lstTraceDetailRedis := collections.NewList[model.TraceDetailRedisPO]()

	lst.Foreach(func(traceContext *model.TraceContextPO) {
		for _, detail := range traceContext.List {
			m := detail.(map[string]any)
			baseDetailPO := mapper.Single[model.BaseTraceDetailPO](m)
			switch baseDetailPO.CallType {
			case eumCallType.Database:
				detailPO := mapper.Single[model.TraceDetailDatabasePO](m)
				lstTraceDetailDatabase.Add(detailPO)
			case eumCallType.Http:
				detailPO := mapper.Single[model.TraceDetailHttpPO](m)
				lstTraceDetailHttp.Add(detailPO)
			case eumCallType.Grpc:
				detailPO := mapper.Single[model.TraceDetailGrpcPO](m)
				lstTraceDetailGrpc.Add(detailPO)
			case eumCallType.Redis:
				detailPO := mapper.Single[model.TraceDetailRedisPO](m)
				lstTraceDetailRedis.Add(detailPO)
			case eumCallType.Mq:
				detailPO := mapper.Single[model.TraceDetailMqPO](m)
				lstTraceDetailMq.Add(detailPO)
			case eumCallType.Elasticsearch:
				detailPO := mapper.Single[model.TraceDetailEsPO](m)
				lstTraceDetailEs.Add(detailPO)
			case eumCallType.Etcd:
				detailPO := mapper.Single[model.TraceDetailEtcdPO](m)
				lstTraceDetailEtcd.Add(detailPO)
			case eumCallType.Hand:
				detailPO := mapper.Single[model.TraceDetailHandPO](m)
				lstTraceDetailHand.Add(detailPO)
			}
		}
	})

	if linkTrace.Config.Driver == "clickhouse" {
		// 写入明细
		if lstTraceDetailDatabase.Count() > 0 {
			if _, err := context.CHContext.TraceDetailDatabase.InsertList(lstTraceDetailDatabase, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailEs.Count() > 0 {
			if _, err := context.CHContext.TraceDetailEs.InsertList(lstTraceDetailEs, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailEtcd.Count() > 0 {
			if _, err := context.CHContext.TraceDetailEtcd.InsertList(lstTraceDetailEtcd, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailHand.Count() > 0 {
			if _, err := context.CHContext.TraceDetailHand.InsertList(lstTraceDetailHand, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailHttp.Count() > 0 {
			if _, err := context.CHContext.TraceDetailHttp.InsertList(lstTraceDetailHttp, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailGrpc.Count() > 0 {
			if _, err := context.CHContext.TraceDetailGrpc.InsertList(lstTraceDetailGrpc, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailMq.Count() > 0 {
			if _, err := context.CHContext.TraceDetailMq.InsertList(lstTraceDetailMq, 2000); err != nil {
				return err
			}
		}
		if lstTraceDetailRedis.Count() > 0 {
			if _, err := context.CHContext.TraceDetailRedis.InsertList(lstTraceDetailRedis, 2000); err != nil {
				return err
			}
		}
	}

	return nil
}

func (receiver *linkTraceRepository) setTraceIdN(lstPO collections.PageList[model.TraceContextPO]) collections.PageList[linkTraceCom.TraceContext] {
	lst := mapper.ToPageList[linkTraceCom.TraceContext](lstPO)
	lst.List.Foreach(func(item *linkTraceCom.TraceContext) {
		item.SetTraceIdN()
	})
	return lst
}
