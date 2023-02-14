package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSForbiddenHandler struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSForbiddenHandler))
}

func (ts *TSForbiddenHandler) BeforeTest(_, _ string) {
	// TODO
}
