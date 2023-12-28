// @area /linkTrace/
package linkTraceApp

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/linkTrace/driver/clickhouse"
	"github.com/farseer-go/linkTrace/eumTraceType"
	"strings"
	"time"
)

// WebApiList WebApi链路追踪列表
// @get webApiList
func WebApiList(appName, appIp, requestIp, searchUrl string, statusCode int, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceContextPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	requestIp = strings.TrimSpace(requestIp)
	searchUrl = strings.TrimSpace(searchUrl)
	return linkTrace_clickhouse.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,start_ts,end_ts,use_ts,web_domain,web_path,web_method,web_content_type,web_status_code,web_request_ip").
		Where("trace_type = ? and parent_app_name = ''", eumTraceType.WebApi).
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(requestIp != "", "web_request_ip = ?", requestIp).
		WhereIf(searchUrl != "", "web_path like ?", searchUrl).
		WhereIf(statusCode > 0, "web_status_code = ?", statusCode).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)
}

// TaskList Task链路追踪列表
// @get taskList
func TaskList(appName, appIp, taskName string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceContextPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	taskName = strings.TrimSpace(taskName)
	return linkTrace_clickhouse.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,start_ts,end_ts,use_ts,task_name").
		Where("trace_type = ? and parent_app_name = ''", eumTraceType.Task).
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(taskName != "", "task_Name like ?", taskName).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)
}

// FScheduleList FSchedule链路追踪列表
// @get fScheduleList
func FScheduleList(appName, appIp, taskName string, taskGroupId, taskId, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceContextPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	taskName = strings.TrimSpace(taskName)
	return linkTrace_clickhouse.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,start_ts,end_ts,use_ts,task_name,task_group_id,task_id").
		Where("trace_type = ? and parent_app_name = ''", eumTraceType.FSchedule).
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(taskName != "", "task_Name like ?", taskName).
		WhereIf(taskGroupId > 0, "task_group_id = ?", taskGroupId).
		WhereIf(taskId > 0, "task_id = ?", taskId).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)
}

// ConsumerList Consumer链路追踪列表
// @get consumerList
func ConsumerList(appName, appIp, server, queueName, routingKey string, searchUseTs int64, startMin int, pageSize, pageIndex int) collections.PageList[linkTrace_clickhouse.TraceContextPO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	server = strings.TrimSpace(server)
	queueName = strings.TrimSpace(queueName)
	routingKey = strings.TrimSpace(routingKey)
	return linkTrace_clickhouse.CHContext.TraceContext.Select("trace_id,app_id,app_name,app_ip,parent_app_name,start_ts,end_ts,use_ts,consumer_server,consumer_queue_name,consumer_routing_key").
		Where("(trace_type = ? or trace_type = ?) and parent_app_name = ''", eumTraceType.MqConsumer, eumTraceType.QueueConsumer).
		WhereIf(appName != "", "app_name = ?", appName).
		WhereIf(appIp != "", "app_ip = ?", appIp).
		WhereIf(searchUseTs > 0, "use_ts >= ?", searchUseTs*int64(time.Microsecond)).
		WhereIf(server != "", "consumer_server like ?", server).
		WhereIf(queueName != "", "consumer_queue_name like ?", queueName).
		WhereIf(routingKey != "", "consumer_routing_key like ?", routingKey).
		WhereIf(startMin > 0, "start_ts >= ?", dateTime.Now().AddMinutes(-startMin).UnixMicro()).
		Desc("start_ts").ToPageList(pageSize, pageIndex)
}
