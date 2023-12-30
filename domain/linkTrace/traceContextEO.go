package linkTrace

import (
	"github.com/farseer-go/fs/trace"
	"github.com/farseer-go/linkTrace"
	"github.com/farseer-go/linkTrace/eumTraceType"
	"time"
)

type TraceContextEO struct {
	TraceId       int64                // 上下文ID
	AppId         int64                // 应用ID
	AppName       string               // 应用名称
	AppIp         string               // 应用IP
	ParentAppName string               // 上游应用
	StartTs       int64                // 调用开始时间戳（微秒）
	EndTs         int64                // 调用结束时间戳（微秒）
	UseTs         time.Duration        // 总共使用时间微秒
	TraceType     eumTraceType.Enum    // 状态码
	List          []any                // 调用的上下文
	ignore        bool                 // 忽略这次的链路追踪
	Exception     trace.ExceptionStack // 异常信息
	linkTrace.WebContext
	linkTrace.ConsumerContext
	linkTrace.TaskContext
	linkTrace.WatchKeyContext
}
