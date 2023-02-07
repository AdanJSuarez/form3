package client

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	requestBodyTest *RequestBody
	bodyTest        = io.NopCloser(bytes.NewBuffer([]byte("abc")))
	sizeTest        = 3
)

type TSRequestBody struct{ suite.Suite }

func TestRunTSRequestBody(t *testing.T) {
	suite.Run(t, new(TSRequestBody))
}

func (ts *TSRequestBody) BeforeTest(_, _ string) {
	requestBodyTest = NewRequestBody(bodyTest, sizeTest)
	ts.IsType(new(RequestBody), requestBodyTest)
}

func (ts *TSRequestBody) TestBody() {
	body := requestBodyTest.Body()
	ts.Equal(bodyTest, body)
}

func (ts *TSRequestBody) TestSize() {
	size := requestBodyTest.Size()
	ts.Equal(sizeTest, size)
}

func (ts *TSRequestBody) TestNilBody() {
	requestBodyTest = NewRequestBody(nil, 0)
	body := requestBodyTest.Body()
	ts.Nil(body)
}
