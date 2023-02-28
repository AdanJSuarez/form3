package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	uncovered  StatusErrorHandler
	responseOK = &http.Response{
		StatusCode: http.StatusOK,
	}
	responseFake13 = &http.Response{
		StatusCode: 613,
	}
)

type TSUncoveredHandler struct{ suite.Suite }

func TestRunUncoveredSuite(t *testing.T) {
	suite.Run(t, new(TSUncoveredHandler))
}

func (ts *TSUncoveredHandler) BeforeTest(_, _ string) {
	uncovered = NewUncoveredHandler()
}

func (ts *TSUncoveredHandler) TestUncoveredResponseReturnsUncoveredError() {
	err := uncovered.Execute(responseOK)
	ts.ErrorContains(err, "status code 200:")
	ts.ErrorContains(err, uncoveredMessage)
}
