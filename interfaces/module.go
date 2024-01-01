package interfaces

import (
	"context"
	"fops/application"
	"fops/interfaces/job"
	"github.com/farseer-go/fs/modules"
	"github.com/farseer-go/tasks"
	"github.com/farseer-go/webapi"
	"time"
)

type Module struct {
}

func (module Module) DependsModule() []modules.FarseerModule {
	return []modules.FarseerModule{webapi.Module{}, application.Module{}}
}

func (module Module) PostInitialize() {
	// 开启构建
	tasks.Run("开启构建", time.Second*1, job.BuildingJob, context.Background())
}
