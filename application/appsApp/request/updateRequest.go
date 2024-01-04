package request

import "github.com/farseer-go/collections"

type UpdateRequest struct {
	AppName           string                  // 应用名称
	ShellScript       string                  // Shell脚本
	AppGit            int64                   // 应用的源代码
	FrameworkGits     collections.List[int64] // 依赖的框架源代码
	Dockerfile        string                  // Dockerfile内容
	HealthInstance    int                     // 健康的实例数量
	DockerReplicas    int                     // 副本数量
	DockerNodeRole    string                  // 容器节点角色 manager or worker
	AdditionalScripts string                  // 首次创建应用时附加脚本
}
