package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	methodNotAllowed         StatusErrorHandler
	responseMethodNotAllowed = &http.Response{
		StatusCode: http.StatusMethodNotAllowed,
	}
	responseFake6 = &http.Response{
		StatusCode: 606,
	}
)

type TSMethodNotAllowedHandler struct{ suite.Suite }

func TestRunMethodNotAllowedSuite(t *testing.T) {
	suite.Run(t, new(TSMethodNotAllowedHandler))
}

func (ts *TSMethodNotAllowedHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	methodNotAllowed = NewMethodNotAllowedHandler()
	methodNotAllowed.SetNext(uncovered)
}

func (ts *TSMethodNotAllowedHandler) TestMethodNotAllowedResponseReturnsErrorCorrectly() {
	err := methodNotAllowed.Execute(responseMethodNotAllowed)
	ts.ErrorContains(err, "status code 405")
	ts.ErrorContains(err, methodNotAllowedMessage)
}

func (ts *TSMethodNotAllowedHandler) TestNotMethodNotAllowedResponseReturnsUncoveredError() {
	err := methodNotAllowed.Execute(responseFake6)
	ts.ErrorContains(err, "status code 606:")
	ts.ErrorContains(err, uncoveredMessage)
}
