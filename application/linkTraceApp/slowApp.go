// @area /linkTrace/
package linkTraceApp

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/linkTrace/driver/clickhouse"
	"strings"
	"time"
)

// SlowDbList 慢数据库列表
// @get slowDbList
func SlowDbList(appName, appIp, dbName, tableName string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailDatabasePO] {
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
	lst := linkTrace_clickhouse.CHContext.TraceDetailDatabase.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(dbName != "", "db_name like ?", "%"+dbName+"%").
		WhereIf(tableName != "", "table_name like ?", "%"+tableName+"%").
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailDatabasePO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowEsList 慢Es列表
// @get slowEsList
func SlowEsList(appName, appIp, indexName, aliasesName string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailEsPO] {
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
	lst := linkTrace_clickhouse.CHContext.TraceDetailEs.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(indexName != "", "index_name like ?", "%"+indexName+"%").
		WhereIf(aliasesName != "", "aliases_name like ?", "%"+aliasesName+"%").
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailEsPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowEtcdList 慢Etcd列表
// @get slowEtcdList
func SlowEtcdList(appName, appIp, key string, leaseID int64, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailEtcdPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	key = strings.TrimSpace(key)
	lst := linkTrace_clickhouse.CHContext.TraceDetailEtcd.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(key != "", "key like ?", "%"+key+"%").
		WhereIf(leaseID > 0, "leaseID = ?", leaseID).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailEtcdPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowHandList 慢手动列表
// @get slowHandList
func SlowHandList(appName, appIp, name string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailHandPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	name = strings.TrimSpace(name)
	lst := linkTrace_clickhouse.CHContext.TraceDetailHand.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(name != "", "name like ?", "%"+name+"%").
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailHandPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowHttpList 慢Http列表
// @get slowHttpList
func SlowHttpList(appName, appIp, method, url, requestBody, responseBody string, statusCode int, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailHttpPO] {
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
	lst := linkTrace_clickhouse.CHContext.TraceDetailHttp.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(method != "", "method like ?", "%"+method+"%").
		WhereIf(url != "", "url like ?", "%"+url+"%").
		WhereIf(requestBody != "", "request_body like ?", "%"+requestBody+"%").
		WhereIf(responseBody != "", "response_body like ?", "%"+responseBody+"%").
		WhereIf(statusCode > 0, "status_code >= ?", statusCode).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailHttpPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowMqList 慢Mq列表
// @get slowMqList
func SlowMqList(appName, appIp, server, exchange, routingKey string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailMqPO] {
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
	lst := linkTrace_clickhouse.CHContext.TraceDetailMq.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(server != "", "server like ?", "%"+server+"%").
		WhereIf(exchange != "", "url like ?", "%"+exchange+"%").
		WhereIf(routingKey != "", "url like ?", "%"+routingKey+"%").
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailMqPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}

// SlowRedisList 慢Redis列表
// @get slowRedisList
func SlowRedisList(appName, appIp, key, field string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceDetailRedisPO] {
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
	lst := linkTrace_clickhouse.CHContext.TraceDetailRedis.
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(key != "", "key like ?", "%"+key+"%").
		WhereIf(field != "", "field like ?", "%"+field+"%").
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)

	lst.List.Foreach(func(item *linkTrace_clickhouse.TraceDetailRedisPO) {
		item.UseDesc = item.UseTs.String()
	})
	return lst
}
