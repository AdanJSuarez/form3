package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

const (
	HOST_KEY           = "Host"
	DATE_KEY           = "Date"
	ACCEPT_KEY         = "Accept"
	CONTENT_TYPE_KEY   = "Content-Type"
	CONTENT_LENGTH_KEY = "Content-Length"
	CONTENT_TYPE_VALUE = "application/vnd.api+json"
)

const (
	timeout     = 10 * time.Second
	emptySize   = 0
	defaultHost = "api.form3.tech"
)

type Client struct {
	baseURL   *url.URL
	stringURL string
	client    httpClient
}

func New(baseURL *url.URL, valuesURL ...string) *Client {
	stringURL, _ := url.JoinPath(baseURL.String(), valuesURL...)
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		baseURL:   baseURL,
		stringURL: stringURL,
		client:    client,
	}
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := c.joinValueToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(GET, url, NewRequestBody(nil, emptySize))
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Post(body RequestBody) (*http.Response, error) {
	request, err := c.request(POST, c.stringURL, body)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Delete(value string) (*http.Response, error) {
	url, err := c.joinValueToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(DELETE, url, NewRequestBody(nil, emptySize))
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) request(method string, url string, requestBody RequestBody) (*http.Request, error) {
	request, err := http.NewRequest(method, url, requestBody.Body())
	if err != nil {
		return nil, err
	}

	c.addRequiredHeader(request)

	if requestBody.Body() != nil {
		c.addHeaderToRequestWithBody(request, requestBody.Size())
	}

	return request, nil
}

func (c *Client) doRequest(request *http.Request) (*http.Response, error) {
	emptyResponse := &http.Response{}
	response, err := c.client.Do(request)
	if err != nil {
		return emptyResponse, err
	}
	if response == nil {
		return emptyResponse, fmt.Errorf("nil response")
	}
	return response, err
}

func (c *Client) joinValueToURL(value string) (string, error) {
	url, err := url.JoinPath(c.stringURL, value)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *Client) addRequiredHeader(request *http.Request) {
	request.Header.Add(HOST_KEY, c.baseURL.Host)
	request.Header.Add(DATE_KEY, time.Now().String())
	request.Header.Add(ACCEPT_KEY, CONTENT_TYPE_VALUE)
}

func (c *Client) addHeaderToRequestWithBody(request *http.Request, size int) {
	request.Header.Add(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	request.Header.Add(CONTENT_LENGTH_KEY, fmt.Sprint(size))
}
