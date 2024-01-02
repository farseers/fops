package repository

import (
	_ "embed"
	"fops/domain/register"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/data"
	"github.com/farseer-go/mapper"
)

type registerRepository struct {
	// IRepository 通用的仓储接口
	data.IRepository[register.DomainObject]
}

//go:embed model/sql/statRegisterApp.sql
var statRegisterAppSql string

func (receiver *registerRepository) StatRegisterApp() map[string][]register.DomainObject {
	var lstPO []model.RegisterPO
	_, _ = context.MysqlContext.ExecuteSqlToResult(&lstPO, statRegisterAppSql)

	lst := mapper.ToList[register.DomainObject](lstPO)
	var group map[string][]register.DomainObject
	lst.GroupBy(&group, func(item register.DomainObject) any {
		return item.AppName
	})
	return group
}
