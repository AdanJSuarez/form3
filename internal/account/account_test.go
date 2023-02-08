package account

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	baseURL     = "http://fakeurl.com"
	accountPath = "/v1/origanisation/"
)

var (
	accountTest    *Account
	baseURLTest, _ = url.Parse(baseURL)
)

type TSAccount struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSAccount))
}

func (ts *TSAccount) BeforeTest(_, _ string) {
	accountTest = New(*baseURLTest, accountPath)
	ts.IsType(new(Account), accountTest)
}
