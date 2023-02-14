package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSBadRequestHandler struct{ suite.Suite }

func TestRunBadRequestSuite(t *testing.T) {
	suite.Run(t, new(TSBadRequestHandler))
}

func (ts *TSBadRequestHandler) BeforeTest(_, _ string) {
	// TODO
}
