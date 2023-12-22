package device

import (
	"fops/domain/apps"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/utils/file"
	"path/filepath"
)

func RegisterCopyToDistDevice() {
	container.Register(func() apps.ICopyToDistDevice { return &copyToDistDevice{} })
}

type copyToDistDevice struct {
}

// Copy 将项目的源代码，复制到目标目录中。用于后续的打包
func (copyToDistDevice) Copy(lstGit []apps.GitEO, env apps.EnvVO, progress chan string) {
	progress <- "---------------------------------------------------------"
	for _, git := range lstGit {
		dest := filepath.Join(apps.DistRoot, git.GetRelativePath())
		progress <- "源文件" + git.GetAbsolutePath() + " 复制到 " + dest
		file.CopyFolder(git.GetAbsolutePath(), dest)
	}
}
