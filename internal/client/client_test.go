package client

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURLTest = "https://api.fakeaddress.tech/fake"
	valueURL       = "/v1/organisation/accounts"
	idTest         = "020cf7d8-01b9-461d-89d4-89d57fd0d998"
)

var (
	clientURLTest     *url.URL
	clientTest        *Client
	httpClientMock    *mockHttpClient
	statusHandlerMock *mockStatusHandler
)

type TSClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSClient))
}

func (ts *TSClient) BeforeTest(_, _ string) {
	clientURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	clientTest = New(*clientURLTest)
	httpClientMock = new(mockHttpClient)
	statusHandlerMock = new(mockStatusHandler)
	clientTest.httpClient = httpClientMock
	clientTest.statusHandler = statusHandlerMock
	ts.IsType(new(Client), clientTest)
}

func (ts *TSClient) TestValidGetWithData() {
	httpClientMock.On("Get", mock.Anything).Return()
	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSClient) TestValidGetWithEmptyData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	httpClientTest.httpClient = mockHTTPClient

	response, err := httpClientTest.Get("")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSClient) TestErrorGetWithData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	httpClientTest.httpClient = mockHTTPClient

	response, err := httpClientTest.Get(idTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestValidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responsePostTest, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Post(reqBodyTest)
	ts.NoError(err)
	ts.NotNil(response)
	ts.Equal(201, response.StatusCode)
}

func (ts *TSClient) TestInvalidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Post(reqBodyTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestValidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSClient) TestValidDeleteEmptyValue() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("", "", "")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSClient) TestValidDeleteNotFound() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
	ts.Equal(404, response.StatusCode)
}

func (ts *TSClient) TestInvalidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.Delete("fakeValue", "version", "1")
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestEmptyResponseAndRequestError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	httpClientTest.httpClient = mockHTTPClient
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
	ts.NoError(err)
	response, err := httpClientTest.doRequest(request)
	ts.ErrorContains(err, "fake error")
	ts.NotNil(response)
}

func (ts *TSClient) TestEmptyResponseAndNilResponseError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	httpClientTest.httpClient = mockHTTPClient
	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
	ts.NoError(err)
	response, err := httpClientTest.doRequest(request)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSClient) TestEmptyResponseAndNilRequest() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	httpClientTest.httpClient = mockHTTPClient
	response, err := httpClientTest.doRequest(nil)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValuesToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
