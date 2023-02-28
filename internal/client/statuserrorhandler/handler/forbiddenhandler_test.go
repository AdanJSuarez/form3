package handler

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	dataTypeDescription2 = `{
		"error": "invalid_grant",
		"error_description": "Wrong email or password."
	}`

	forbidden         StatusErrorHandler
	responseForbidden = &http.Response{
		StatusCode: http.StatusForbidden,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(dataTypeDescription2))),
	}
	responseFake4 = &http.Response{
		StatusCode: 604,
	}
)

type TSForbiddenHandler struct{ suite.Suite }

func TestRunForbiddenSuite(t *testing.T) {
	suite.Run(t, new(TSForbiddenHandler))
}

func (ts *TSForbiddenHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	forbidden = NewForbiddenHandler()
	forbidden.SetNext(uncovered)
}

func (ts *TSForbiddenHandler) TestForbiddenResponseReturnsErrorCorrectly() {
	err := forbidden.Execute(responseForbidden)
	ts.ErrorContains(err, "status code 403")
	ts.ErrorContains(err, "error: invalid_grant")
	ts.ErrorContains(err, "errorDescription: Wrong email or password.")
}

func (ts *TSForbiddenHandler) TestNotForbiddenResponseReturnsUncoveredError() {
	err := forbidden.Execute(responseFake4)
	ts.ErrorContains(err, "status code 604:")
	ts.ErrorContains(err, uncoveredMessage)
}
