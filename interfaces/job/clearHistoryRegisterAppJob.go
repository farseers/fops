package job

import (
	"fops/domain/register"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/tasks"
)

// ClearHistoryRegisterAppJob 清除历史注册信息
func ClearHistoryRegisterAppJob(*tasks.TaskContext) {
	registerRepository := container.Resolve[register.Repository]()
	registerRepository.ClearHistory()
}
