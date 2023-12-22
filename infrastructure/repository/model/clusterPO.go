package model

// ClusterPO 集群
type ClusterPO struct {
	Id               int64  `gorm:"primaryKey;comment:主键"`
	Name             string `gorm:"size:32;not null;comment:集群名称"`
	Ip               string `gorm:"size:64;not null;comment:集群地址"`
	K8sConnectConfig string `gorm:"size:2048;not null;comment:K8s连接配置"`
	DockerName       string `gorm:"size:64;not null;comment:仓库名称"`
	DockerHub        string `gorm:"size:256;not null;comment:托管地址"`
	DockerUserName   string `gorm:"size:32;not null;comment:账户名称"`
	DockerUserPwd    string `gorm:"size:64;not null;comment:账户密码"`
}
