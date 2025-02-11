package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	dataTest       = "{abc: xyzt}"
	digestExpected = "sha-256=9U2Kll78+c7rXln6XrxIu839WxXH7yij2J77+R8d1iM="
	requestURLTest = "http://fake.form3.tech/v1/organisation/accounts"
	hostTest       = "fake.form3.tech"
)

var (
	requestTest     *RequestHandler
	dataByteTest, _ = json.Marshal(dataTest)
	bodyTest        = io.NopCloser(bytes.NewBuffer(dataByteTest))
)

type TSRequest struct{ suite.Suite }

func TestRunTSRequest(t *testing.T) {
	suite.Run(t, new(TSRequest))
}

func (ts *TSRequest) BeforeTest(_, _ string) {
	requestTest = NewRequestHandler()
	ts.IsType(&RequestHandler{}, requestTest)
}

func (ts *TSRequest) TestSetCorrectBody() {
	requestTest.Request(dataTest, http.MethodGet, requestURLTest, hostTest)
	body := requestTest.body
	ts.Equal(bodyTest, body)
}

func (ts *TSRequest) TestSetCorrectSize() {
	requestTest.Request(dataTest, http.MethodGet, requestURLTest, hostTest)
	size := len(requestTest.rawData)
	expected := len(dataByteTest)
	ts.Equal(expected, size)
}

func (ts *TSRequest) TestSetCorrectDigest() {
	requestTest.Request(dataTest, http.MethodGet, requestURLTest, hostTest)
	desire := requestTest.digestFormatted()
	ts.Equal(digestExpected, desire)
}
func (ts *TSRequest) TestSetNilBodyWhenNoData() {
	requestTest.Request(nil, http.MethodGet, requestURLTest, hostTest)
	body := requestTest.body
	ts.Nil(body)
}

func (ts *TSRequest) TestSendValidRequestReturnsNoError() {
	request, err := requestTest.Request(dataTest, http.MethodPost, requestURLTest, hostTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal(hostTest, request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Equal(CONTENT_TYPE_VALUE, request.Header.Get(CONTENT_TYPE_KEY))
	ts.Equal(fmt.Sprint(len(dataByteTest)), request.Header.Get(CONTENT_LENGTH_KEY))
	ts.NotEmpty(request.Header.Get(DIGEST_KEY))
}

func (ts *TSRequest) TestSendValidRequestNilDataSetCorrectValues() {
	request, err := requestTest.Request(nil, http.MethodPost, requestURLTest, hostTest)
	ts.NotNil(request)
	ts.NoError(err)
	ts.Equal(hostTest, request.Header.Get(HOST_KEY))
	ts.NotEmpty(request.Header.Get(DATE_KEY))
	ts.Empty(request.Header.Get(CONTENT_TYPE_KEY))
	ts.Empty(request.Header.Get(CONTENT_LENGTH_KEY))
	ts.Empty(request.Header.Get(DIGEST_KEY))
}
func (ts *TSRequest) TestSendValidRequestForDeleteSetCorrectQuery() {
	request, err := requestTest.Request(nil, http.MethodDelete, requestURLTest, hostTest)
	ts.NoError(err)
	requestTest.SetQuery(request, "fakeKey", "fakeValue")
	ts.Equal("fakeKey=fakeValue", request.URL.RawQuery)
}

func (ts *TSRequest) TestDataToBytesReturnsCorrectly() {
	actual := requestTest.dataToBytes(dataTest)
	ts.Equal(dataByteTest, actual)
}

func (ts *TSRequest) TestDataToBodyReturnsCorrectly() {
	requestTest.Request(dataTest, http.MethodGet, requestURLTest, hostTest)
	actual := requestTest.dataToBody()
	ts.Equal(bodyTest, actual)
}

func (ts *TSRequest) TestDataToByteInvalid() {
	actual := requestTest.dataToBytes(nil)
	ts.IsType([]byte{}, actual)
}

func (ts *TSRequest) TestNowUTCFormatted() {
	ts.Contains(requestTest.nowUTCFormatted(), "GMT")

}
