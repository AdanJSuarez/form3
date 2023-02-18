package client

import (
	"net/http"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/client/request"
	"github.com/AdanJSuarez/form3/internal/client/statushandler"
)

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

type Client struct {
	clientURL     url.URL
	httpClient    httpClient
	statusHandler statusHandler
}

func New(clientURL url.URL) *Client {
	return &Client{
		clientURL:     clientURL,
		statusHandler: statushandler.NewStatusHandler(),
	}
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}

	requestHandler := request.NewRequestHandler(nil)
	request, err := requestHandler.Request(GET, url, c.clientURL.Host)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Get(request)
}

func (c *Client) Post(data interface{}) (*http.Response, error) {
	requestHandler := request.NewRequestHandler(data)
	request, err := requestHandler.Request(POST, c.clientURL.String(), c.clientURL.Host)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Post(request)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) Delete(value, parameterKey, parameterValue string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}
	requestHandler := request.NewRequestHandler(nil)
	request, err := requestHandler.Request(DELETE, url, c.clientURL.Host)
	if err != nil {
		return nil, err
	}
	requestHandler.SetQuery(request, parameterKey, parameterValue)
	return c.httpClient.Delete(request)
}

func (c *Client) StatusCreated(response *http.Response) bool {
	return c.statusHandler.StatusCreated(response)
}

func (c *Client) StatusOK(response *http.Response) bool {
	return c.statusHandler.StatusOK(response)
}

func (c *Client) StatusNoContent(response *http.Response) bool {
	return c.statusHandler.StatusNoContent(response)
}

func (c *Client) HandleError(response *http.Response) error {
	return c.statusHandler.HandleError(response)
}

func (c *Client) joinValuesToURL(values ...string) (string, error) {
	url, err := url.JoinPath(c.clientURL.String(), values...)
	if err != nil {
		return "", err
	}
	return url, nil
}
