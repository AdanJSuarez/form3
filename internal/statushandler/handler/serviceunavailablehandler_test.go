package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSServiceUnavailableHandler struct{ suite.Suite }

func TestRunServiceUnavailableSuite(t *testing.T) {
	suite.Run(t, new(TSGatewayTimeoutHandler))
}

func (ts *TSServiceUnavailableHandler) BeforeTest(_, _ string) {
	// TODO
}
