package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSServerErrorHandler struct{ suite.Suite }

func TestRunServerErrorSuite(t *testing.T) {
	suite.Run(t, new(TSServerErrorHandler))
}

func (ts *TSServerErrorHandler) BeforeTest(_, _ string) {
	// TODO
}
