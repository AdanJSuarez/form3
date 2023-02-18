package httpclient

import (
	"net/http"
)

//go:generate mockery --inpackage --name=httpClient

// type RequestBody interface {
// 	Body() io.ReadCloser
// 	Size() int
// 	Digest() string
// }

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
