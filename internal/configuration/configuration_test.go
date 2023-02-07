package configuration

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	baseURL        = "https://api.fakeaddress/fake"
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
	err := configurationTest.InitializeByValue(baseURL, accountPath, organizationID)
	ts.NoError(err)
	accountURL := baseURL + accountPath
	ts.Equal(baseURL, configurationTest.baseURL)
	ts.Equal(accountURL, configurationTest.accountURL)
	ts.Equal(organizationID, configurationTest.organizationID)
}

func (ts *TSConfiguration) TestInvalidInitializeByValue() {
	err := configurationTest.InitializeByValue(invalidURL1, accountPath, organizationID)
	ts.ErrorContains(err, "parse")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountURL)
	ts.Empty(configurationTest.organizationID)
}

func (ts *TSConfiguration) TestValidAccountURL() {
	err := configurationTest.InitializeByValue(baseURL, accountPath, organizationID)
	ts.NoError(err)
	accountURL := baseURL + accountPath
	ts.Equal(accountURL, configurationTest.AccountURL())
}

func (ts *TSConfiguration) TestValidOrganisationID() {
	err := configurationTest.InitializeByValue(baseURL, accountPath, organizationID)
	ts.NoError(err)
	ts.Equal(organizationID, configurationTest.OrganizationID())
}

func (ts *TSConfiguration) TestValidBaseURL() {
	err := configurationTest.validateBaseURL(baseURL)
	ts.NoError(err)
}

func (ts *TSConfiguration) TestInvalidBaseURL1() {
	err := configurationTest.validateBaseURL(invalidURL1)
	ts.Error(err)
}

func (ts *TSConfiguration) TestInvalidBaseURL2() {
	err := configurationTest.validateBaseURL(invalidURL2)
	ts.Error(err)
}

func (ts *TSConfiguration) TestInvalidBaseURL3() {
	err := configurationTest.validateBaseURL(invalidURL3)
	ts.Error(err)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByYaml() {
	err := configurationTest.InitializeByYaml()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountURL)
	ts.Empty(configurationTest.organizationID)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByEnv() {
	err := configurationTest.InitializeByEnv()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountURL)
	ts.Empty(configurationTest.organizationID)
}

func (ts *TSConfiguration) TestJoinPathToURL() {
	url, err := url.JoinPath(baseURL, accountPath)
	ts.NoError(err)
	ts.NotEmpty(url)
}
