package job

import (
	"fops/application/appsApp"
	"fops/application/appsApp/request"
	"fops/domain/register"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/tasks"
)

// RegisterFopsJob 注册自己
func RegisterFopsJob(*tasks.TaskContext) {
	req := request.RegisterAppRequest{StartupAt: core.StartupAt, AppName: core.AppName, HostName: core.HostName, AppId: core.AppId, AppIp: core.AppIp, ProcessId: core.ProcessId}
	appsApp.Register(req, container.Resolve[register.Repository]())
}
