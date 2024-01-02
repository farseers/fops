package job

import (
	"fops/domain/apps"
	"fops/domain/register"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/mapper"
	"github.com/farseer-go/tasks"
)

// StatAppOnlineJob 统计应用的在线实例
func StatAppOnlineJob(*tasks.TaskContext) {
	registerRepository := container.Resolve[register.Repository]()
	appsRepository := container.Resolve[apps.Repository]()
	lstGroupBy := registerRepository.StatRegisterApp()

	// 开启事务更新
	container.Resolve[core.ITransaction]("default").Transaction(func() {
		lstApps := appsRepository.ToList()
		lstApps.Foreach(func(item *apps.DomainObject) {
			if lstRegisterDO, exists := lstGroupBy[item.AppName]; exists {
				item.ActiveInstance = mapper.Array[apps.ActiveInstanceEO](lstRegisterDO)
			} else {
				item.ActiveInstance = make([]apps.ActiveInstanceEO, 0)
			}
			_ = appsRepository.UpdateApp(*item)
		})
	})
}
