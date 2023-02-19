package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

const (
	//TODO: timeout = 10 * time.Second
	timeout = 1 * time.Microsecond
)

type HTTPClient struct {
	httpClient httpClient
}

func New() *HTTPClient {
	client := &http.Client{
		Timeout: timeout,
	}
	return &HTTPClient{
		httpClient: client,
	}
}

func (c *HTTPClient) Get(request *http.Request) (*http.Response, error) {
	response, err := c.do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) Post(request *http.Request) (*http.Response, error) {
	response, err := c.do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) Delete(request *http.Request) (*http.Response, error) {
	response, err := c.do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) do(request *http.Request) (*http.Response, error) {
	emptyResponse := &http.Response{}
	if request == nil {
		return emptyResponse, fmt.Errorf("nil request")
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return emptyResponse, err
	}
	if response == nil {
		return emptyResponse, fmt.Errorf("nil response")
	}
	return response, err
}
