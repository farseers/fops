package infrastructure

import (
	"fops/domain/apps/event"
	"fops/infrastructure/device"
	"fops/infrastructure/domainEvent"
	"fops/infrastructure/repository"
	"fops/infrastructure/repository/context"
	"github.com/farseer-go/data"
	"github.com/farseer-go/eventBus"
	"github.com/farseer-go/fs/modules"
	linkTrace_clickhouse "github.com/farseer-go/linkTrace/driver/clickhouse"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	// 这些模块都是farseer-go内置的模块
	return []modules.FarseerModule{data.Module{}, eventBus.Module{}, linkTrace_clickhouse.Module{}}
}

func (module Module) PostInitialize() {
	// 初始化数据库上下文
	context.InitMysqlContext()
	// 初始化仓储
	repository.InitRepository()

	// 注册驱动
	device.RegisterCopyToDistDevice()
	device.RegisterDirectoryDevice()
	device.RegisterDockerDevice()
	device.RegisterGitDevice()
	device.RegisterKubectlDevice()
	device.RegisterShellDevice()

	eventBus.Subscribe(event.BuildFinishedEventName, domainEvent.BuildFinishedConsumer)
	eventBus.Subscribe(event.DockerPushedEventName, domainEvent.DockerPushedConsumer)
	eventBus.Subscribe(event.GitCloneOrPulledEventName, domainEvent.GitCloneOrPulledConsumer)
}
