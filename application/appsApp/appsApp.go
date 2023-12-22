// @area /apps/
package appsApp

import (
	"fops/application/appsApp/request"
	"fops/application/appsApp/response"
	"fops/domain/apps"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/mapper"
	"strings"
)

// Add 添加应用
// @post add
func Add(req request.AddRequest, appsRepository apps.Repository) {
	do := mapper.Single[apps.DomainObject](req)
	exception.ThrowWebExceptionBool(appsRepository.IsExists(req.AppName), 403, "应用不能重复")
	// 添加
	err := appsRepository.Add(do)
	exception.ThrowWebExceptionError(403, err)
}

// Update 修改应用
// @post update
func Update(req request.UpdateRequest, appsRepository apps.Repository) {
	do := mapper.Single[apps.DomainObject](req)
	err := appsRepository.UpdateApp(do)
	exception.ThrowWebExceptionError(403, err)
}

// Delete 删除应用
// @post delete
func Delete(appName string, appsRepository apps.Repository) {
	exception.ThrowWebExceptionBool(strings.Trim(appName, "") == "", 403, "参数不完整")
	_, err := appsRepository.Delete(appName)
	exception.ThrowWebExceptionError(403, err)
}

// List 应用列表
// @post list
func List(appsRepository apps.Repository) collections.List[response.AppsResponse] {
	lstDO := appsRepository.ToList()
	lst := collections.NewList[response.AppsResponse]()
	lstDO.Foreach(func(item *apps.DomainObject) {
		item.ShellScript = ""
		item.Dockerfile = ""
		lst.Add(doToAppsResponse(*item))
	})
	return lst
}

// Info 查询应用
// @post info
func Info(appName string, appsRepository apps.Repository) response.AppsResponse {
	do := appsRepository.ToEntity(appName)
	exception.ThrowWebExceptionBool(do.IsNil(), 403, "应用不存在")
	return doToAppsResponse(do)
	//rsp := mapper.Single[response.AppsResponse](do)
	//return rsp
}

func doToAppsResponse(do apps.DomainObject) response.AppsResponse {
	var clusterVers []apps.ClusterVerVO
	for _, clusterVerVO := range do.ClusterVer {
		clusterVers = append(clusterVers, *clusterVerVO)
	}
	return response.AppsResponse{
		AppName:        do.AppName,
		AppId:          do.AppId,
		DockerVer:      do.DockerVer,
		ShellScript:    do.ShellScript,
		ClusterVer:     clusterVers,
		AppGit:         do.AppGit,
		FrameworkGits:  do.FrameworkGits,
		Dockerfile:     do.Dockerfile,
		DockerfilePath: do.DockerfilePath,
	}
}
