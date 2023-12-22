package repository

import (
	"fops/domain/apps"
	"fops/domain/cluster"
	"github.com/farseer-go/fs/container"
)

// InitRepository 初始化仓储
func InitRepository() {
	// 应用
	container.Register(func() apps.Repository { return &appsRepository{} })
	// 集群
	container.Register(func() cluster.Repository { return &clusterRepository{} })
}
