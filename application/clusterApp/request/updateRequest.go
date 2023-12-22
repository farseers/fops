package request

type UpdateRequest struct {
	Id             int64  // 集群Id
	Name           string // 集群名称
	Ip             string // 集群地址
	DockerName     string // 仓库名称
	DockerHub      string // 托管地址
	DockerUserName string // 账户名称
	DockerUserPwd  string // 账户密码
}
