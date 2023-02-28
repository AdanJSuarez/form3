package account

import (
	"net/http"
	"net/url"
)

//go:generate mockery --inpackage --name=Client
//go:generate mockery --inpackage --name=Configuration
type Client interface {
	Get(accountID string) (*http.Response, error)
	Post(data interface{}) (*http.Response, error)
	Delete(accountID, parameterKey, parameterValue string) (*http.Response, error)
}

type Configuration interface {
	BaseURL() *url.URL
	AccountPath() string
}
