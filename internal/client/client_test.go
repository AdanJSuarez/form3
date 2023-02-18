package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	clientURLTest       *url.URL
	clientTest          *Client
	httpClientMock      *mockHttpClient
	statusHandlerMock   *mockStatusHandler
	requestHandlerMock  *mockRequestHandler
	dataTest            = "{data: {moreData: 55}}"
	dataBytesMarshal, _ = json.Marshal(dataTest)
	requestGetTest      = http.Request{
		Method: http.MethodGet,
	}
	requestPostTest = http.Request{
		Method: http.MethodPost,
	}
	requestDeleteTest = http.Request{
		Method: http.MethodDelete,
	}
	responseGetTest = http.Response{
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
	responseDeleteTest = http.Response{
		Status:     "204 Delete",
		StatusCode: 204,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
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

func (ts *TSClient) TestValidGet() {
	httpClientMock.On("Get", mock.Anything).Return(&responseGetTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)
	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSClient) TestValidGetNotFound() {
	httpClientMock.On("Get", mock.Anything).Return(&responseNotFoundTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)

	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSClient) TestErrorGet() {
	httpClientMock.On("Get", mock.Anything).Return(nil, fmt.Errorf("fakeError1"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)

	response, err := clientTest.Get(idTest)
	ts.ErrorContains(err, "fakeError1")
	ts.Nil(response)
}

func (ts *TSClient) TestRequestErrorGet() {
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf("fakeErrorRequest"))

	response, err := clientTest.Get(idTest)
	ts.ErrorContains(err, "fakeErrorRequest")
	ts.Nil(response)
}

func (ts *TSClient) TestValidPost() {
	httpClientMock.On("Post", mock.Anything).Return(&responsePostTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestPostTest, nil)
	response, err := clientTest.Post(idTest)
	ts.NoError(err)
	ts.Equal(&responsePostTest, response)
}

func (ts *TSClient) TestErrorPost() {
	httpClientMock.On("Post", mock.Anything).Return(nil, fmt.Errorf("fakeError2"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestPostTest, nil)
	response, err := clientTest.Post(idTest)
	ts.ErrorContains(err, "fakeError2")
	ts.Nil(response)
}

func (ts *TSClient) TestRequestErrorPost() {
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf("fakeErrorRequest2"))
	response, err := clientTest.Post(idTest)
	ts.ErrorContains(err, "fakeErrorRequest2")
	ts.Nil(response)
}

func (ts *TSClient) TestValidDelete() {
	httpClientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(&responseDeleteTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.NoError(err)
	ts.Equal(&responsePostTest, response)
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
