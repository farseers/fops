// 该文件由fsctl route命令自动生成，请不要手动修改此文件
package main

import (
	"fops/application/appsApp"
	"fops/application/clusterApp"
	"fops/application/gitApp"
	"fops/application/linkTraceApp"
	"github.com/farseer-go/webapi"
	"github.com/farseer-go/webapi/context"
)

var route = []webapi.Route{
	{"POST", "/apps/add", appsApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/apps/update", appsApp.Update, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/apps/delete", appsApp.Delete, "", []context.IFilter{}, []string{"appName", ""}},
	{"POST", "/apps/list", appsApp.List, "", []context.IFilter{}, []string{""}},
	{"POST", "/apps/info", appsApp.Info, "", []context.IFilter{}, []string{"appName", ""}},
	{"POST", "/apps/build/add", appsApp.BuildAdd, "", []context.IFilter{}, []string{"appName", "clusterId", ""}},
	{"POST", "/apps/build/list", appsApp.BuildList, "", []context.IFilter{}, []string{"appName", "pageSize", "pageIndex", ""}},
	{"GET", "/apps/build/view", appsApp.View, "", []context.IFilter{}, []string{"buildId"}},
	{"POST", "/cluster/add", clusterApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/cluster/update", clusterApp.Update, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/cluster/list", clusterApp.List, "", []context.IFilter{}, []string{""}},
	{"POST", "/cluster/delete", clusterApp.Delete, "", []context.IFilter{}, []string{"clusterId", ""}},
	{"POST", "/git/add", gitApp.Add, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/git/update", gitApp.Update, "", []context.IFilter{}, []string{"req", ""}},
	{"POST", "/git/list", gitApp.List, "", []context.IFilter{}, []string{""}},
	{"POST", "/git/delete", gitApp.Delete, "", []context.IFilter{}, []string{"gitId", ""}},
	{"POST", "/git/info", gitApp.Info, "", []context.IFilter{}, []string{"gitId", ""}},
	{"GET", "/info", linkTraceApp.Info, "", []context.IFilter{}, []string{"traceId"}},
	{"GET", "/linkTrace/webApiList", linkTraceApp.WebApiList, "", []context.IFilter{}, []string{"appName", "appIp", "requestIp", "searchUrl", "searchUseTs", "statusCode", "pageSize", "pageIndex"}},
	{"GET", "/linkTrace/taskList", linkTraceApp.TaskList, "", []context.IFilter{}, []string{"appName", "appIp", "taskName", "searchUseTs", "pageSize", "pageIndex"}},
	{"GET", "/linkTrace/fScheduleList", linkTraceApp.FScheduleList, "", []context.IFilter{}, []string{"appName", "appIp", "taskName", "taskGroupId", "taskId", "searchUseTs", "pageSize", "pageIndex"}},
	{"GET", "/linkTrace/consumerList", linkTraceApp.ConsumerList, "", []context.IFilter{}, []string{"appName", "appIp", "server", "queueName", "routingKey", "searchUseTs", "pageSize", "pageIndex"}},
}
