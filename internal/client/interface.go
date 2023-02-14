package client

import (
	"io"
	"net/http"
)

//go:generate mockery --inpackage --name=httpClient
//go:generate mockery --inpackage --name=statusHandler

type RequestBody interface {
	Body() io.ReadCloser
	Size() int
	Digest() string
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type statusHandler interface {
	HandleError(response *http.Response) error
}
