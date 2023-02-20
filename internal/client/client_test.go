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
	clientURLTest          *url.URL
	clientTest             *Client
	httpClientMock         *mockHttpClient
	statusErrorHandlerMock *mockStatusErrorHandler
	requestHandlerMock     *mockRequestHandler
	dataTest               = "{data: {moreData: 55}}"
	dataBytesMarshal, _    = json.Marshal(dataTest)
	requestGetTest         = http.Request{
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
	httpClientMock = newMockHttpClient(ts.T())
	statusErrorHandlerMock = newMockStatusErrorHandler(ts.T())
	requestHandlerMock = newMockRequestHandler(ts.T())
	clientTest.httpClient = httpClientMock
	clientTest.statusErrorHandler = statusErrorHandlerMock
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

func (ts *TSClient) TestGetReturnNotFound() {
	httpClientMock.On("SendRequest", mock.Anything).Return(&responseNotFoundTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)
	statusErrorHandlerMock.On("StatusError", mock.Anything).Return(nil, fmt.Errorf("not found"))
	response, err := clientTest.Get(idTest)
	ts.ErrorContains(err, "not found")
	ts.Nil(response)
}

func (ts *TSClient) TestGetWithErrorOnSendRequest() {
	httpClientMock.On("SendRequest", mock.Anything).Return(nil, fmt.Errorf("fakeError1"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestGetTest, nil)

	response, err := clientTest.Get(idTest)
	ts.ErrorContains(err, "fakeError1")
	ts.Nil(response)
}

func (ts *TSClient) TestGetWithErrorOnRequest() {
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

func (ts *TSClient) TestPostWithErrorOnSendRequest() {
	httpClientMock.On("SendRequest", mock.Anything).Return(nil, fmt.Errorf("fakeError2"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestPostTest, nil)
	response, err := clientTest.Post(idTest)
	ts.ErrorContains(err, "fakeError2")
	ts.Nil(response)
}

func (ts *TSClient) TestPostWithErrorOnRequest() {
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

func (ts *TSClient) TestDeleteWithErrorOnRequest() {
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf("fakeErrorRequestDelete"))
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.ErrorContains(err, "fakeErrorRequestDelete")
	ts.Nil(response)
}

func (ts *TSClient) TestDeleteReturnNotFound() {
	httpClientMock.On("SendRequest", mock.Anything, mock.Anything, mock.Anything).Return(&responseNotFoundTest, nil)
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	requestHandlerMock.On("SetQuery", mock.Anything, mock.Anything, mock.Anything).Return()
	statusErrorHandlerMock.On("StatusError", mock.Anything).Return(nil, fmt.Errorf("not found"))
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.ErrorContains(err, "not found")
	ts.Nil(response)
}

func (ts *TSClient) TestDeleteWithErrorOnSendRequest() {
	httpClientMock.On("SendRequest", mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("fakeErrorDelete"))
	requestHandlerMock.On("Request", mock.Anything, mock.Anything, mock.Anything,
		mock.Anything).Return(&requestDeleteTest, nil)
	requestHandlerMock.On("SetQuery", mock.Anything, mock.Anything, mock.Anything).Return()
	response, err := clientTest.Delete(idTest, "version", "0")
	ts.ErrorContains(err, "fakeErrorDelete")
	ts.Nil(response)
}

func (ts *TSClient) TestStatusCreatedTrue() {
	actual := clientTest.statusCreated(&responsePostTest)
	ts.True(actual)
}
func (ts *TSClient) TestStatusCreatedFalse() {
	actual := clientTest.statusCreated(&responseNotFoundTest)
	ts.False(actual)
}
func (ts *TSClient) TestStatusCreatedNilResponse() {
	actual := clientTest.statusCreated(nil)
	ts.False(actual)
}

func (ts *TSClient) TestStatusOKTrue() {
	actual := clientTest.statusOK(&responseGetTest)
	ts.True(actual)
}

func (ts *TSClient) TestStatusOKFalse() {
	actual := clientTest.statusOK(&responseNotFoundTest)
	ts.False(actual)
}
func (ts *TSClient) TestStatusOKNilResponse() {
	actual := clientTest.statusOK(nil)
	ts.False(actual)
}

func (ts *TSClient) TestStatusNoContentTrue() {
	actual := clientTest.statusNoContent(&responseDeleteTest)
	ts.True(actual)
}

func (ts *TSClient) TestStatusNoContentFalse() {
	actual := clientTest.statusNoContent(&responseNotFoundTest)
	ts.False(actual)
}

func (ts *TSClient) TestStatusNoContentNilResponse() {
	actual := clientTest.statusNoContent(nil)
	ts.False(actual)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValuesToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
