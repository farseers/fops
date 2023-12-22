package model

import "time"

type GitPO struct {
	Id       int64     `gorm:"primaryKey;comment:主键"`
	Name     string    `gorm:"size:32;not null;comment:Git名称"`
	Hub      string    `gorm:"size:256;not null;comment:托管地址"`
	Branch   string    `gorm:"size:64;not null;comment:Git分支"`
	UserName string    `gorm:"size:32;not null;comment:账户名称"`
	UserPwd  string    `gorm:"size:64;not null;comment:账户密码"`
	Dir      string    `gorm:"size:64;not null;comment:存储目录"`
	PullAt   time.Time `gorm:"type:timestamp;size:6;not null;default:CURRENT_TIMESTAMP;comment:拉取时间"`
}
