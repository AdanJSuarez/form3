package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	gatewayTimeout         StatusErrorHandler
	responseGatewayTimeout = &http.Response{
		StatusCode: http.StatusGatewayTimeout,
	}
	responseFake5 = &http.Response{
		StatusCode: 605,
	}
)

type TSGatewayTimeoutHandler struct{ suite.Suite }

func TestRunGatewayTimeoutSuite(t *testing.T) {
	suite.Run(t, new(TSGatewayTimeoutHandler))
}

func (ts *TSGatewayTimeoutHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	gatewayTimeout = NewGatewayTimeoutHandler()
	gatewayTimeout.SetNext(uncovered)
}

func (ts *TSGatewayTimeoutHandler) TestGatewayTimeoutResponse() {
	err := gatewayTimeout.Execute(responseGatewayTimeout)
	ts.ErrorContains(err, "status code 504")
	ts.ErrorContains(err, gatewayTimeoutMessage)
}

func (ts *TSGatewayTimeoutHandler) TestNotGatewayTimeoutResponse() {
	err := gatewayTimeout.Execute(responseFake5)
	ts.ErrorContains(err, "status code 605:")
	ts.ErrorContains(err, uncoveredMessage)
}
