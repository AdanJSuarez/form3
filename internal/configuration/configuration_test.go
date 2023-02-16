package configuration

import (
	"os"
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
	configurationTest = new(Configuration)
	ts.IsType(new(Configuration), configurationTest)
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

func (ts *TSConfiguration) TestValidInitializeByEnv() {
	_, ok := os.LookupEnv(baseURLEnvKey)
	if !ok {
		os.Setenv(baseURLEnvKey, rawBaseURL)
		defer os.Unsetenv(baseURLEnvKey)
	}
	_, ok = os.LookupEnv(accountPathEnvKey)
	if !ok {
		os.Setenv(accountPathEnvKey, accountPath)
		defer os.Unsetenv(accountPathEnvKey)
	}

	err := configurationTest.InitializeByEnv()
	ts.NoError(err)
	ts.NotEmpty(configurationTest.baseURL)
	ts.NotEmpty(configurationTest.accountPath)
}

func (ts *TSConfiguration) TestInvalidInitializeByEnv() {
	err := configurationTest.InitializeByEnv()
	ts.ErrorContains(err, "failed to get BASE_URL from environment variables")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
}
