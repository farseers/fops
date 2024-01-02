package response

import (
	"fops/domain/apps"
	"github.com/farseer-go/collections"
)

type AppsResponse struct {
	AppName        string                  // 应用名称（链路追踪）
	DockerVer      int                     // 镜像版本
	ShellScript    string                  // Shell脚本
	ClusterVer     []apps.ClusterVerVO     // 集群版本
	AppGit         int64                   // 应用的git仓库
	AppGitName     string                  // 应用的git仓库名称
	FrameworkGits  collections.List[int64] // 依赖的框架源代码
	Dockerfile     string                  // Dockerfile内容
	DockerfilePath string                  // Dockerfile路径
}
