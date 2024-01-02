package repository

import (
	"fops/domain/register"
	"github.com/farseer-go/data"
)

type registerRepository struct {
	// IRepository 通用的仓储接口
	data.IRepository[register.DomainObject]
}
