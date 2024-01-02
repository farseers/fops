package register

import (
	"github.com/farseer-go/data"
)

// Repository 仓储接口
type Repository interface {
	// IRepository 通用的仓储接口
	data.IRepository[DomainObject]
	// StatRegisterApp 统计应用的在线实例
	StatRegisterApp() map[string][]DomainObject
	// ClearHistory 清除历史注册信息
	ClearHistory()
}
