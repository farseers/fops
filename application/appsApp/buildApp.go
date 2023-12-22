// @area /apps/
package appsApp

import (
	"fops/domain/apps"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/webapi/action"
	"strings"
)

// BuildAdd 添加构建
// @post build/add
func BuildAdd(appName string, clusterId int64, appsRepository apps.Repository) {
	appDO := appsRepository.ToEntity(appName)
	exception.ThrowWebExceptionfBool(appDO.IsNil(), 403, "应用不存在")

	buildNumber := appsRepository.GetBuildNumber(appName) + 1
	buildDO := apps.BuildEO{
		ClusterId:   clusterId,
		BuildNumber: buildNumber,
		CreateAt:    dateTime.Now(),
		FinishAt:    dateTime.Now(),
		Env:         apps.EnvVO{},
		AppName:     appName,
		Dockerfile:  appDO.Dockerfile,
		ShellScript: appDO.ShellScript,
	}
	err := appsRepository.AddBuild(buildDO)
	exception.ThrowWebExceptionError(403, err)
}

// BuildList 构建列表
// @post build/list
func BuildList(appName string, pageSize int, pageIndex int, appsRepository apps.Repository) collections.PageList[apps.BuildEO] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	return appsRepository.ToBuildList(appName, pageSize, pageIndex)
}

// View 构建日志
// @get build/view
func View(buildId int64) action.IResult {
	logQueue := apps.LogQueue{
		BuildId: buildId,
	}
	return action.Content(strings.Join(logQueue.View(), "\r\n"))
}
