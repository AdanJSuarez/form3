package httpclient

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
	httpClientTest      *HTTPClient
	mockHTTPClient      *mockHttpClient
	dataTest            = "{data: {moreData: 3}}"
	dataBytesMarshal, _ = json.Marshal(dataTest)
	requestTest         = &http.Request{}
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
	responseDeleteTest = http.Response{
		Status:     "204 Delete",
		StatusCode: 204,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
)

type TSHTTPClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSHTTPClient))
}

func (ts *TSHTTPClient) BeforeTest(_, _ string) {
	clientURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	httpClientTest = New()
	ts.IsType(new(HTTPClient), httpClientTest)
	mockHTTPClient = new(mockHttpClient)
	httpClientTest.httpClient = mockHTTPClient
}

func (ts *TSHTTPClient) TestValidGet() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseGetTest, nil)
	response, err := httpClientTest.Get(requestTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSHTTPClient) TestErrorGet() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fakeError"))
	response, err := httpClientTest.Get(requestTest)
	ts.ErrorContains(err, "fakeError")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestNilRequestGet() {
	response, err := httpClientTest.Get(nil)
	ts.ErrorContains(err, "nil request")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestNotFoundGet() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	response, err := httpClientTest.Get(requestTest)
	ts.NoError(err)
	ts.Equal(404, response.StatusCode)
}

func (ts *TSHTTPClient) TestValidPost() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responsePostTest, nil)
	response, err := httpClientTest.Get(requestTest)
	ts.NoError(err)
	ts.Equal(&responsePostTest, response)
}

func (ts *TSHTTPClient) TestErrorPost() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fakeError2"))
	response, err := httpClientTest.Post(requestTest)
	ts.ErrorContains(err, "fakeError2")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestNilRequestPost() {
	response, err := httpClientTest.Post(nil)
	ts.ErrorContains(err, "nil request")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestValidDelete() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDeleteTest, nil)
	response, err := httpClientTest.Delete(requestTest)
	ts.NoError(err)
	ts.Equal(&responseDeleteTest, response)
}

func (ts *TSHTTPClient) TestErrorDelete() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fakeError3"))
	response, err := httpClientTest.Delete(requestTest)
	ts.ErrorContains(err, "fakeError3")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestNilRequestDelete() {
	response, err := httpClientTest.Delete(nil)
	ts.ErrorContains(err, "nil request")
	ts.Nil(response)
}

func (ts *TSHTTPClient) TestNotFoundDelete() {
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
	response, err := httpClientTest.Delete(requestTest)
	ts.NoError(err)
	ts.Equal(404, response.StatusCode)
}

// func (ts *TSHTTPClient) TestValidGetWithData() {
// 	response, err := httpClientTest.Get(idTest)
// 	ts.NoError(err)
// 	ts.Equal(&responseGetTest, response)
// }

// func (ts *TSHTTPClient) TestValidGetWithEmptyData() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
// 	httpClientTest.httpClient = mockHTTPClient

// 	response, err := httpClientTest.Get("")
// 	ts.NoError(err)
// 	ts.Equal(&responseNotFoundTest, response)
// }

// func (ts *TSHTTPClient) TestErrorGetWithData() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
// 	httpClientTest.httpClient = mockHTTPClient

// 	response, err := httpClientTest.Get(idTest)
// 	ts.Error(err)
// 	ts.Nil(response)
// }

// func (ts *TSHTTPClient) TestValidPost() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responsePostTest, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Post(reqBodyTest)
// 	ts.NoError(err)
// 	ts.NotNil(response)
// 	ts.Equal(201, response.StatusCode)
// }

// func (ts *TSHTTPClient) TestInvalidPost() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake http error"))
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Post(reqBodyTest)
// 	ts.Error(err)
// 	ts.Nil(response)
// }

// func (ts *TSHTTPClient) TestValidDelete() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Delete("fakeValue", "version", "1")
// 	ts.NoError(err)
// 	ts.NotNil(response)
// }

// func (ts *TSHTTPClient) TestValidDeleteEmptyValue() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseDelete, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Delete("", "", "")
// 	ts.NoError(err)
// 	ts.NotNil(response)
// }

// func (ts *TSHTTPClient) TestValidDeleteNotFound() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseNotFoundTest, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Delete("fakeValue", "version", "1")
// 	ts.NoError(err)
// 	ts.Equal(&responseNotFoundTest, response)
// 	ts.Equal(404, response.StatusCode)
// }

// func (ts *TSHTTPClient) TestInvalidDelete() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.Delete("fakeValue", "version", "1")
// 	ts.Error(err)
// 	ts.Nil(response)
// }

// func (ts *TSHTTPClient) TestEmptyResponseAndRequestError() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, fmt.Errorf("fake error"))
// 	httpClientTest.httpClient = mockHTTPClient
// 	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
// 	ts.NoError(err)
// 	response, err := httpClientTest.doRequest(request)
// 	ts.ErrorContains(err, "fake error")
// 	ts.NotNil(response)
// }

// func (ts *TSHTTPClient) TestEmptyResponseAndNilResponseError() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	request, err := httpClientTest.request(POST, httpClientTest.clientURL.String(), reqBodyTest)
// 	ts.NoError(err)
// 	response, err := httpClientTest.doRequest(request)
// 	ts.ErrorContains(err, "nil response")
// 	ts.NotNil(response)
// }

// func (ts *TSHTTPClient) TestEmptyResponseAndNilRequest() {
// 	mockHTTPClient := new(mockHttpClient)
// 	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(nil, nil)
// 	httpClientTest.httpClient = mockHTTPClient
// 	response, err := httpClientTest.doRequest(nil)
// 	ts.ErrorContains(err, "nil response")
// 	ts.NotNil(response)
// }
