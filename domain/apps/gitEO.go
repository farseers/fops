package apps

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/utils/str"
	"path/filepath"
	"strings"
)

// GitEO git仓库
type GitEO struct {
	Id       int               // 主键
	Name     string            // Git名称
	Hub      string            // git地址
	Branch   string            // Git分支
	UserName string            // 账户名称
	UserPwd  string            // 账户密码
	Dir      string            // 存储目录
	PullAt   dateTime.DateTime // 拉取时间
}

func (receiver *GitEO) IsNil() bool {
	return receiver.Id == 0
}

// GetAbsolutePath 获取git存储的绝对路径
func (receiver *GitEO) GetAbsolutePath() string {
	return GitRoot + receiver.GetRelativePath()
}

// GetRelativePath 获取git存储的相对路径
func (receiver *GitEO) GetRelativePath() string {
	if receiver.Dir == "" || receiver.Dir == "/" {
		receiver.Dir = receiver.GetName()
	}
	// 移除前后/
	receiver.Dir = strings.TrimPrefix(receiver.Dir, "/")
	receiver.Dir = strings.TrimSuffix(receiver.Dir, "/")
	return receiver.Dir + "/"
}

func (receiver *GitEO) GetName() string {
	git := filepath.Base(receiver.Hub)
	return str.CutRight(git, ".git")
}
