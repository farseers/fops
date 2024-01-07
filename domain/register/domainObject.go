package register

import "github.com/farseer-go/fs/dateTime"

// DomainObject 注册应用
type DomainObject struct {
	AppName     string            // 应用名称
	AppId       string            // 应用ID
	AppIp       string            // 应用IP
	HostName    string            // 主机名称
	ProcessId   int               // 进程Id
	StartupAt   dateTime.DateTime // 应用启动时间
	ActivateAt  dateTime.DateTime // 活动时间
	CpuUsage    float64           // CPU百分比
	MemoryUsage float64           // 内存百分比
}
