package request

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/flog"
)

type UploadRequest struct {
	List collections.List[flog.LogData]
}
