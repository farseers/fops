package device

import (
	"context"
	"fmt"
	"fops/domain/apps"
	"fops/domain/apps/event"
	"fops/domain/cluster"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/utils/exec"
	"github.com/farseer-go/utils/file"
	"github.com/farseer-go/utils/str"
	"os"
	"strconv"
)

func RegisterDockerDevice() {
	container.Register(func() apps.IDockerDevice { return &dockerDevice{} })
}

type dockerDevice struct {
}

func (dockerDevice) GetDockerHub(dockerHubAddress string) string {
	var dockerHub = "localhost"
	if dockerHubAddress != "" {
		dockerHub = dockerHubAddress
		dockerHub = str.CutRight(dockerHub, "/")
	}
	return dockerHub
}

func (device dockerDevice) GetDockerImage(dockerHubAddress string, projectName string, buildNumber int) string {
	return device.GetDockerHub(dockerHubAddress) + "/" + projectName + ":" + strconv.Itoa(buildNumber)
}

func (dockerDevice) Login(dockerHub string, loginName string, loginPwd string, progress chan string, env apps.EnvVO, ctx context.Context) bool {
	progress <- "---------------------------------------------------------"
	if dockerHub != "" && loginName != "" {
		var result = exec.RunShellContext(ctx, "docker login "+dockerHub+" -u "+loginName+" -p "+loginPwd, progress, env.ToMap(), "")
		if result != 0 {
			progress <- "镜像仓库登陆失败。"
			return false
		}
	}

	progress <- "镜像仓库登陆成功。"
	return true
}

func (dockerDevice) ExistsDockerfile(dockerfilePath string) bool {
	return file.IsExists(dockerfilePath)
}

func (dockerDevice) CreateDockerfile(projectName string, dockerfileContent string, ctx context.Context) {
	if file.IsExists(apps.DockerfilePath) {
		_ = os.RemoveAll(apps.DockerfilePath)
	}
	file.WriteString(apps.DockerfilePath, dockerfileContent)
}

func (dockerDevice) Build(env apps.EnvVO, progress chan string, ctx context.Context) bool {
	progress <- "---------------------------------------------------------"
	progress <- "开始镜像打包。"

	// 打包
	var result = exec.RunShellContext(ctx, "docker build -t "+env.DockerImage+" --network=host -f "+apps.DockerfilePath+" "+apps.DistRoot, progress, env.ToMap(), apps.DistRoot)
	if result == 0 {
		progress <- "镜像打包完成。"
	} else {
		progress <- "镜像打包出错了。"
	}
	return result == 0
}

func (dockerDevice) Push(env apps.EnvVO, progress chan string, ctx context.Context) bool {
	defer func() {
		// 上传完后，删除本地镜像
		exec.RunShellContext(ctx, "docker rmi "+env.DockerImage, progress, env.ToMap(), "")
	}()

	// 上传
	var result = exec.RunShellContext(ctx, "docker push "+env.DockerImage, progress, env.ToMap(), "")

	if result == 0 {
		progress <- "镜像上传完成。"

		// 上传成功后，需要更新项目中的镜像版本属性
		event.DockerPushedEvent{
			BuildNumber: env.BuildNumber,
			AppName:     env.AppName,
		}.PublishEvent()
		return true
	}

	progress <- "镜像上传出错了。"
	return false
}

func (dockerDevice) SetImages(cluster cluster.DomainObject, appName string, dockerImages string, progress chan string, ctx context.Context) bool {
	progress <- "---------------------------------------------------------"
	progress <- "开始更新Docker Swarm的镜像版本。"

	var exitCode = exec.RunShellContext(ctx, fmt.Sprintf("docker service update --image %s %s", dockerImages, appName), progress, nil, "")
	if exitCode != 0 {
		progress <- "Docker Swarm更新镜像失败。"
		return false
	}
	progress <- "Docker Swarm更新镜像版本完成。"
	return true
}
