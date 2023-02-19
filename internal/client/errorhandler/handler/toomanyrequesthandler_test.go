package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	tooManyRequests         StatusErrorHandler
	responseTooManyRequests = &http.Response{
		StatusCode: http.StatusTooManyRequests,
	}
	responseFake11 = &http.Response{
		StatusCode: 611,
	}
)

type TSTooManyRequestHandler struct{ suite.Suite }

func TestRunTooManyRequestSuite(t *testing.T) {
	suite.Run(t, new(TSTooManyRequestHandler))
}

func (ts *TSTooManyRequestHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	tooManyRequests = NewTooManyRequestsHandler()
	tooManyRequests.SetNext(uncovered)
}

func (ts *TSTooManyRequestHandler) TestTooManyRequestsResponse() {
	err := tooManyRequests.Execute(responseTooManyRequests)
	ts.ErrorContains(err, "status code 429:")
	ts.ErrorContains(err, tooManyRequestsMessage)
}

func (ts *TSTooManyRequestHandler) TestNotTooManyRequestResponse() {
	err := tooManyRequests.Execute(responseFake11)
	ts.ErrorContains(err, "status code 611:")
	ts.ErrorContains(err, uncoveredMessage)
}
