package repository

import (
	"fmt"
	"fops/domain/linkTrace"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/mapper"
)

type logDataRepository struct {
}

func (receiver *logDataRepository) Save(lstEO collections.List[flog.LogData]) error {
	lstPO := mapper.ToList[model.LogDataPO](lstEO)
	if linkTrace.Config.Driver == "clickhouse" {
		// 写入上下文
		_, err := context.CHContext.LogData.InsertList(lstPO, 2000)
		flog.ErrorIfExists(err)
	} else {
		return fmt.Errorf("不支持的链路追踪驱动：%s", linkTrace.Config.Driver)
	}
	return nil
}

func (receiver *logDataRepository) ToList(traceId int64, appName, appIp, logContent string, logLevel eumLogLevel.Enum, pageSize, pageIndex int) collections.PageList[flog.LogData] {
	var lst collections.PageList[flog.LogData]
	if linkTrace.Config.Driver == "clickhouse" {
		lstPO := context.CHContext.LogData.
			WhereIf(traceId > 0, "trace_id = ?", traceId).
			WhereIf(appName != "", "app_name = ?", appName).
			WhereIf(appIp != "", "app_ip = ?", appIp).
			WhereIf(logLevel > -1, "log_level = ?", logLevel).
			WhereIf(logContent != "", "log_content like ?", "%"+logContent+"%").
			Desc("create_at").
			ToPageList(pageSize, pageIndex)
		lst = mapper.ToPageList[flog.LogData](lstPO)
	} else {
		exception.ThrowRefuseExceptionf("不支持的链路追踪驱动：%s", linkTrace.Config.Driver)
	}
	return lst
}

func (receiver *logDataRepository) ToInfo(id int64) flog.LogData {
	var do flog.LogData
	if linkTrace.Config.Driver == "clickhouse" {
		po := context.CHContext.LogData.Where("id = ?", id).ToEntity()
		do = mapper.Single[flog.LogData](po)
	} else {
		exception.ThrowRefuseExceptionf("不支持的链路追踪驱动：%s", linkTrace.Config.Driver)
	}
	return do
}
