package device

import (
	"fops/domain/apps"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/utils/file"
)

func RegisterDirectoryDevice() {
	container.Register(func() apps.IDirectoryDevice { return &directoryDevice{} })
}

type directoryDevice struct {
}

// Check 检查目录
func (directoryDevice) Check(progress chan string) {
	progress <- "---------------------------------------------------------"
	progress <- "前置检查。"
	// 先删除之前编译的目标文件
	progress <- "先删除之前的目标文件。"

	file.ClearFile(apps.DistRoot)

	// 自动创建目录
	progress <- "自动创建目录。"

	if !file.IsExists(apps.FopsRoot) {
		file.CreateDir766(apps.FopsRoot)
	}
	if !file.IsExists(apps.NpmModulesRoot) {
		file.CreateDir766(apps.NpmModulesRoot)
	}
	if !file.IsExists(apps.DistRoot) {
		file.CreateDir766(apps.DistRoot)
	}
	if !file.IsExists(apps.KubeRoot) {
		file.CreateDir766(apps.KubeRoot)
	}
	if !file.IsExists(apps.ShellRoot) {
		file.CreateDir766(apps.ShellRoot)
	}
	if !file.IsExists(apps.GitRoot) {
		file.CreateDir766(apps.GitRoot)
	}
	progress <- "前置检查通过。"
}
