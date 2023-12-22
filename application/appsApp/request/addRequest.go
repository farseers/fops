package request

import "github.com/farseer-go/collections"

type AddRequest struct {
	AppName       string                  `validate:"required" label:"应用名称"` // 应用名称
	AppId         string                  // 应用ID（链路追踪）
	ShellScript   string                  // Shell脚本
	AppGit        int64                   // 应用的源代码
	FrameworkGits collections.List[int64] // 依赖的框架源代码
}
