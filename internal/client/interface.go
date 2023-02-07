package client

import "net/http"

//go:generate mockery --inpackage --name=httpClient

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
