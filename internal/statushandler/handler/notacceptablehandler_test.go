package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSNotAcceptableHandler struct{ suite.Suite }

func TestRunNotAcceptableSuite(t *testing.T) {
	suite.Run(t, new(TSGatewayTimeoutHandler))
}

func (ts *TSNotAcceptableHandler) BeforeTest(_, _ string) {
	// TODO
}
