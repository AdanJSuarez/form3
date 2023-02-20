package client

import (
	"net/http"
)

//go:generate mockery --inpackage --name=httpClient
//go:generate mockery --inpackage --name=requestHandler
//go:generate mockery --inpackage --name=errorHandler

type httpClient interface {
	SendRequest(request *http.Request) (*http.Response, error)
}

type requestHandler interface {
	Request(data interface{}, method, url, host string) (*http.Request, error)
	SetQuery(request *http.Request, parameterKey, parameterValue string)
}

type statusErrorHandler interface {
	StatusError(response *http.Response) (*http.Response, error)
}
