package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	conflict         StatusErrorHandler
	responseConflict = &http.Response{
		StatusCode: http.StatusConflict,
	}
	responseFake3 = &http.Response{
		StatusCode: 600,
	}
)

type TSConflictHandler struct{ suite.Suite }

func TestRunConflictSuite(t *testing.T) {
	suite.Run(t, new(TSConflictHandler))
}

func (ts *TSConflictHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	conflict = NewConflictHandler()
	conflict.SetNext(uncovered)
}

func (ts *TSConflictHandler) TestConflictResponse() {
	err := conflict.Execute(responseConflict)
	ts.ErrorContains(err, "status code 409")
	ts.ErrorContains(err, conflictHandlerMessage)
}

func (ts *TSConflictHandler) TestNotConflictResponse() {
	err := conflict.Execute(responseFake3)
	ts.ErrorContains(err, "status code 600:")
	ts.ErrorContains(err, uncoveredMessage)
}
