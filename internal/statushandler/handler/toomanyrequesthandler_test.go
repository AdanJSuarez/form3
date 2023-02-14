package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSTooManyRequestHandler struct{ suite.Suite }

func TestRunTooManyRequestSuite(t *testing.T) {
	suite.Run(t, new(TSTooManyRequestHandler))
}

func (ts *TSTooManyRequestHandler) BeforeTest(_, _ string) {
	// TODO
}
