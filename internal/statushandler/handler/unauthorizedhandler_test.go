package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSUnauthorizedHandler struct{ suite.Suite }

func TestRunUnauthorizedSuite(t *testing.T) {
	suite.Run(t, new(TSUnauthorizedHandler))
}

func (ts *TSUnauthorizedHandler) BeforeTest(_, _ string) {
	// TODO
}
