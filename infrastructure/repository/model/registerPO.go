package model

import (
	"github.com/farseer-go/fs/dateTime"
	"github.com/shopspring/decimal"
)

type RegisterPO struct {
	Id          int64             `gorm:"primaryKey;comment:主键"`
	AppName     string            `gorm:"size:32;not null;comment:应用名称"`
	AppId       string            `gorm:"not null;default:'';comment:应用ID"`
	AppIp       string            `gorm:"size:32;not null;comment:应用IP"`
	HostName    string            `gorm:"size:32;not null;comment:主机名称"`
	ProcessId   int               `gorm:"type:int;not null;default:0;comment:进程Id"`
	StartupAt   dateTime.DateTime `gorm:"type:timestamp;size:6;not null;default:CURRENT_TIMESTAMP;comment:应用启动时间"`
	ActivateAt  dateTime.DateTime `gorm:"type:timestamp;size:6;index:idx_activate_at;not null;default:CURRENT_TIMESTAMP;comment:活动时间"`
	CpuUsage    decimal.Decimal   `gorm:"type:Decimal(6,4);not null;default:0;comment:CPU百分比"`
	MemoryUsage decimal.Decimal   `gorm:"type:Decimal(6,4);not null;default:0;comment:内存百分比"`
}
