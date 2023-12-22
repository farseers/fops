package apps

import (
	"fmt"
	"github.com/farseer-go/fs/parse"
)

const (
	FopsRoot         = "/var/lib/fops/"                   // FopsRoot Fops根目录
	KubeRoot         = "/var/lib/fops/kube/"              // KubeRoot kubectlConfig配置
	NpmModulesRoot   = "/var/lib/fops/npm"                // NpmModulesRoot NpmModules
	DistRoot         = "/var/lib/fops/dist/"              // DistRoot 编译保存的根目录
	GitRoot          = "/var/lib/fops/git/"               // GitRoot GIT根目录
	DockerfilePath   = "/var/lib/fops/dist/Dockerfile"    // DockerfilePath Dockerfile文件地址
	DockerIgnorePath = "/var/lib/fops/dist/.dockerignore" // DockerIgnorePath DockerIgnore文件地址
	ShellRoot        = "/var/lib/fops/shell/"             // ShellRoot 生成Shell脚本的存放路径
)

// EnvVO 构建时的环境变量
type EnvVO struct {
	BuildId     int64  // 构建主键
	BuildNumber int    // 构建版本号
	AppName     string // 项目名称
	AppGitRoot  string // Git仓库源代码根目录 /var/lib/fops/git/{gitName}/
	GitHub      string // Git仓库地址
	GitName     string // Git名称（项目的目录名称）
	DockerHub   string // Docker仓库地址
	DockerImage string // Docker镜像
}

// Print 打印环境变量
func (env *EnvVO) Print(progress chan string) {
	// 打印环境变量
	progress <- "---------------------------------------------------------"
	progress <- "环境变量："

	for k, v := range env.ToMap() {
		progress <- fmt.Sprint(k, "=", v)
	}
}

// ToMap 转成字典
func (env *EnvVO) ToMap() map[string]string {
	return map[string]string{
		"FopsRoot":       FopsRoot,
		"NpmModulesRoot": NpmModulesRoot,
		"DistRoot":       DistRoot,
		"KubeRoot":       KubeRoot,
		"Git_Root":       GitRoot,
		"Git_Hub":        env.GitHub,
		"Build_Id":       parse.ToString(env.BuildId),
		"Build_Number":   parse.ToString(env.BuildNumber),
		"App_Name":       env.AppName,
		"App_GitRoot":    env.AppGitRoot,
		//"App_Domain":     env.ProjectDomain,
		//"App_EntryPoint": env.ProjectEntryPoint,
		//"App_EntryPort":  strconv.Itoa(env.ProjectEntryPort),
		"Docker_Hub":   env.DockerHub,
		"Docker_Image": env.DockerImage,
		"Git_Name":     env.GitName,
	}
}
