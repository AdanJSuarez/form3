package form3

import url "net/url"

//go:generate mockery --inpackage --name=Configuration

type Configuration interface {
	BaseURL() *url.URL
	AccountPath() string
	InitializeByValue(rawBaseURL, accountPath string) error
	InitializeByEnv() error
}
