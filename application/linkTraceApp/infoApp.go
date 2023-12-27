package linkTraceApp

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/linkTrace/driver/clickhouse"
)

// Info 链路追踪日志详情
// @get info
func Info(traceId int64) collections.List[linkTrace_clickhouse.TraceContextPO] {
	lst := linkTrace_clickhouse.CHContext.TraceContext.Where("trace_id", traceId).Asc("start_ts").ToList()
	return lst
}
