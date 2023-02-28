package statuserrorhandler

import (
	"net/http"
	"testing"

	"github.com/AdanJSuarez/form3/internal/client/statuserrorhandler/handler"
	"github.com/stretchr/testify/suite"
)

var (
	statusHandlerTest *StatusErrorHandler
	responseOK        = &http.Response{
		StatusCode: http.StatusOK,
	}
	responseErrorInternalServerError = &http.Response{
		StatusCode: http.StatusInternalServerError,
	}
)

type TSStatusHandler struct{ suite.Suite }

func TestRunStatusHandlerSuite(t *testing.T) {
	suite.Run(t, new(TSStatusHandler))
}

func (ts *TSStatusHandler) BeforeTest(_, _ string) {
	statusHandlerTest = NewStatusErrorHandler()
	ts.IsType(new(StatusErrorHandler), statusHandlerTest)
}

func (ts *TSStatusHandler) TestNextInitializedAndCorrectType() {
	ts.NotNil(statusHandlerTest.next)
	ts.Implements(new(handler.StatusErrorHandler), statusHandlerTest.next)
}

func (ts *TSStatusHandler) TestUncoverResponseReturnsUncoveredError() {
	response, err := statusHandlerTest.StatusError(responseOK)
	ts.ErrorContains(err, "status code 200:")
	ts.ErrorContains(err, "uncovered status code for this request")
	ts.Nil(response)
}

func (ts *TSStatusHandler) TestStatusErrorReturnsErrorCorrectly() {
	response, err := statusHandlerTest.StatusError(responseErrorInternalServerError)
	ts.ErrorContains(err, "status code 500:")
	ts.ErrorContains(err, "an internal error occurs or the request\n\t\ttimes out")
	ts.Nil(response)
}

func (ts *TSStatusHandler) TestStatusErrorNilResponseReturnErrorCorrectly() {
	response, err := statusHandlerTest.StatusError(nil)
	ts.ErrorContains(err, nilResponseError)
	ts.Nil(response)
}
