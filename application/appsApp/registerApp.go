// @area /apps/
package appsApp

import (
	"fops/application/appsApp/request"
	"fops/domain/register"
	"github.com/farseer-go/fs/dateTime"
	"github.com/farseer-go/mapper"
)

// Register 上传当前应用信息
// @post register
func Register(req request.RegisterAppRequest, registerRepository register.Repository) {
	do := mapper.Single[register.DomainObject](req)
	do.ActivateAt = dateTime.Now()
	_ = registerRepository.Add(do)
}
