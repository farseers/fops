// @area /linkTrace/
package linkTraceApp

import (
	"fops/domain/linkTrace"
	"github.com/farseer-go/collections"
	linkTraceCom "github.com/farseer-go/linkTrace"
	"strings"
)

// SlowDbList 慢数据库列表
// @get slowDbList
func SlowDbList(appName, appIp, dbName, tableName string, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailDatabase] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	dbName = strings.TrimSpace(dbName)
	tableName = strings.TrimSpace(tableName)

	return linkTraceRepository.ToSlowDbList(appName, appIp, dbName, tableName, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowEsList 慢Es列表
// @get slowEsList
func SlowEsList(appName, appIp, indexName, aliasesName string, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailEs] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	indexName = strings.TrimSpace(indexName)
	aliasesName = strings.TrimSpace(aliasesName)

	return linkTraceRepository.ToSlowEsList(appName, appIp, indexName, aliasesName, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowEtcdList 慢Etcd列表
// @get slowEtcdList
func SlowEtcdList(appName, appIp, key string, leaseID int64, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailEtcd] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	key = strings.TrimSpace(key)

	return linkTraceRepository.ToSlowEtcdList(appName, appIp, key, leaseID, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowHandList 慢手动列表
// @get slowHandList
func SlowHandList(appName, appIp, name string, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailHand] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	name = strings.TrimSpace(name)

	return linkTraceRepository.ToSlowHandList(appName, appIp, name, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowHttpList 慢Http列表
// @get slowHttpList
func SlowHttpList(appName, appIp, method, url, requestBody, responseBody string, statusCode int, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailHttp] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	method = strings.TrimSpace(method)
	url = strings.TrimSpace(url)
	requestBody = strings.TrimSpace(requestBody)
	responseBody = strings.TrimSpace(responseBody)

	return linkTraceRepository.ToSlowHttpList(appName, appIp, method, url, requestBody, responseBody, statusCode, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowMqList 慢Mq列表
// @get slowMqList
func SlowMqList(appName, appIp, server, exchange, routingKey string, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailMq] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	server = strings.TrimSpace(server)
	exchange = strings.TrimSpace(exchange)
	routingKey = strings.TrimSpace(routingKey)

	return linkTraceRepository.ToSlowMqList(appName, appIp, server, exchange, routingKey, searchUseTs, startMin, pageSize, pageIndex)
}

// SlowRedisList 慢Redis列表
// @get slowRedisList
func SlowRedisList(appName, appIp, key, field string, searchUseTs int64, startMin int, pageSize, pageIndex int, linkTraceRepository linkTrace.Repository) collections.PageList[linkTraceCom.TraceDetailRedis] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	key = strings.TrimSpace(key)
	field = strings.TrimSpace(field)

	return linkTraceRepository.ToSlowRedisList(appName, appIp, key, field, searchUseTs, startMin, pageSize, pageIndex)
}
