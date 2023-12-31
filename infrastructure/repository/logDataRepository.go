package repository

import (
	"fmt"
	"fops/domain/linkTrace"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/collections"
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
