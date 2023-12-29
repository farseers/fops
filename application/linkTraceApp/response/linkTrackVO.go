package response

import "github.com/farseer-go/linkTrace/driver/clickhouse"

var RgbaList = []string{"95,184,120,0.4", "65,105,225,0.4", "219,112,147,0.4", "128,0,128,0.4", "153,50,204,0.4", "123,104,238,0.4", "119,136,153,0.4", "70,130,180,0.4", "0,139,139,0.4", "34,139,34,0.4", "128,128,0,0.4", "238,232,170,0.4", "218,165,32,0.4", "255,165,0,0.4", "255,140,0,0.4", "210,105,30,0.4"}

type LinkTraceVO struct {
	Rgba      string // 背景颜色
	AppId     int64  // 应用实例ID
	AppIp     string // 应用IP
	AppName   string // 应用名称
	StartTs   int64  // 计算服务的相对调用时间，第一个服务从0ms计算
	UseTs     int64  // 使用时间
	Caption   string // 标题
	Desc      string // tips描述
	Exception *linkTrace_clickhouse.ExceptionStackPO
}
