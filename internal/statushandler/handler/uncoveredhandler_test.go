package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSUncoveredHandler struct{ suite.Suite }

func TestRunUncoveredSuite(t *testing.T) {
	suite.Run(t, new(TSUncoveredHandler))
}

func (ts *TSUncoveredHandler) BeforeTest(_, _ string) {
	// TODO
}
