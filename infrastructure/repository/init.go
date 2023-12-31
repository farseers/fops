package repository

import (
	"fops/domain/apps"
	"fops/domain/cluster"
	"fops/domain/linkTrace"
	"fops/domain/logData"
	"github.com/farseer-go/fs/container"
)

// InitRepository 初始化仓储
func InitRepository() {
	// 应用
	container.Register(func() apps.Repository { return &appsRepository{} })
	// 集群
	container.Register(func() cluster.Repository { return &clusterRepository{} })
	// 链路追踪
	container.Register(func() linkTrace.Repository { return &linkTraceRepository{} })
	// 日志
	container.Register(func() logData.Repository { return &logDataRepository{} })
}
