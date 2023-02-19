package account

import (
	"net/http"
)

//go:generate mockery --inpackage --name=Client

type Client interface {
	Get(accoutID string) (*http.Response, error)
	Post(data interface{}) (*http.Response, error)
	Delete(accountID, parameterKey, parameterValue string) (*http.Response, error)
	// StatusCreated(response *http.Response) bool
	// StatusOK(response *http.Response) bool
	// StatusNoContent(response *http.Response) bool
	// HandleError(response *http.Response) error
}
