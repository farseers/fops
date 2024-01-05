// 该文件由fsctl route命令自动生成，请不要手动修改此文件
package main

import (
	"fops/application/appsApp"
	"fops/application/clusterApp"
	"fops/application/flogApp"
	"fops/application/gitApp"
	"fops/application/linkTraceApp"
	"github.com/farseer-go/webapi"
	"github.com/farseer-go/webapi/context"
)

var route = []webapi.Route{
	{"POST", "/apps/add", appsApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/apps/update", appsApp.Update, "", []context.IFilter{}, []string{"req", "", ""}},
	{"POST", "/apps/delete", appsApp.Delete, "", []context.IFilter{}, []string{"appName", "", ""}},
	{"POST", "/apps/list", appsApp.List, "", []context.IFilter{}, []string{"clusterId", ""}},
	{"POST", "/apps/info", appsApp.Info, "", []context.IFilter{}, []string{"appName", ""}},
	{"POST", "/apps/build/add", appsApp.BuildAdd, "", []context.IFilter{}, []string{"appName", "clusterId", "", ""}},
	{"POST", "/apps/build/list", appsApp.BuildList, "", []context.IFilter{}, []string{"appName", "pageSize", "pageIndex", ""}},
	{"POST", "/apps/build/clearDockerImage", appsApp.ClearDockerImage, "", []context.IFilter{}, []string{""}},
	{"GET", "/apps/build/view-{buildId}", appsApp.View, "", []context.IFilter{}, []string{"buildId"}},
	{"POST", "/apps/register", appsApp.Register, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/cluster/add", clusterApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/cluster/update", clusterApp.Update, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/cluster/list", clusterApp.List, "", []context.IFilter{}, []string{""}},
	{"POST", "/cluster/delete", clusterApp.Delete, "", []context.IFilter{}, []string{"clusterId", ""}},
	{"GET", "/flog/list", flogApp.List, "", []context.IFilter{}, []string{"traceId", "appName", "appIp", "logContent", "logLevel", "pageSize", "pageIndex", ""}},
	{"GET", "/flog/info-{id}", flogApp.Info, "", []context.IFilter{}, []string{"id", ""}},
	{"POST", "/flog/upload", flogApp.Upload, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/git/add", gitApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/git/update", gitApp.Update, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/git/list", gitApp.List, "", []context.IFilter{}, []string{"isApp", "", ""}},
	{"POST", "/git/delete", gitApp.Delete, "", []context.IFilter{}, []string{"gitId", "", ""}},
	{"POST", "/git/clear", gitApp.Clear, "", []context.IFilter{}, []string{"gitId", "", ""}},
	{"POST", "/git/pull", gitApp.Pull, "", []context.IFilter{}, []string{"gitId", "", ""}},
	{"POST", "/git/info", gitApp.Info, "", []context.IFilter{}, []string{"gitId", ""}},
	{"GET", "/linkTrace/info/{traceId}", linkTraceApp.Info, "", []context.IFilter{}, []string{"traceId", ""}},
	{"GET", "/linkTrace/webApiList", linkTraceApp.WebApiList, "", []context.IFilter{}, []string{"appName", "appIp", "requestIp", "searchUrl", "statusCode", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/taskList", linkTraceApp.TaskList, "", []context.IFilter{}, []string{"appName", "appIp", "taskName", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/fScheduleList", linkTraceApp.FScheduleList, "", []context.IFilter{}, []string{"appName", "appIp", "taskName", "taskGroupId", "taskId", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/consumerList", linkTraceApp.ConsumerList, "", []context.IFilter{}, []string{"appName", "appIp", "server", "queueName", "routingKey", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowDbList", linkTraceApp.SlowDbList, "", []context.IFilter{}, []string{"appName", "appIp", "dbName", "tableName", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowEsList", linkTraceApp.SlowEsList, "", []context.IFilter{}, []string{"appName", "appIp", "indexName", "aliasesName", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowEtcdList", linkTraceApp.SlowEtcdList, "", []context.IFilter{}, []string{"appName", "appIp", "key", "leaseID", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowHandList", linkTraceApp.SlowHandList, "", []context.IFilter{}, []string{"appName", "appIp", "name", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowHttpList", linkTraceApp.SlowHttpList, "", []context.IFilter{}, []string{"appName", "appIp", "method", "url", "requestBody", "responseBody", "statusCode", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowMqList", linkTraceApp.SlowMqList, "", []context.IFilter{}, []string{"appName", "appIp", "server", "exchange", "routingKey", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"GET", "/linkTrace/slowRedisList", linkTraceApp.SlowRedisList, "", []context.IFilter{}, []string{"appName", "appIp", "key", "field", "searchUseTs", "startMin", "pageSize", "pageIndex", ""}},
	{"POST", "/linkTrace/upload", linkTraceApp.Upload, "", []context.IFilter{}, []string{"req", ""}},
}
