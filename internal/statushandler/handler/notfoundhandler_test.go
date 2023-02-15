package handler

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	notFound         StatusErrorHandler
	responseNotFound = &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBuffer([]byte(dataCodeMessage))),
	}
	responseFake8 = &http.Response{
		StatusCode: 608,
	}
)

type TSNotFoundHandler struct{ suite.Suite }

func TestRunNotFoundSuite(t *testing.T) {
	suite.Run(t, new(TSNotFoundHandler))
}

func (ts *TSNotFoundHandler) BeforeTest(_, _ string) {
	uncovered := NewUncoveredHandler()
	notFound = NewNotFoundHandler()
	notFound.SetNext(uncovered)
}

func (ts *TSNotFoundHandler) TestNotFoundResponse() {
	err := notFound.Execute(responseNotFound)
	ts.ErrorContains(err, "status code 404")
	ts.ErrorContains(err, notFoundMessage)
}

func (ts *TSNotFoundHandler) TestNotANotFoundResponse() {
	err := notFound.Execute(responseFake8)
	ts.ErrorContains(err, "status code 608:")
	ts.ErrorContains(err, uncoveredMessage)
}
