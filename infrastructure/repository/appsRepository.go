package repository

import (
	"encoding/json"
	"fops/domain/apps"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseer-go/data"
	"github.com/farseer-go/mapper"
)

type appsRepository struct {
	// IRepository 通用的仓储接口
	data.IRepository[apps.DomainObject]
	buildRepository
	gitRepository
}

func (receiver *appsRepository) UpdateApp(do apps.DomainObject) error {
	po := mapper.Single[model.AppsPO](do)
	_, err := context.MysqlContext.Apps.Where("app_name = ?", po.AppName).Omit("app_name", "docker_ver", "cluster_ver").Update(po)
	return err
}

// UpdateDockerVer 修改镜像版本
func (receiver *appsRepository) UpdateDockerVer(appName string, dockerVer int) (int64, error) {
	return context.MysqlContext.Apps.Where("app_name = ?", appName).UpdateValue("docker_ver", dockerVer)
}

// UpdateClusterVer 修改集群的镜像版本
func (receiver *appsRepository) UpdateClusterVer(appName string, dicClusterVer map[int64]*apps.ClusterVerVO) (int64, error) {
	marshal, _ := json.Marshal(dicClusterVer)
	return context.MysqlContext.Apps.Where("app_name = ?", appName).UpdateValue("cluster_ver", string(marshal))
}
