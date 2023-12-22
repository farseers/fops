package repository

import (
	"fops/domain/cluster"
	"github.com/farseer-go/data"
)

type clusterRepository struct {
	// IRepository 通用的仓储接口
	data.IRepository[cluster.DomainObject]
}
