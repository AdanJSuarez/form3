package handler

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	badRequest StatusErrorHandler
	data1      = `{"error_message": "Message parsing failed: Unexpected character (';' (code 34)): was expecting comma to separate Object entries ",
		"error_code": "d0a17902-63ed-4cb6-a8e8-fac5ca31b0b7"}`
	badRequestResponse = &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(data1))),
	}
	responseFake2 = &http.Response{
		StatusCode: 602,
	}
)

type TSBadRequestHandler struct{ suite.Suite }

func TestRunBadRequestSuite(t *testing.T) {
	suite.Run(t, new(TSBadRequestHandler))
}

func (ts *TSBadRequestHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	badRequest = NewBadRequestHandler()
	badRequest.SetNext(uncovered)
}

func (ts *TSBadRequestHandler) TestBadRequest() {
	err := badRequest.Execute(badRequestResponse)
	ts.ErrorContains(err, "status code 400")
	ts.ErrorContains(err, "errorCode: d0a17902-63ed-4cb6-a8e8-fac5ca31b0b7")
	ts.ErrorContains(err, "errorMessage: Message parsing failed: Unexpected character (';' (code 34)): was expecting comma to separate Object entries")
}

func (ts *TSBadRequestHandler) TestNotBadRequest() {
	err := badRequest.Execute(responseFake2)
	ts.ErrorContains(err, "status code 602:")
	ts.ErrorContains(err, uncoveredMessage)
}
