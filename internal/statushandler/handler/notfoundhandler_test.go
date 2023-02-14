package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSNotFoundHandler struct{ suite.Suite }

func TestRunNotFoundSuite(t *testing.T) {
	suite.Run(t, new(TSNotFoundHandler))
}

func (ts *TSNotFoundHandler) BeforeTest(_, _ string) {
	// TODO
}
