package handler

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	conflict         StatusErrorHandler
	responseConflict = &http.Response{
		StatusCode: http.StatusConflict,
		Body: io.NopCloser(bytes.NewBuffer([]byte(`
		{
			"error_message": "Duplicate id f72c5098-bf0f-4526-a215-54e5c1e2e687",
			"error_code": "4bc0fa5d-231e-43f3-af79-8fc371d95a31"
		}
		`))),
	}
	responseFake3 = &http.Response{
		StatusCode: 603,
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

func (ts *TSConflictHandler) TestConflictResponseReturnsCorrectError() {
	err := conflict.Execute(responseConflict)
	ts.ErrorContains(err, "status code 409")
	ts.ErrorContains(err, "errorCode: 4bc0fa5d-231e-43f3-af79-8fc371d95a31")
}

func (ts *TSConflictHandler) TestNotConflictResponseReturnsUncoveredError() {
	err := conflict.Execute(responseFake3)
	ts.ErrorContains(err, "status code 603:")
	ts.ErrorContains(err, uncoveredMessage)
}
