package request

type AddRequest struct {
	Name     string // Git名称
	Hub      string // git地址
	Branch   string // Git分支
	UserName string // 账户名称
	UserPwd  string // 账户密码
	Dir      string // 存储目录
}
