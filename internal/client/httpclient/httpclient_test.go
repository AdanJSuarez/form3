package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/AdanJSuarez/form3/internal/client/request"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURLTest = "https://api.fakeaddress.tech/fake"
	valueURL       = "/v1/organisation/accounts"
	idTest         = "020cf7d8-01b9-461d-89d4-89d57fd0d998"
)

var (
	clientURLTest       *url.URL
	httpClientTest      *HTTPClient
	dataTest            = "{data: {moreData: 3}}"
	dataBytesMarshal, _ = json.Marshal(dataTest)
	responseGetTest     = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataBytesMarshal)),
	}
	responsePostTest = http.Response{
		Status:     "201 Created",
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer(dataBytesMarshal)),
	}
	responseNotFoundTest = http.Response{
		Status:     "404 Not Found",
		StatusCode: 404,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	responseDelete = http.Response{
		Status:     "204 Delete",
		StatusCode: 204,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
	reqBodyTest = request.NewRequestHandler(dataTest)
)

type TSHTTPClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSHTTPClient))
}

func (ts *TSHTTPClient) BeforeTest(_, _ string) {
	clientURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	httpClientTest = New(*clientURLTest)
	ts.IsType(new(HTTPClient), httpClientTest)
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseGetTest, nil)
	httpClientTest.httpClient = mockHTTPClient
}

func (ts *TSHTTPClient) TestValidGetWithData() {
	response, err := httpClientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSHTTPClient) TestValidGetWithEmptyData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	httpClientTest.httpClient = mockHTTPClient

	response, err := httpClientTest.Get("")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSHTTPClient) TestErrorGetWithData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	httpClientTest.httpClient = mockHTTPClient

	response, err := httpClientTest.Get(idTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestValidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responsePostTest, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Post(reqBodyTest)
	ts.NoError(err)
	ts.NotNil(response)
	ts.Equal(201, response.StatusCode)
}

func (ts *TSHTTPClient) TestInvalidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Post(reqBodyTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestValidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSHTTPClient) TestValidDeleteEmptyValue() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("", "", "")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSHTTPClient) TestValidDeleteNotFound() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
	ts.Equal(404, response.StatusCode)
}

func (ts *TSHTTPClient) TestInvalidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestValidRequest() {
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal("api.fakeaddress.tech", request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Equal(CONTENT_TYPE_VALUE, request.Header.Get(CONTENT_TYPE_KEY))
	ts.Equal(fmt.Sprint(len(dataBytesMarshal)), request.Header.Get(CONTENT_LENGTH_KEY))
	ts.NotEmpty(request.Header.Get(DIGEST_KEY))
}

func (ts *TSHTTPClient) TestValidRequestNotBody() {
	rbTest := request.NewRequestHandler(nil)
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), rbTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal("api.fakeaddress.tech", request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Equal("", request.Header.Get(CONTENT_TYPE_KEY))
	ts.Equal("", request.Header.Get(CONTENT_LENGTH_KEY))
	ts.Equal("", request.Header.Get(DIGEST_KEY))
}

func (ts *TSHTTPClient) TestEmptyResponseAndRequestError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	httpClientTest.httpClient = mockHTTPClient
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
	ts.NoError(err)
	response, err := httpClientTest.doRequest(request)
	ts.ErrorContains(err, "fake error")
	ts.NotNil(response)
}

func (ts *TSHTTPClient) TestEmptyResponseAndNilResponseError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	httpClientTest.httpClient = mockHTTPClient
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
	ts.NoError(err)
	response, err := httpClientTest.doRequest(request)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSHTTPClient) TestEmptyResponseAndNilRequest() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.doRequest(nil)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSHTTPClient) TestQuery() {
	request, err := httpClientTest.request(DELETE, httpClientTest.clientURL.String(),
		request.NewRequestHandler(nil))
	ts.NoError(err)
	httpClientTest.setQuery(request, "fakeKey", "fakeValue")
	ts.Equal("fakeKey=fakeValue", request.URL.RawQuery)
}
