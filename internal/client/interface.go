package client

import (
	"net/http"
)

type httpClient interface {
	Get(request *http.Request) (*http.Response, error)
	Post(request *http.Request) (*http.Response, error)
	Delete(request *http.Request) (*http.Response, error)
}

type statusHandler interface {
	StatusCreated(response *http.Response) bool
	StatusOK(response *http.Response) bool
	StatusNoContent(response *http.Response) bool
	HandleError(response *http.Response) error
}
