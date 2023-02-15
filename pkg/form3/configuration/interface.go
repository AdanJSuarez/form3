package configuration

import "net/url"

//go:generate mockery --inpackage --name=Configuration

type Configuration interface {
	BaseURL() *url.URL
	AccountPath() string
	InitializeByValue(rawBaseURL, accountPath string) error
	InitializeByYaml() error
	InitializeByEnv() error
}
