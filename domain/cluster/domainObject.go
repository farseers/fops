package cluster

// DomainObject 集群
type DomainObject struct {
	Id               int64  // 集群Id
	Name             string // 集群名称
	Ip               string // 集群地址
	K8sConnectConfig string // K8s连接配置
	DockerName       string // 仓库名称
	DockerHub        string // 托管地址
	DockerUserName   string // 账户名称
	DockerUserPwd    string // 账户密码
	DockerNetwork    string // Docker网络
}
