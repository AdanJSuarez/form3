package configuration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURL     = "https://api.fakeaddress/fake"
	invalidURL1    = "https//api.fakeaddress/"
	invalidURL2    = ""
	invalidURL3    = "☺xldkj"
	accountPath    = "/v1/organisation/accounts"
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
)

var configurationTest *Configuration

type TSConfiguration struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSConfiguration))
}

func (ts *TSConfiguration) BeforeTest(_, _ string) {
	configurationTest = New()
	ts.IsType(new(Configuration), configurationTest)
}

func (ts *TSConfiguration) TestValidInitializeByValue() {
	err := configurationTest.InitializeByValue(rawBaseURL, accountPath, organizationID)
	ts.NoError(err)
	ts.Equal(rawBaseURL, configurationTest.BaseURL().String())
	ts.Equal(accountPath, configurationTest.AccountPath())
	ts.Equal(organizationID, configurationTest.OrganizationID())
}

func (ts *TSConfiguration) TestInvalidInitializeByValue() {
	err := configurationTest.InitializeByValue(invalidURL1, accountPath, organizationID)
	ts.ErrorContains(err, "parse")
	ts.Empty(configurationTest.BaseURL())
	ts.Empty(configurationTest.AccountPath())
	ts.Empty(configurationTest.OrganizationID())
}

func (ts *TSConfiguration) TestInvalidBaseURL1() {
	url, err := configurationTest.parseRawBaseURL(invalidURL1)
	ts.Error(err)
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL2() {
	url, err := configurationTest.parseRawBaseURL(invalidURL2)
	ts.Error(err)
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL3() {
	url, err := configurationTest.parseRawBaseURL(invalidURL3)
	ts.Error(err)
	ts.Nil(url)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByYaml() {
	err := configurationTest.InitializeByYaml()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
	ts.Empty(configurationTest.organizationID)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByEnv() {
	err := configurationTest.InitializeByEnv()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
	ts.Empty(configurationTest.organizationID)
}
