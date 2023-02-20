package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	dataTest = "{data: {moreData: 3}}"
)

var (
	httpClientTest      *HTTPClient
	mockHTTPClient      *mockHttpClient
	dataBytesMarshal, _ = json.Marshal(dataTest)
	requestTest         = &http.Request{}
	responseGetTest     = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataBytesMarshal)),
	}
	responseTooManyRequestTest = http.Response{
		Status:     "429 Too Many Request",
		StatusCode: http.StatusTooManyRequests,
		Body:       io.NopCloser(bytes.NewBuffer(dataBytesMarshal)),
	}
	responseNotFoundTest = http.Response{
		Status:     "404 Not Found",
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	responseInternalServerErrorTest = http.Response{
		Status:     "500 Internal Server Error",
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	responseBadGatewayErrorTest = http.Response{
		Status:     "502 Bad Gateway",
		StatusCode: http.StatusBadGateway,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	responseServiceUnavailableErrorTest = http.Response{
		Status:     "503 Service Unavailable",
		StatusCode: http.StatusServiceUnavailable,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	responseGatewayTimeoutErrorTest = http.Response{
		Status:     "504 Gateway Timeout",
		StatusCode: http.StatusGatewayTimeout,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
)

type timeoutError struct {
	error
}

func (e timeoutError) Timeout() bool {
	return true
}
func (e timeoutError) Temporary() bool {
	return true
}
func (e timeoutError) Error() string {
	return "timeout error super fake"
}

type TSHTTPClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSHTTPClient))
}

func (ts *TSHTTPClient) BeforeTest(_, _ string) {
	httpClientTest = New()
	ts.IsType(new(HTTPClient), httpClientTest)
	mockHTTPClient = newMockHttpClient(ts.T())
	httpClientTest.httpClient = mockHTTPClient
}

func (ts *TSHTTPClient) TestValidRequest() {
	mockHTTPClient.On("Do", mock.Anything).Return(&responseGetTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSHTTPClient) TestReturnErr() {
	mockHTTPClient.On("Do", mock.Anything).Return(nil, fmt.Errorf("fakeError"))
	response, err := httpClientTest.SendRequest(requestTest)
	ts.ErrorContains(err, "fakeError")
	ts.Nil(response)
}
func (ts *TSHTTPClient) TestTimeoutError() {
	errTest := timeoutError{}
	mockHTTPClient.On("Do", mock.Anything).Return(nil, errTest)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.True(os.IsTimeout(err))
	ts.Nil(response)
}
func (ts *TSHTTPClient) TestNilRequest() {
	response, err := httpClientTest.SendRequest(nil)
	ts.ErrorContains(err, "nil request")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestResponseNotFound() {
	mockHTTPClient.On("Do", mock.Anything).Return(&responseNotFoundTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(404, response.StatusCode)
}

func (ts *TSHTTPClient) TestTooManyRequest() {
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
	mockHTTPClient.On("Do", mock.Anything).Return(&responseTooManyRequestTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseTooManyRequestTest, response)
}

func (ts *TSHTTPClient) TestInternalServerError() {
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
	mockHTTPClient.On("Do", mock.Anything).Return(&responseInternalServerErrorTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseInternalServerErrorTest, response)
}

func (ts *TSHTTPClient) TestBadGatewayError() {
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
	mockHTTPClient.On("Do", mock.Anything).Return(&responseBadGatewayErrorTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseBadGatewayErrorTest, response)
}
func (ts *TSHTTPClient) TestServiceUnavailableError() {
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
	mockHTTPClient.On("Do", mock.Anything).Return(&responseServiceUnavailableErrorTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseServiceUnavailableErrorTest, response)
}

func (ts *TSHTTPClient) TestGatewayTimeoutError() {
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
	mockHTTPClient.On("Do", mock.Anything).Return(&responseGatewayTimeoutErrorTest, nil)
	response, err := httpClientTest.SendRequest(requestTest)
	ts.NoError(err)
	ts.Equal(&responseGatewayTimeoutErrorTest, response)
}
