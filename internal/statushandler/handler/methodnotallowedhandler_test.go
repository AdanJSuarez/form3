package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSMethodNotAllowedHandler struct{ suite.Suite }

func TestRunMethodNotAllowedSuite(t *testing.T) {
	suite.Run(t, new(TSMethodNotAllowedHandler))
}

func (ts *TSMethodNotAllowedHandler) BeforeTest(_, _ string) {
	// TODO
}
