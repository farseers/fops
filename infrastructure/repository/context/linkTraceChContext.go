package context

import (
	"fops/domain/linkTrace"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/data"
)

var LinkTraceCHContext *linkTraceChContext

// EsContext 链路追踪上下文
type linkTraceChContext struct {
	TraceContext        data.TableSet[model.TraceContextPO]        `data:"name=link_trace;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id) PARTITION BY app_name"`
	TraceDetailDatabase data.TableSet[model.TraceDetailDatabasePO] `data:"name=trace_detail_database;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,db_name,table_name,connection_string) PARTITION BY app_name"`
	TraceDetailEs       data.TableSet[model.TraceDetailEsPO]       `data:"name=trace_detail_es;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,index_name,aliases_name) PARTITION BY app_name"`
	TraceDetailEtcd     data.TableSet[model.TraceDetailEtcdPO]     `data:"name=trace_detail_etcd;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,key) PARTITION BY app_name"`
	TraceDetailHand     data.TableSet[model.TraceDetailHandPO]     `data:"name=trace_detail_hand;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,name) PARTITION BY app_name"`
	TraceDetailHttp     data.TableSet[model.TraceDetailHttpPO]     `data:"name=trace_detail_http;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,method,url) PARTITION BY app_name"`
	TraceDetailGrpc     data.TableSet[model.TraceDetailGrpcPO]     `data:"name=trace_detail_grpc;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,method,url) PARTITION BY app_name"`
	TraceDetailMq       data.TableSet[model.TraceDetailMqPO]       `data:"name=trace_detail_mq;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,server,exchange,routing_key) PARTITION BY app_name"`
	TraceDetailRedis    data.TableSet[model.TraceDetailRedisPO]    `data:"name=trace_detail_redis;migrate=ReplacingMergeTree() ORDER BY (app_name,parent_app_name,app_ip,app_id,trace_id,key,field) PARTITION BY app_name"`
}

// InitLinkTraceChContextContext 初始化上下文
func InitLinkTraceChContextContext() {
	if linkTrace.Config.ConnString == "" {
		panic("[farseer.yaml]FOPS.LinkTrace.ConnString，配置不正确")
	}
	data.RegisterInternalContext("LinkTrace", linkTrace.Config.ConnString)
	LinkTraceCHContext = data.NewContext[linkTraceChContext]("LinkTrace")
}
