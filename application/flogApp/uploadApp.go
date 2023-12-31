// @area /flog/
package flogApp

import (
	"fops/application/flogApp/request"
	"fops/domain/logData"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"strings"
)

// Upload 上传链路记录
// @post upload
func Upload(req request.UploadRequest, logDataRepository logData.Repository) {
	err := logDataRepository.Save(req.List)
	exception.ThrowWebExceptionError(403, err)
}

// List 日志列表
// @get list
func List(traceId int64, appName, appIp, logContent string, logLevel eumLogLevel.Enum, pageSize, pageIndex int, logDataRepository logData.Repository) collections.PageList[flog.LogData] {
	if pageSize < 1 {
		pageSize = 20
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	appName = strings.TrimSpace(appName)
	appIp = strings.TrimSpace(appIp)
	logContent = strings.TrimSpace(logContent)

	return logDataRepository.ToList(traceId, appName, appIp, logContent, logLevel, pageSize, pageIndex)
}

// Info 日志详情
// @get info
func Info(id int64, logDataRepository logData.Repository) flog.LogData {
	return logDataRepository.ToInfo(id)
}
