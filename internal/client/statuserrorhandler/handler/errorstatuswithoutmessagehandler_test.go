package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	errorStatusHandlerTest StatusErrorHandler
	responseUnauthorized   = &http.Response{
		StatusCode: http.StatusUnauthorized,
	}
	responseNotAcceptable = &http.Response{
		StatusCode: http.StatusNotAcceptable,
	}
	responseServerError = &http.Response{
		StatusCode: http.StatusInternalServerError,
	}
	responseBadGateway = &http.Response{
		StatusCode: http.StatusBadGateway,
	}
	responseServiceUnavailable = &http.Response{
		StatusCode: http.StatusServiceUnavailable,
	}
	responseGatewayTimeout = &http.Response{
		StatusCode: http.StatusGatewayTimeout,
	}
	responseFake = &http.Response{
		StatusCode: 605,
	}
)

type TSErrorStatusWithoutMessageHandler struct{ suite.Suite }

func TestRunErrorStatusWithoutMessageHandlerSuite(t *testing.T) {
	suite.Run(t, new(TSErrorStatusWithoutMessageHandler))
}

func (ts *TSErrorStatusWithoutMessageHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	errorStatusHandlerTest = NewErrorStatusWithoutMessageHandler()
	errorStatusHandlerTest.SetNext(uncovered)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestUnauthorizedResponse() {
	err := errorStatusHandlerTest.Execute(responseUnauthorized)
	ts.ErrorContains(err, "status code 401:")
	ts.ErrorContains(err, unauthorizedMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestNotAcceptableResponse() {
	err := errorStatusHandlerTest.Execute(responseNotAcceptable)
	ts.ErrorContains(err, "status code 406")
	ts.ErrorContains(err, notAcceptableMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestServerErrorResponse() {
	err := errorStatusHandlerTest.Execute(responseServerError)
	ts.ErrorContains(err, "status code 500")
	ts.ErrorContains(err, serverErrorMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestBadGatewayResponse() {
	err := errorStatusHandlerTest.Execute(responseBadGateway)
	ts.ErrorContains(err, "status code 502")
	ts.ErrorContains(err, badGatewayMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestServiceUnavailableResponse() {
	err := errorStatusHandlerTest.Execute(responseServiceUnavailable)
	ts.ErrorContains(err, "status code 503")
	ts.ErrorContains(err, serviceUnavailableMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestGatewayTimeoutResponse() {
	err := errorStatusHandlerTest.Execute(responseGatewayTimeout)
	ts.ErrorContains(err, "status code 504")
	ts.ErrorContains(err, gatewayTimeoutMessage)
}

func (ts *TSErrorStatusWithoutMessageHandler) TestNotGatewayTimeoutResponse() {
	err := errorStatusHandlerTest.Execute(responseFake)
	ts.ErrorContains(err, "status code 605:")
	ts.ErrorContains(err, uncoveredMessage)
}
