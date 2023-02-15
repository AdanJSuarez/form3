package account

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client"
)

//go:generate mockery --inpackage --name=Client
//go:generate mockery --inpackage --name=StatusHandler

type Client interface {
	Get(accoutURL string) (*http.Response, error)
	Post(requestBody client.RequestBody) (*http.Response, error)
	Delete(value, parameterKey, parameterValue string) (*http.Response, error)
}

type StatusHandler interface {
	StatusCreated(response *http.Response) bool
	StatusOK(response *http.Response) bool
	StatusNoContent(response *http.Response) bool
	HandleError(response *http.Response) error
}
