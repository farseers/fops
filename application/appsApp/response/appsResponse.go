package response

import (
	"fops/domain/apps"
	"github.com/farseer-go/collections"
)

type AppsResponse struct {
	AppName           string                  // 应用名称（链路追踪）
	DockerVer         int                     // 镜像版本
	ShellScript       string                  // Shell脚本
	ClusterVer        apps.ClusterVerVO       // 集群版本
	DockerImage       string                  // 仓库镜像名称
	AppGit            int64                   // 应用的git仓库
	AppGitName        string                  // 应用的git仓库名称
	FrameworkGits     collections.List[int64] // 依赖的框架源代码
	Dockerfile        string                  // Dockerfile内容
	DockerfilePath    string                  // Dockerfile路径
	ActiveInstance    []apps.ActiveInstanceEO // 正在运行的实例
	IsHealth          bool                    // 应用是否健康
	DockerReplicas    int                     // 副本数量
	DockerNodeRole    string                  // 容器节点角色 manager or worker
	AdditionalScripts string                  // 首次创建应用时附加脚本
}
