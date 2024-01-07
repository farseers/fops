package logData

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/flog"
)

// Repository 仓储接口
type Repository interface {
	Save(lstEO collections.List[flog.LogData]) error
	ToInfo(id string) flog.LogData
	ToList(traceId string, appName, appIp, logContent string, logLevel eumLogLevel.Enum, pageSize, pageIndex int) collections.PageList[flog.LogData]
}
