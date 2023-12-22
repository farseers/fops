package model

import (
	"fops/domain/_/eumBuildStatus"
	"fops/domain/apps"
	"time"
)

// BuildPO 实体类
type BuildPO struct {
	Id            int64               `gorm:"primaryKey;comment:主键"`
	ClusterId     int64               `gorm:"not null;default:0;comment:集群信息"`
	BuildNumber   int                 `gorm:"type:int;not null;default:0;comment:构建号"`
	Status        eumBuildStatus.Enum `gorm:"type:tinyint;not null;default:0;comment:状态：0=未开始，1=构建中，2=完成"`
	IsSuccess     bool                `gorm:"size:1;not null;default:0;comment:是否成功"`
	CreateAt      time.Time           `gorm:"type:timestamp;size:6;not null;default:CURRENT_TIMESTAMP;comment:开始时间"`
	FinishAt      time.Time           `gorm:"type:timestamp;size:6;not null;default:CURRENT_TIMESTAMP;comment:完成时间"`
	BuildServerId int64               `gorm:"not null;default:0;comment:构建的服务端id"`
	Log           []string            `gorm:"type:text;json;not null;comment:构建日志"`
	Env           apps.EnvVO          `gorm:"type:text;json;not null;comment:环境变量"`
	AppName       string              `gorm:"size:32;not null;comment:应用名称"`
	Dockerfile    string              `gorm:"type:text;not null;comment:Dockerfile内容"`
	ShellScript   string              `gorm:"type:mediumtext;not null;comment:Shell脚本"`
}
