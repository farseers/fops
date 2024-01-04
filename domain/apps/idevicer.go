package apps

import (
	"context"
	"fops/domain/_/eumK8SControllers"
	"fops/domain/cluster"
)

type ICopyToDistDevice interface {
	Copy(lstGit []GitEO, env EnvVO, progress chan string)
}
type IDirectoryDevice interface {
	// Check 检查目录
	Check(progress chan string)
}
type IDockerDevice interface {
	// GetDockerHub 取得dockerHub
	GetDockerHub(dockerHubAddress string) string
	// GetDockerImage 生成镜像名称
	GetDockerImage(dockerHubAddress string, projectName string, buildNumber int) string
	// Login 登陆容器仓库
	Login(dockerHub string, loginName string, loginPwd string, progress chan string, env EnvVO, ctx context.Context) bool
	// ExistsDockerfile dockerfile文件是否存在
	ExistsDockerfile(dockerfilePath string) bool
	// CreateDockerfile 生成Dockerfile文件
	// projectName dockerfile文件地址
	// dockerfileContent 文件内容
	CreateDockerfile(projectName string, dockerfileContent string, ctx context.Context)
	// Build 容器构建
	Build(env EnvVO, progress chan string, ctx context.Context) bool
	// Push 上传镜像
	Push(env EnvVO, progress chan string, ctx context.Context) bool
	// SetImages 更新镜像版本
	SetImages(cluster cluster.DomainObject, appName string, dockerImages string, progress chan string, ctx context.Context) bool
	// ClearImages 清除镜像
	ClearImages(progress chan string) bool
}

type IGitDevice interface {
	// GetGitPath 获取项目GIT源代码存的位置/var/lib/fops/git/{gitName}/
	GetGitPath(gitHub string) string
	// RememberPassword 记住密码
	RememberPassword(env EnvVO, progress chan string)
	// ExistsGitProject 项目GIT是否存在
	ExistsGitProject(gitPath string) bool
	// Clear 消除仓库
	Clear(git GitEO, progress chan string) bool
	// CloneOrPull 根据判断是否存在Git目录，来决定返回Clone or pull
	CloneOrPull(git GitEO, progress chan string, ctx context.Context) bool
	// CloneOrPullAndDependent 拉取主仓库及依赖仓库
	CloneOrPullAndDependent(lstGit []GitEO, progress chan string, ctx context.Context) bool
}
type IKubectlDevice interface {
	// GetConfigFile 获取存储k8s Config的路径
	GetConfigFile(clusterName string) string
	// CreateConfigFile 创建用于K8S远程管理的配置文件
	CreateConfigFile(clusterName string, clusterConfig string) string
	// SetYaml 生成yaml文件，并执行kubectl apply命令
	SetYaml(clusterName string, projectName string, yamlContent string, progress chan string, ctx context.Context) bool
	// SetImages 更新k8s的镜像版本
	SetImages(cluster cluster.DomainObject, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool
	// SetImagesByClusterName 更新k8s的镜像版本
	SetImagesByClusterName(clusterName string, clusterConfig string, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool
}
type IShellDevice interface {
	// ExecShell 执行Shell脚本
	ExecShell(env EnvVO, shellScript string, progress chan string, ctx context.Context) bool
}
