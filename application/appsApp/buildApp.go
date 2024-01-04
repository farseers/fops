// @area /apps/
package appsApp

import (
	"fmt"
	"fops/domain/apps"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/webapi/action"
	"regexp"
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

// ClearDockerImage 清除Docker镜像
// @post build/clearDockerImage
func ClearDockerImage(device apps.IDockerDevice) {
	c := make(chan string, 100)
	if !device.ClearImages(c) {
		lstLog := collections.NewListFromChan(c)
		exception.ThrowWebExceptionf(403, "无用镜像清除失败:<br />%s", lstLog.ToString("<br />"))
	}
}

// 语法高亮
var chineseTips = collections.NewList(
	"环境变量：", "前置检查。", "先删除之前编译的目标文件。", "自动创建目录。", "前置检查通过。", "已经是最新的。", "拉取完成。", "登陆镜像仓库。", "镜像仓库登陆成功。",
	"开始镜像打包。", "镜像打包完成。", "开始上传镜像。", "镜像上传完成。", "开始更新K8S POD的镜像版本。", "更新镜像版本完成。")

// （黄色）
var cmdPrefix = collections.NewList(
	"git -C ",
	"docker login ",
	"docker build -t ",
	"docker push ",
	"kubectl set image ",
	"Successfully built ",
	"Successfully tagged ",
	"成功执行。",
	"构建完成。")

var cmdResultTips = collections.NewList(
	"Login Succeeded",
	"Determining projects to restore...")

var errorTips = collections.NewList(
	"Exception ",
	"Failed ",
	"ERROR:",
	"error ",
	"镜像打包出错了。",
	"镜像仓库登陆失败。",
	"执行失败，退出构建。",
	"K8S更新镜像失败。",
	"Unable to connect",
	"Cannot connect")

// View 构建日志
// @get build/view-{buildId}
func View(buildId int64) action.IResult {
	logQueue := apps.LogQueue{
		BuildId: buildId,
	}
	logContent := logQueue.View()
	for i := 0; i < len(logContent); i++ {
		if len(logContent[i]) < 20 {
			continue
		}
		dateTimePart := logContent[i][:19] // 提取日期时间部分
		logPart := logContent[i][20:]      // 提取日志内容部分
		dateTimePart = fmt.Sprintf("<span style=\"color:#9caf62\">%s</span>", dateTimePart)

		if chineseTips.Contains(logPart) {
			logPart = fmt.Sprintf("<span style=\"color:#cfbbfc\">%s</span>", logPart)
		} else if cmdResultTips.Contains(logPart) || strings.HasPrefix(logPart, "The push refers to repository ") || regexp.MustCompile(`\w+\.apps/\w+ image updated`).MatchString(logPart) {
			logPart = fmt.Sprintf("<span style=\"color:#fff\">%s</span>", logPart)
		} else if startsWithAny(cmdPrefix, logPart) {
			logPart = fmt.Sprintf("<span style=\"color:#ffe127\">%s</span>", logPart)
		} else if containsAny(errorTips, strings.ToLower(logPart)) {
			logPart = fmt.Sprintf("<span style=\"color:#ff5b5b\">%s</span>", logPart)
		} else {
			dockerLogMatch := regexp.MustCompile(`#\d+ \[.+ \d+/\d+\] (?P<cmd>.+)`).FindStringSubmatch(logPart)
			if dockerLogMatch == nil {
				dockerLogMatch = regexp.MustCompile(`#\d+ \d+\.\d+ (?P<cmd>.+)`).FindStringSubmatch(logPart)
			}
			if dockerLogMatch == nil {
				dockerLogMatch = regexp.MustCompile(`Step \d+/\d+ :(?P<cmd>.+)`).FindStringSubmatch(logPart)
			}
			if dockerLogMatch != nil {
				cmd := dockerLogMatch[1]
				logPart = strings.Replace(logPart, cmd, fmt.Sprintf("<span style=\"color:#38e4c6\">%s</span>", cmd), 1)
				logPart = fmt.Sprintf("<span style=\"color:#6a6964\">%s</span>", logPart)
			} else {
				logPart = fmt.Sprintf("<span style=\"color:#6a6964\">%s</span>", logPart)
			}
		}
		logContent[i] = dateTimePart + " " + logPart
	}
	return action.Content(strings.Join(logContent, "\n"))
}

// 检查字符串是否以指定字符串切片中的任意字符串开头
func startsWithAny(lst collections.List[string], str string) bool {
	return lst.Where(func(item string) bool {
		return strings.HasPrefix(str, item)
	}).Any()
}

// 检查字符串切片中是否包含指定子字符串
func containsAny(lst collections.List[string], substr string) bool {
	return lst.Where(func(item string) bool {
		return strings.Contains(substr, strings.ToLower(item))
	}).Any()
}
