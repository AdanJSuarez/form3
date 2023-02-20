package errorhandler

import (
	"net/http"
	"testing"

	"github.com/AdanJSuarez/form3/internal/client/statuserrorhandler/handler"
	"github.com/stretchr/testify/suite"
)

var (
	statusHandlerTest *StatusErrorHandler
	responseCreated   = &http.Response{
		StatusCode: http.StatusCreated,
	}
	responseOK = &http.Response{
		StatusCode: http.StatusOK,
	}
	responseNoContent = &http.Response{
		StatusCode: http.StatusNoContent,
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

func (ts *TSStatusHandler) TestNextInitialized() {
	ts.NotNil(statusHandlerTest.next)
	ts.Implements(new(handler.StatusErrorHandler), statusHandlerTest.next)
}

// func (ts *TSStatusHandler) TestStatusCreated() {
// 	actual := statusHandlerTest.StatusCreated(responseCreated)
// 	ts.True(actual)
// }

// func (ts *TSStatusHandler) TestNotStatusCreated() {
// 	actual := statusHandlerTest.StatusCreated(responseOK)
// 	ts.False(actual)
// }

// func (ts *TSStatusHandler) TestStatusCreatedNilResponse() {
// 	actual := statusHandlerTest.StatusCreated(nil)
// 	ts.False(actual)
// }

// func (ts *TSStatusHandler) TestStatusOK() {
// 	actual := statusHandlerTest.StatusOK(responseOK)
// 	ts.True(actual)
// }

// func (ts *TSStatusHandler) TestNotStatusOK() {
// 	actual := statusHandlerTest.StatusOK(responseCreated)
// 	ts.False(actual)
// }

// func (ts *TSStatusHandler) TestStatusOKNilResponse() {
// 	actual := statusHandlerTest.StatusOK(nil)
// 	ts.False(actual)
// }

// func (ts *TSStatusHandler) TestStatusNoContent() {
// 	actual := statusHandlerTest.StatusNoContent(responseNoContent)
// 	ts.True(actual)
// }

// func (ts *TSStatusHandler) TestNotStatusNoContent() {
// 	actual := statusHandlerTest.StatusNoContent(responseOK)
// 	ts.False(actual)
// }

// func (ts *TSStatusHandler) TestStatusNoContentNilResponse() {
// 	actual := statusHandlerTest.StatusNoContent(nil)
// 	ts.False(actual)
// }

func (ts *TSStatusHandler) TestHandleErrorUncovered() {
	response, err := statusHandlerTest.StatusError(responseOK)
	ts.ErrorContains(err, "status code 200:")
	ts.ErrorContains(err, "uncovered status code for this request")
	ts.Equal("", response)
}

func (ts *TSStatusHandler) TestHandleError() {
	response, err := statusHandlerTest.StatusError(responseErrorInternalServerError)
	ts.ErrorContains(err, "status code 500:")
	ts.ErrorContains(err, "an internal error occurs or the request times out")
	ts.Nil(response)
}

func (ts *TSStatusHandler) TestHandleErrorNilResponse() {
	response, err := statusHandlerTest.StatusError(nil)
	ts.ErrorContains(err, nilResponseError)
	ts.Nil(response)
}
