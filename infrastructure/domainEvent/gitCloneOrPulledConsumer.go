package domainEvent

import (
	"fops/domain/apps"
	"fops/domain/apps/event"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/core"
	"time"
)

func GitCloneOrPulledConsumer(message any, ea core.EventArgs) {
	appsRepository := container.Resolve[apps.Repository]()
	gitCloneOrPulledEvent := message.(event.GitCloneOrPulledEvent)

	// 更新git拉取时间
	_, _ = appsRepository.UpdateForTime(gitCloneOrPulledEvent.GitId, time.UnixMicro(ea.CreateAt))
}
