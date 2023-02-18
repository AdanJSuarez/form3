package client

import (
	"net/http"
)

//go:generate mockery --inpackage --name=httpClient
//go:generate mockery --inpackage --name=requestHandler
//go:generate mockery --inpackage --name=statusHandler

type httpClient interface {
	Get(request *http.Request) (*http.Response, error)
	Post(request *http.Request) (*http.Response, error)
	Delete(request *http.Request) (*http.Response, error)
}

type requestHandler interface {
	Request(data interface{}, method, url, host string) (*http.Request, error)
	SetQuery(request *http.Request, parameterKey, parameterValue string)
}

type statusHandler interface {
	StatusCreated(response *http.Response) bool
	StatusOK(response *http.Response) bool
	StatusNoContent(response *http.Response) bool
	HandleError(response *http.Response) error
}
