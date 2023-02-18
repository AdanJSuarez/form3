package client

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client/requestbody"
)

type httpClient interface {
	Get(url string) (*http.Response, error)
	Post(body *requestbody.RequestBody) (*http.Response, error)
	Delete(url, parameterKey, parameterValue string) (*http.Response, error)
}

type statusHandler interface {
	StatusCreated(response *http.Response) bool
	StatusOK(response *http.Response) bool
	StatusNoContent(response *http.Response) bool
	HandleError(response *http.Response) error
}
