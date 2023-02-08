package client

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
	requestBodyTest RequestBody
	dataByteTest, _ = json.Marshal(rbDataTest)
	bodyTest        = io.NopCloser(bytes.NewBuffer(dataByteTest))
)

type TSRequestBody struct{ suite.Suite }

func TestRunTSRequestBody(t *testing.T) {
	suite.Run(t, new(TSRequestBody))
}

func (ts *TSRequestBody) BeforeTest(_, _ string) {
	requestBodyTest = NewRequestBody(rbDataTest)
	ts.IsType(RequestBody{}, requestBodyTest)
}

func (ts *TSRequestBody) TestBody() {
	body := requestBodyTest.Body()
	ts.Equal(bodyTest, body)
}

func (ts *TSRequestBody) TestSize() {
	size := requestBodyTest.Size()
	ts.Equal(len(dataByteTest), size)
}

func (ts *TSRequestBody) TestDesire() {
	desire := requestBodyTest.Digest()
	ts.Equal(desireExpected, desire)
}
func (ts *TSRequestBody) TestNilBody() {
	requestBodyTest = NewRequestBody(nil)
	body := requestBodyTest.Body()
	ts.Equal(nil, body)
}
