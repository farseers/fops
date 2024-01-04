// @area /apps/
package appsApp

import (
	"fops/application/appsApp/request"
	"fops/application/appsApp/response"
	"fops/domain/apps"
	"fops/domain/cluster"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/parse"
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
func Update(req request.UpdateRequest, appsRepository apps.Repository, appsIDockerDevice apps.IDockerDevice) {
	do := appsRepository.ToEntity(req.AppName)
	exception.ThrowWebExceptionBool(do.IsNil(), 403, "应用不存在")

	// 判断副本数量是否有变更
	if do.DockerReplicas != req.DockerReplicas {
		c := make(chan string, 100)
		if !appsIDockerDevice.SetReplicas(cluster.DomainObject{}, req.AppName, do.DockerReplicas, c) {
			lstLog := collections.NewListFromChan(c)
			exception.ThrowWebExceptionf(403, "更新副本失败:<br />%s", lstLog.ToString("<br />"))
		}
	}

	// 更新应用信息
	do = mapper.Single[apps.DomainObject](req)
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
	lstGit := appsRepository.ToGitListAll(-1)

	lst := collections.NewList[response.AppsResponse]()
	lstDO.Foreach(func(item *apps.DomainObject) {
		item.ShellScript = ""
		item.Dockerfile = ""
		appsResponse := doToAppsResponse(*item)
		appsResponse.AppGitName = lstGit.Where(func(gitItem apps.GitEO) bool {
			return item.AppGit == parse.ToInt64(gitItem.Id)
		}).First().Name
		lst.Add(appsResponse)
	})
	return lst
}

// Info 查询应用
// @post info
func Info(appName string, appsRepository apps.Repository) response.AppsResponse {
	do := appsRepository.ToEntity(appName)
	exception.ThrowWebExceptionBool(do.IsNil(), 403, "应用不存在")
	rsp := doToAppsResponse(do)
	rsp.AppGitName = appsRepository.ToGitEntity(do.AppGit).Name
	return rsp
}

func doToAppsResponse(do apps.DomainObject) response.AppsResponse {
	var clusterVers []apps.ClusterVerVO
	for _, clusterVerVO := range do.ClusterVer {
		clusterVers = append(clusterVers, *clusterVerVO)
	}
	return response.AppsResponse{
		AppName:        do.AppName,
		ActiveInstance: do.ActiveInstance,
		DockerVer:      do.DockerVer,
		ShellScript:    do.ShellScript,
		ClusterVer:     clusterVers,
		AppGit:         do.AppGit,
		FrameworkGits:  do.FrameworkGits,
		Dockerfile:     do.Dockerfile,
		DockerfilePath: do.DockerfilePath,
		IsHealth:       len(do.ActiveInstance) >= do.DockerReplicas,
	}
}
