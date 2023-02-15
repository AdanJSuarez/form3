package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	serviceUnavailable         StatusErrorHandler
	responseServiceUnavailable = &http.Response{
		StatusCode: http.StatusServiceUnavailable,
	}
	responseFake10 = &http.Response{
		StatusCode: 610,
	}
)

type TSServiceUnavailableHandler struct{ suite.Suite }

func TestRunServiceUnavailableSuite(t *testing.T) {
	suite.Run(t, new(TSServiceUnavailableHandler))
}

func (ts *TSServiceUnavailableHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	serviceUnavailable = NewServiceUnavailableHandler()
	serviceUnavailable.SetNext(uncovered)
}

func (ts *TSServiceUnavailableHandler) TestServiceUnavailableResponse() {
	err := serviceUnavailable.Execute(responseServiceUnavailable)
	ts.ErrorContains(err, "status code 503")
	ts.ErrorContains(err, serviceUnavailableMessage)
}

func (ts *TSServiceUnavailableHandler) TestNotServiceUnavailableResponse() {
	err := serviceUnavailable.Execute(responseFake10)
	ts.ErrorContains(err, "status code 610:")
	ts.ErrorContains(err, uncoveredMessage)
}
