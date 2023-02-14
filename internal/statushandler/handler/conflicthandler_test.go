package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSConflictHandler struct{ suite.Suite }

func TestRunConflictSuite(t *testing.T) {
	suite.Run(t, new(TSConflictHandler))
}

func (ts *TSConflictHandler) BeforeTest(_, _ string) {
	// TODO
}
