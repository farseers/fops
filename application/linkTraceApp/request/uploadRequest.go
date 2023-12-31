package request

import (
	"github.com/farseer-go/collections"
	linkTraceCom "github.com/farseer-go/linkTrace"
)

type UploadRequest struct {
	List collections.List[linkTraceCom.TraceContext]
}
