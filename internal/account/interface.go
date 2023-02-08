package account

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client"
)

//go:generate mockery --inpackage --name=Client

type Client interface {
	Get(accoutURL string) (*http.Response, error)
	Post(requestBody client.RequestBody) (*http.Response, error)
	Delete(value, paramKey, paramValue string) (*http.Response, error)
}
