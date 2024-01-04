package apps

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/fs/exception"
	"github.com/farseer-go/utils/str"
	"net/url"
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
	IsApp    bool              // 是否为应用
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

// GetAuthHub 获取带账号密码的地址
func (receiver *GitEO) GetAuthHub() string {
	parsedURL, err := url.Parse(receiver.Hub)
	exception.ThrowRefuseExceptionfBool(err != nil, "解析 URL 失败:%s", err)

	// 设置用户名和密码
	parsedURL.User = url.UserPassword(receiver.UserName, receiver.UserPwd)

	return parsedURL.String()
}
