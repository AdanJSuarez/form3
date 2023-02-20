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
	errorHandlerMock    *mockErrorHandler
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
	errorHandlerMock = new(mockErrorHandler)
	requestHandlerMock = new(mockRequestHandler)
	clientTest.httpClient = httpClientMock
	clientTest.errorHandler = errorHandlerMock
	clientTest.requestHandler = requestHandlerMock
	ts.IsType(new(Client), clientTest)
}

func (ts *TSClient) TestValidGet() {
	httpClientMock.On("SendRequest", mock.Anything).Return(&responseGetTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)
	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseGetTest, response)
}

func (ts *TSClient) TestValidGetNotFound() {
	httpClientMock.On("SendRequest", mock.Anything).Return(&responseNotFoundTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)

	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSClient) TestErrorGet() {
	httpClientMock.On("SendRequest", mock.Anything).Return(nil, fmt.Errorf("fakeError1"))
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
	httpClientMock.On("SendRequest", mock.Anything).Return(&responsePostTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestPostTest, nil)
	response, err := clientTest.Post(idTest)
	ts.NoError(err)
	ts.Equal(&responsePostTest, response)
}

func (ts *TSClient) TestErrorPost() {
	httpClientMock.On("SendRequest", mock.Anything).Return(nil, fmt.Errorf("fakeError2"))
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
	httpClientMock.On("SendRequest", mock.Anything, mock.Anything, mock.Anything).Return(&responseDeleteTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	requestHandlerMock.On("SetQuery", mock.Anything, mock.Anything, mock.Anything).Return()
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.NoError(err)
	ts.Equal(&responseDeleteTest, response)
}

func (ts *TSClient) TestErrorRequestDelete() {
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf("fakeErrorRequestDelete"))
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.ErrorContains(err, "fakeErrorRequestDelete")
	ts.Nil(response)
}

func (ts *TSClient) TestValidDeleteNotFound() {
	httpClientMock.On("SendRequest", mock.Anything, mock.Anything, mock.Anything).Return(&responseNotFoundTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	requestHandlerMock.On("SetQuery", mock.Anything, mock.Anything, mock.Anything).Return()
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.NoError(err)
	ts.Equal(&responseNotFoundTest, response)
}

func (ts *TSClient) TestInvalidDelete() {
	httpClientMock.On("SendRequest", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("fakeErrorDelete"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	requestHandlerMock.On("SetQuery", mock.Anything, mock.Anything, mock.Anything).Return()
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.ErrorContains(err, "fakeErrorDelete")
	ts.Nil(response)
}

func (ts *TSClient) TestTrueStatusCreated() {
	errorHandlerMock.On("StatusCreated", mock.Anything).Return(true)
	actual := clientTest.statusCreated(&responsePostTest)
	ts.True(actual)
}
func (ts *TSClient) TestFalseStatusCreated() {
	errorHandlerMock.On("StatusCreated", mock.Anything).Return(false)
	actual := clientTest.statusCreated(&responsePostTest)
	ts.False(actual)
}

func (ts *TSClient) TestTrueStatusOK() {
	errorHandlerMock.On("StatusOK", mock.Anything).Return(true)
	actual := clientTest.statusOK(&responseGetTest)
	ts.True(actual)
}

func (ts *TSClient) TestFalseStatusOK() {
	errorHandlerMock.On("StatusOK", mock.Anything).Return(false)
	actual := clientTest.statusOK(&responseGetTest)
	ts.False(actual)
}

func (ts *TSClient) TestTrueStatusNoContent() {
	errorHandlerMock.On("StatusNoContent", mock.Anything).Return(true)
	actual := clientTest.statusNoContent(&responseDeleteTest)
	ts.True(actual)
}

func (ts *TSClient) TestFalseStatusNoContent() {
	errorHandlerMock.On("StatusNoContent", mock.Anything).Return(false)
	actual := clientTest.statusNoContent(&responseDeleteTest)
	ts.False(actual)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValuesToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
