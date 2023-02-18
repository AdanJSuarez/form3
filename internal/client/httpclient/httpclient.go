package httpclient

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// const (
// 	HOST_KEY              = "Host"
// 	DATE_KEY              = "Date"
// 	ACCEPT_KEY            = "Accept"
// 	ACCEPT_ENCODING_KEY   = "Accept-Encoding"
// 	CONTENT_TYPE_KEY      = "Content-Type"
// 	CONTENT_LENGTH_KEY    = "Content-Length"
// 	DIGEST_KEY            = "Digest"
// 	CONTENT_TYPE_VALUE    = "application/vnd.api+json"
// 	ACCEPT_ENCODING_VALUE = "gzip"
// )

const (
	timeout = 10 * time.Second
)

type HTTPClient struct {
	clientURL  url.URL
	httpClient httpClient
}

func New(clientURL url.URL) *HTTPClient {
	client := &http.Client{
		Timeout: timeout,
	}
	return &HTTPClient{
		clientURL:  clientURL,
		httpClient: client,
	}
}

func (c *HTTPClient) Get(request *http.Request) (*http.Response, error) {
	// request, err := c.request(GET, url, request.NewRequestHandler(nil))
	// if err != nil {
	// 	return nil, err
	// }

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) Post(request *http.Request) (*http.Response, error) {
	// request, err := c.request(POST, c.clientURL.String(), body)
	// if err != nil {
	// 	return nil, err
	// }

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *HTTPClient) Delete(request *http.Request) (*http.Response, error) {
	// request, err := c.request(DELETE, url, request.NewRequestHandler(nil))
	// if err != nil {
	// 	return nil, err
	// }

	// c.setQuery(request, parameterKey, parameterValue)

	response, err := c.doRequest(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// func (c *HTTPClient) request(method string, url string, requestBody RequestBody) (*http.Request, error) {
// 	request, err := http.NewRequest(method, url, requestBody.Body())
// 	if err != nil {
// 		return nil, err
// 	}

// 	c.addRequiredHeader(request)

// 	if requestBody.Body() != nil {
// 		c.addHeaderToRequestWithBody(request, requestBody.Size(), requestBody.Digest())
// 	}

// 	return request, nil
// }

func (c *HTTPClient) doRequest(request *http.Request) (*http.Response, error) {
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

// func (c *HTTPClient) setQuery(request *http.Request, parameterKey, parameterValue string) {
// 	query := request.URL.Query()
// 	query.Add(parameterKey, parameterValue)
// 	request.URL.RawQuery = query.Encode()
// }

// func (c *HTTPClient) addRequiredHeader(request *http.Request) {
// 	request.Header.Add(HOST_KEY, c.clientURL.Host)
// 	request.Header.Add(DATE_KEY, time.Now().Format(time.RFC1123))
// 	request.Header.Add(ACCEPT_KEY, CONTENT_TYPE_VALUE)
// 	request.Header.Add(ACCEPT_ENCODING_KEY, ACCEPT_ENCODING_VALUE)
// }

// func (c *HTTPClient) addHeaderToRequestWithBody(request *http.Request, size int, desire string) {
// 	request.Header.Add(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
// 	request.Header.Add(CONTENT_LENGTH_KEY, fmt.Sprint(size))
// 	request.Header.Add(DIGEST_KEY, desire)
// }
