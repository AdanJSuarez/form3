package httpclient

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultTimeout = 10 * time.Second
	maxRetries     = 3
	// maxJitter set small to avoid test timeout during retry. Set bigger for production
	maxJitter       = 100
	exponentialBase = 1.5
	// periodMultiplier set small to avoid test timeout during retry. Set bigger for production
	periodMultiplier = 100 * time.Millisecond
	nilRequest       = "nil request"
)

type HTTPClient struct {
	httpClient httpClient
}

func New() *HTTPClient {
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	return &HTTPClient{
		httpClient: client,
	}
}

func (c *HTTPClient) SendRequest(request *http.Request) (*http.Response, error) {
	var retries float64 = 0
	var response *http.Response
	var err error

	for retries <= maxRetries {
		if retries > 0 {
			c.exponentialDelay(retries)
		}

		response, err = c.do(request)

		if !c.needRetry(response) {
			return response, nil
		}
		retries++
	}

	return response, err
}

func (c *HTTPClient) needRetry(response *http.Response) bool {
	return response == nil || response.StatusCode >= http.StatusTooManyRequests
}

func (c *HTTPClient) exponentialDelay(retries float64) {
	rand.Seed(time.Now().UnixNano())
	period := time.Duration(math.Pow(exponentialBase, retries)) * periodMultiplier
	jitter := time.Duration(rand.Intn(maxJitter)) * time.Millisecond
	time.Sleep(period + jitter)
}

func (c *HTTPClient) do(request *http.Request) (*http.Response, error) {
	if request == nil {
		return nil, fmt.Errorf(nilRequest)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, err
}
