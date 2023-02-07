package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	urlTest = "https://api.fakeaddress/fake/v1/organisation/accounts"
	idTest  = "020cf7d8-01b9-461d-89d4-89d57fd0d998"
)

var (
	clientTest   *Client
	dataTest     = []byte("{data: {moreData: 3}}")
	responseTest = http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataTest)),
	}
	responseNotFoundTest = http.Response{
		Status:     "404 Not Found",
		StatusCode: 404,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(""))),
	}
)

type TSClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSClient))
}

func (ts *TSClient) BeforeTest(_, _ string) {
	clientTest = New(urlTest)
	ts.IsType(new(Client), clientTest)
}

func (ts *TSClient) TestValidGetWithData() {
	mockHTTPClient := new(mockHttpClient)
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseTest, nil)
	clientTest.client = mockHTTPClient

	response, err := clientTest.Get(idTest)
	ts.NoError(err)
	ts.Equal(&responseTest, response)
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
	mockHTTPClient.On("Do", mock.AnythingOfType("*http.Request")).Return(&responseTest, fmt.Errorf("fake http error"))
	clientTest.client = mockHTTPClient

	response, err := clientTest.Get(idTest)
	ts.Error(err)
	ts.Nil(response)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValueToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
