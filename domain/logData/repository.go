package logData

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/flog"
)

// Repository 仓储接口
type Repository interface {
	Save(lstEO collections.List[flog.LogData]) error
}
