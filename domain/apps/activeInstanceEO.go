package apps

import "github.com/farseer-go/fs/dateTime"

// ActiveInstanceEO 正在运行的实例
type ActiveInstanceEO struct {
	AppId       string            // 应用ID
	AppIp       string            // 应用IP
	HostName    string            // 主机名称
	ProcessId   int               // 进程Id
	StartupAt   dateTime.DateTime // 应用启动时间
	CpuUsage    float64           // CPU百分比
	MemoryUsage float64           // 内存百分比
}
