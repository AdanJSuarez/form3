package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AdanJSuarez/form3/internal/client/requestbody"
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
	DIGEST_KEY         = "Digest"
	CONTENT_TYPE_VALUE = "application/vnd.api+json"
)

const (
	timeout = 10 * time.Second
)

type Client struct {
	clientURL  url.URL
	httpClient httpClient
}

func New(clientURL url.URL) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		clientURL:  clientURL,
		httpClient: client,
	}
}

func NewRequestBody(data interface{}) RequestBody {
	return requestbody.NewRequestBody(data)
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(GET, url, requestbody.NewRequestBody(nil))
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
	request, err := c.request(POST, c.clientURL.String(), body)
	if err != nil {
		return nil, err
	}

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Delete(value, parameterKey, parameterValue string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(DELETE, url, NewRequestBody(nil))
	if err != nil {
		return nil, err
	}

	c.setQuery(request, parameterKey, parameterValue)

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
		c.addHeaderToRequestWithBody(request, requestBody.Size(), requestBody.Digest())
	}

	return request, nil
}

func (c *Client) doRequest(request *http.Request) (*http.Response, error) {
	emptyResponse := &http.Response{}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return emptyResponse, err
	}
	if response == nil {
		return emptyResponse, fmt.Errorf("nil response")
	}
	return response, err
}

func (c *Client) joinValuesToURL(values ...string) (string, error) {
	url, err := url.JoinPath(c.clientURL.String(), values...)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *Client) setQuery(request *http.Request, parameterKey, parameterValue string) {
	query := request.URL.Query()
	query.Add(parameterKey, parameterValue)
	request.URL.RawQuery = query.Encode()
}

//TODO: Include header with "gzip"

func (c *Client) addRequiredHeader(request *http.Request) {
	request.Header.Add(HOST_KEY, c.clientURL.Host)
	request.Header.Add(DATE_KEY, time.Now().Format(time.RFC1123))
	request.Header.Add(ACCEPT_KEY, CONTENT_TYPE_VALUE)
}

func (c *Client) addHeaderToRequestWithBody(request *http.Request, size int, desire string) {
	request.Header.Add(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	request.Header.Add(CONTENT_LENGTH_KEY, fmt.Sprint(size))
	request.Header.Add(DIGEST_KEY, desire)
}
