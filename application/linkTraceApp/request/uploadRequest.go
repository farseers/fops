package request

import (
	"fops/domain/linkTrace"
	"github.com/farseer-go/collections"
)

type UploadRequest struct {
	List collections.List[linkTrace.TraceContextEO]
}
