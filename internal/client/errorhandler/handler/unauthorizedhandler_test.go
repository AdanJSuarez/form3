package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	unauthorized         StatusErrorHandler
	responseUnauthorized = &http.Response{
		StatusCode: http.StatusUnauthorized,
	}
	responseFake12 = &http.Response{
		StatusCode: 612,
	}
)

type TSUnauthorizedHandler struct{ suite.Suite }

func TestRunUnauthorizedSuite(t *testing.T) {
	suite.Run(t, new(TSUnauthorizedHandler))
}

func (ts *TSUnauthorizedHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	unauthorized = NewUnauthorizedHandler()
	unauthorized.SetNext(uncovered)
}

func (ts *TSUnauthorizedHandler) TestUnauthorizedResponse() {
	err := unauthorized.Execute(responseUnauthorized)
	ts.ErrorContains(err, "status code 401:")
	ts.ErrorContains(err, unauthorizedMessage)
}

func (ts *TSUnauthorizedHandler) TestNotUnauthorizedResponse() {
	err := unauthorized.Execute(responseFake12)
	ts.ErrorContains(err, "status code 612:")
	ts.ErrorContains(err, uncoveredMessage)
}
