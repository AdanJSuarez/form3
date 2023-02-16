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

var configurationTest *configuration

type TSConfiguration struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSConfiguration))
}

func (ts *TSConfiguration) BeforeTest(_, _ string) {
	configurationTest = &configuration{}
}

func (ts *TSConfiguration) TestNewType() {
	config := New()
	ts.IsType(new(configuration), config)
}

func (ts *TSConfiguration) TestValidInitializeByValue() {
	err := configurationTest.InitializeByValue(rawBaseURL, accountPath)
	ts.NoError(err)
	ts.Equal(rawBaseURL, configurationTest.BaseURL().String())
	ts.Equal(accountPath, configurationTest.AccountPath())
}

func (ts *TSConfiguration) TestInvalidInitializeByValue() {
	err := configurationTest.InitializeByValue(invalidURL1, accountPath)
	ts.ErrorContains(err, "parse")
	ts.Empty(configurationTest.BaseURL())
	ts.Empty(configurationTest.AccountPath())
}

func (ts *TSConfiguration) TestInvalidBaseURL1() {
	url, err := configurationTest.parseRawBaseURL(invalidURL1)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL2() {
	url, err := configurationTest.parseRawBaseURL(invalidURL2)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL3() {
	url, err := configurationTest.parseRawBaseURL(invalidURL3)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByYaml() {
	err := configurationTest.InitializeByYaml()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
}

func (ts *TSConfiguration) TestNotImplementedInitializeByEnv() {
	err := configurationTest.InitializeByEnv()
	ts.ErrorContains(err, "not implemented")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
}
