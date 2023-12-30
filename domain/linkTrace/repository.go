package linkTrace

import (
	"github.com/farseer-go/collections"
	linkTraceCom "github.com/farseer-go/linkTrace"
)

// Repository 仓储接口
type Repository interface {
	// ToEntity 获取列表
	ToEntity(traceId int64) collections.List[linkTraceCom.TraceContext]
	ToWebApiList(appName, appIp, requestIp, searchUrl string, statusCode int, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext]
	ToTaskList(appName, appIp, taskName string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext]
	ToFScheduleList(appName, appIp, taskName string, taskGroupId, taskId, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext]
	ToConsumerList(appName, appIp, server, queueName, routingKey string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceContext]

	ToSlowDbList(appName, appIp, dbName, tableName string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailDatabase]
	ToSlowEsList(appName, appIp, indexName, aliasesName string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailEs]
	ToSlowEtcdList(appName, appIp, key string, leaseID, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailEtcd]
	ToSlowHandList(appName, appIp, name string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailHand]
	ToSlowHttpList(appName, appIp, method, url, requestBody, responseBody string, statusCode int, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailHttp]
	ToSlowMqList(appName, appIp, server, exchange, routingKey string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailMq]
	ToSlowRedisList(appName, appIp, key, field string, searchUseTs int64, startMin, pageSize, pageIndex int) collections.PageList[linkTraceCom.TraceDetailRedis]
	Save(lstEO collections.List[TraceContextEO]) error
}
