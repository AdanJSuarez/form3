package httpclient

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultTimeout  = 30 * time.Second
	exponentialBase = 1.5
	maxJitter       = 10
	nilRequest      = "nil request"
	maxConnections  = 100
)
const (
	maxRetries       float64 = 3
	periodMultiplier float64 = 1
)

var timeframe = time.Second

type HTTPClient struct {
	httpClient httpClient
}

func New() *HTTPClient {
	hc := &HTTPClient{}
	client := &http.Client{
		Timeout:   defaultTimeout,
		Transport: hc.basicTransport(),
	}
	hc.httpClient = client
	return hc
}

func (c *HTTPClient) SendRequest(request *http.Request) (*http.Response, error) {
	var retries float64 = 0
	var response *http.Response
	var err error

	for retries <= maxRetries {
		if c.hasRetried(retries) {
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

func (c *HTTPClient) basicTransport() *http.Transport {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = maxConnections
	transport.MaxIdleConnsPerHost = maxConnections
	transport.MaxConnsPerHost = maxConnections

	return transport
}

func (c *HTTPClient) hasRetried(retries float64) bool {
	return retries > 0
}

func (c *HTTPClient) needRetry(response *http.Response) bool {
	return response == nil || response.StatusCode >= http.StatusTooManyRequests
}

func (c *HTTPClient) exponentialDelay(retries float64) {
	rand.Seed(time.Now().UnixNano())
	period := time.Duration(math.Pow(exponentialBase, retries)*periodMultiplier) * timeframe
	jitter := time.Duration(rand.Intn(maxJitter)) * timeframe
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
