package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSError struct{ suite.Suite }

func TestRunErrorSuite(t *testing.T) {
	suite.Run(t, new(TSError))
}

func (ts *TSError) BeforeTest(_, _ string) {
	// TODO
}
