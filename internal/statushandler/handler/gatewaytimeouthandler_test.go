package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSGatewayTimeoutHandler struct{ suite.Suite }

func TestRunGatewayTimeoutSuite(t *testing.T) {
	suite.Run(t, new(TSGatewayTimeoutHandler))
}

func (ts *TSGatewayTimeoutHandler) BeforeTest(_, _ string) {
	// TODO
}
