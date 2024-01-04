// @area /git/
package gitApp

import (
	"context"
	"fops/application/gitApp/request"
	"fops/domain/apps"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/mapper"
	"strings"
)

// Add 添加Git
// @post add
func Add(req request.AddRequest, appsRepository apps.Repository) {
	do := mapper.Single[apps.GitEO](req)
	if !strings.HasPrefix(do.Dir, "/") {
		do.Dir = "/" + do.Dir
	}
	// 添加
	err := appsRepository.AddGit(do)
	exception.ThrowWebExceptionError(403, err)
}

// Update 修改Git
// @post update
func Update(req request.UpdateRequest, appsRepository apps.Repository) {
	do := mapper.Single[apps.GitEO](req)
	exception.ThrowWebExceptionBool(!appsRepository.ExistsGit(req.Id), 403, "GitId不存在")

	if !strings.HasPrefix(do.Dir, "/") {
		do.Dir = "/" + do.Dir
	}
	_, err := appsRepository.UpdateGit(do)
	exception.ThrowWebExceptionError(403, err)
}

// List Git列表
// @post list
func List(appsRepository apps.Repository, appsIGitDevice apps.IGitDevice) collections.List[apps.GitEO] {
	return appsRepository.ToGitListAll()
}

// Delete 删除Git
// @post delete
func Delete(gitId int64, appsRepository apps.Repository, appsIGitDevice apps.IGitDevice) {
	exception.ThrowWebExceptionBool(gitId < 1, 403, "GitId没有填")
	gitEO := appsRepository.ToGitEntity(gitId)
	c := make(chan string, 100)
	if !appsIGitDevice.Clear(gitEO, c) {
		lstLog := collections.NewListFromChan(c)
		exception.ThrowWebExceptionf(403, "清除目录失败:<br />%s", lstLog.ToString("<br />"))
	}
	_, err := appsRepository.DeleteGit(gitId)
	exception.ThrowWebExceptionError(403, err)
}

// Clear 清除目录
// @post clear
func Clear(gitId int64, appsRepository apps.Repository, appsIGitDevice apps.IGitDevice) {
	exception.ThrowWebExceptionBool(gitId < 1, 403, "GitId没有填")
	gitEO := appsRepository.ToGitEntity(gitId)
	c := make(chan string, 100)
	if !appsIGitDevice.Clear(gitEO, c) {
		lstLog := collections.NewListFromChan(c)
		exception.ThrowWebExceptionf(403, "清除目录失败:<br />%s", lstLog.ToString("<br />"))
	}
}

// Pull 拉取仓库
// @post pull
func Pull(gitId int64, appsRepository apps.Repository, appsIGitDevice apps.IGitDevice) {
	exception.ThrowWebExceptionBool(gitId < 1, 403, "GitId没有填")
	gitEO := appsRepository.ToGitEntity(gitId)
	c := make(chan string, 100)
	if !appsIGitDevice.CloneOrPull(gitEO, c, context.Background()) {
		lstLog := collections.NewListFromChan(c)
		exception.ThrowWebExceptionf(403, "拉取目录失败:<br />%s", lstLog.ToString("<br />"))
	}
}

// Info 查询Git
// @post info
func Info(gitId int64, appsRepository apps.Repository) apps.GitEO {
	return appsRepository.ToGitEntity(gitId)
}
