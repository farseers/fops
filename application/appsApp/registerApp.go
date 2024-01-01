// @area /apps/
package appsApp

import (
	"fmt"
	"fops/application/appsApp/request"
	"github.com/farseer-go/fs/flog"
)

// Register 上传当前应用信息
// @post register
func Register(req request.RegisterAppRequest) {
	flog.Info(fmt.Sprintf("%v+", req))
}
