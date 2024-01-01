package model

import (
	"github.com/farseer-go/fs/core/eumLogLevel"
	"github.com/farseer-go/fs/dateTime"
)

type LogDataPO struct {
	Id        int64             `gorm:"not null;default:0;comment:主键ID"`
	AppName   string            `gorm:"not null;default:'';comment:应用名称"`
	AppIp     string            `gorm:"not null;default:'';comment:应用IP"`
	AppId     int64             `gorm:"not null;default:0;comment:应用ID"`
	TraceId   int64             `gorm:"not null;default:0;comment:上下文ID"`
	Component string            `gorm:"not null;default:'';comment:组件名称"`
	LogLevel  eumLogLevel.Enum  `gorm:"not null;comment:日志等级"`
	Content   string            `gorm:"not null;default:'';comment:日志内容"`
	CreateAt  dateTime.DateTime `gorm:"type:DateTime64(3);not null;comment:发生时间"`
}
