package response

import (
	"github.com/farseer-go/collections"
	linkTraceCom "github.com/farseer-go/linkTrace"
)

type LinkTraceResponse struct {
	Entry linkTraceCom.TraceContext
	List  collections.List[LinkTraceVO]
}
