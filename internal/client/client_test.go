package client

import (
	"bytes"
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
	baseURLTest     *url.URL
	clientTest      *Client
	dataTest        = []byte("{data: {moreData: 3}}")
	responseGetTest = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataTest)),
	}
	responsePostTest = http.Response{
		Status:     "201 Created",
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer(dataTest)),
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
	reqBodyTest = NewRequestBody(responseNotFoundTest.Body, len(dataTest))
)

type TSClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSClient))
}

func (ts *TSClient) BeforeTest(_, _ string) {
	baseURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	clientTest = New(baseURLTest, valueURL)
	ts.IsType(new(Client), clientTest)
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseGetTest, nil)
	clientTest.client = mockHTTPClient
}

func (ts *TSClient) TestValidGetWithData() {
	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSClient) TestValidGetWithEmptyData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	clientTest.client = mockHTTPClient

	response, err := clientTest.Get("")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSClient) TestErrorGetWithData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	clientTest.client = mockHTTPClient

	response, err := clientTest.Get(idTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestValidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responsePostTest, nil)
	clientTest.client = mockHTTPClient
	response, err := clientTest.Post(reqBodyTest)
	ts.NoError(err)
	ts.NotNil(response)
	ts.Equal(201, response.StatusCode)
}

func (ts *TSClient) TestInvalidPost() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
	clientTest.client = mockHTTPClient
	response, err := clientTest.Post(reqBodyTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestValidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	clientTest.client = mockHTTPClient
	response, err := clientTest.Delete("fakeValue")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSClient) TestValidDeleteEmptyValue() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
	clientTest.client = mockHTTPClient
	response, err := clientTest.Delete("")
	ts.NoError(err)
	ts.NotNil(response)
}

func (ts *TSClient) TestValidDeleteNotFound() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	clientTest.client = mockHTTPClient
	response, err := clientTest.Delete("fakeValue")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
	ts.Equal(404, response.StatusCode)
}

func (ts *TSClient) TestInvalidDelete() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	clientTest.client = mockHTTPClient
	response, err := clientTest.Delete("fakeValue")
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestValidRequest() {
	request, err := clientTest.request(POST, clientTest.stringURL, reqBodyTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal("api.fakeaddress.tech", request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Equal(CONTENT_TYPE_VALUE, request.Header.Get(CONTENT_TYPE_KEY))
	ts.Equal(fmt.Sprint(len(dataTest)), request.Header.Get(CONTENT_LENGTH_KEY))
}

func (ts *TSClient) TestValidRequestNotBody() {
	rbTest := NewRequestBody(nil, 0)
	request, err := clientTest.request(POST, clientTest.stringURL, rbTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal("api.fakeaddress.tech", request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Equal("", request.Header.Get(CONTENT_TYPE_KEY))
	ts.Equal("", request.Header.Get(CONTENT_LENGTH_KEY))
}

func (ts *TSClient) TestEmptyResponseAndRequestError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
	clientTest.client = mockHTTPClient
	request, err := clientTest.request(POST, clientTest.stringURL, reqBodyTest)
	ts.NoError(err)
	response, err := clientTest.doRequest(request)
	ts.ErrorContains(err, "fake error")
	ts.NotNil(response)
}

func (ts *TSClient) TestEmptyResponseAndNilResponseError() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	clientTest.client = mockHTTPClient
	request, err := clientTest.request(POST, clientTest.stringURL, reqBodyTest)
	ts.NoError(err)
	response, err := clientTest.doRequest(request)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSClient) TestEmptyResponseAndNilRequest() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
	clientTest.client = mockHTTPClient
	response, err := clientTest.doRequest(nil)
	ts.ErrorContains(err, "nil response")
	ts.NotNil(response)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValueToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
