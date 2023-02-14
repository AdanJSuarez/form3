package handler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSBadGatewayHandler struct{ suite.Suite }

func TestRunBadGatewaySuite(t *testing.T) {
	suite.Run(t, new(TSBadGatewayHandler))
}

func (ts *TSBadGatewayHandler) BeforeTest(_, _ string) {
	// TODO
}
