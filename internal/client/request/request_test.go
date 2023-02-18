package request

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	rbDataTest     = "{abc: xyzt}"
	desireExpected = "sha-256=9U2Kll78+c7rXln6XrxIu839WxXH7yij2J77+R8d1iM="
)

var (
	requestBodyTest *RequestHandler
	dataByteTest, _ = json.Marshal(rbDataTest)
	bodyTest        = io.NopCloser(bytes.NewBuffer(dataByteTest))
)

type TSRequestBody struct{ suite.Suite }

func TestRunTSRequestBody(t *testing.T) {
	suite.Run(t, new(TSRequestBody))
}

func (ts *TSRequestBody) BeforeTest(_, _ string) {
	requestBodyTest = NewRequestHandler(rbDataTest)
	ts.IsType(&RequestHandler{}, requestBodyTest)
}

func (ts *TSRequestBody) TestBody() {
	body := requestBodyTest.Body()
	ts.Equal(bodyTest, body)
}

func (ts *TSRequestBody) TestSize() {
	size := len(requestBodyTest.data)
	expected := len(dataByteTest)
	ts.Equal(expected, size)
}

func (ts *TSRequestBody) TestDigest() {
	desire := requestBodyTest.digestFormatted()
	ts.Equal(desireExpected, desire)
}
func (ts *TSRequestBody) TestNilBody() {
	requestBodyTest = NewRequestHandler(nil)
	body := requestBodyTest.Body()
	ts.Equal(nil, body)
}
