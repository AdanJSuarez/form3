package client

import (
	"net/http"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/client/errorhandler"
	"github.com/AdanJSuarez/form3/internal/client/httpclient"
	"github.com/AdanJSuarez/form3/internal/client/request"
)

type Client struct {
	clientURL      url.URL
	httpClient     httpClient
	requestHandler requestHandler
	errorHandler   errorHandler
}

func New(clientURL url.URL) *Client {
	return &Client{
		clientURL:      clientURL,
		httpClient:     httpclient.New(),
		requestHandler: request.NewRequestHandler(),
		errorHandler:   errorhandler.NewErrorHandler(),
	}
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.requestHandler.Request(nil, http.MethodGet, url, c.clientURL.Host)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Get(request)
	if err != nil {
		return c.errorHandler.Error(request, err)
	}

	if !c.statusOK(response) {
		return c.errorHandler.StatusError(request, response)
	}

	return response, nil
}

func (c *Client) Post(data interface{}) (*http.Response, error) {
	request, err := c.requestHandler.Request(data, http.MethodPost, c.clientURL.String(), c.clientURL.Host)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Post(request)
	if err != nil {
		return c.errorHandler.Error(request, err)
	}

	if !c.statusCreated(response) {
		return c.errorHandler.StatusError(request, response)
	}

	return response, nil
}

func (c *Client) Delete(value, parameterKey, parameterValue string) (*http.Response, error) {
	url, err := c.joinValuesToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.requestHandler.Request(nil, http.MethodDelete, url, c.clientURL.Host)
	if err != nil {
		return nil, err
	}
	c.requestHandler.SetQuery(request, parameterKey, parameterValue)

	response, err := c.httpClient.Delete(request)
	if err != nil {
		return c.errorHandler.Error(request, err)
	}

	if !c.statusNoContent(response) {
		return c.errorHandler.StatusError(request, response)
	}

	return response, nil
}

func (c *Client) statusCreated(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusCreated
}

func (c *Client) statusOK(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}

func (c *Client) statusNoContent(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusNoContent
}

// func (c *Client) HandleError(response *http.Response) error {
// 	return c.errorHandler.HandleStatusError(response)
// }

func (c *Client) joinValuesToURL(values ...string) (string, error) {
	url, err := url.JoinPath(c.clientURL.String(), values...)
	if err != nil {
		return "", err
	}
	return url, nil
}
