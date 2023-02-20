package httpclient

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

/*
# DEFAULT_TIMEOUT = 60 seconds
# MAX_RETRIES = 3
# retries = 0
#
# DO
#     TRY
#       IF retries > 0
#         WAIT for (1.5^retries * 500) milliseconds +- some jitter
#
#       status = makeCallToForm3(timeout:DEFAULT_TIMEOUT)
#
#       IF status = SUCCESS (2xx) or CONFLICT (409)
#           retry = false
#       ELSE IF status = THROTTLED (429) # You have reached your request limit and are being throttled
#           retry = true
#      ELSE IF status >= 500 # A temporary issue has occurred, all requests are idempotent and safe to retry
#           retry = true
#       ELSE # Another http response such as 400 bad request, client must fix request before retrying
#           retry = false
#       END IF
#     CATCH EXCEPTION
#       retry = true # connection timeout, connection dropped etc...
#     END TRY
#
#     retries = retries + 1
# WHILE (retry AND (retries <= MAX_RETRIES))
*/
const (
	timeout = 100 * time.Second
	// timeout          = 1 * time.Microsecond
	maxRetries       = 3
	maxJitter        = 500
	exponentialBase  = 1.5
	periodMultiplier = 500 * time.Millisecond
	nilRequest       = "nil request"
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

func (c *HTTPClient) SendRequest(request *http.Request) (*http.Response, error) {
	var retries float64 = 0
	var response *http.Response
	var err error

	for retries <= maxRetries {
		if retries > 0 {
			c.exponentialDelay(retries)
		}

		response, err = c.do(request)
		if err != nil && !os.IsTimeout(err) {
			return nil, err
		}

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
	period := time.Duration(math.Pow(exponentialBase, retries)) * periodMultiplier
	jitter := time.Duration(rand.Intn(maxJitter))
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
