package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURL  = "https://api.fakeaddress/fake"
	invalidURL1 = "https//api.fakeaddress/"
	invalidURL2 = ""
	invalidURL3 = "☺xldkj"
	accountPath = "/v1/organisation/accounts"
)

var configurationTest *Configuration

type TSConfiguration struct{ suite.Suite }

func TestRunConfigurationSuite(t *testing.T) {
	suite.Run(t, new(TSConfiguration))
}

func (ts *TSConfiguration) BeforeTest(_, _ string) {
	configurationTest = New()
	ts.IsType(new(Configuration), configurationTest)
}

func (ts *TSConfiguration) TestValidValuesInitializeByValueReturnsNoError() {
	err := configurationTest.InitializeByValue(rawBaseURL, accountPath)
	ts.NoError(err)
	ts.Equal(rawBaseURL, configurationTest.BaseURL().String())
	ts.Equal(accountPath, configurationTest.AccountPath())
}

func (ts *TSConfiguration) TestInvalidValuesInitializeByValueReturnsError() {
	err := configurationTest.InitializeByValue(invalidURL1, accountPath)
	ts.ErrorContains(err, "parse")
	ts.Empty(configurationTest.BaseURL())
	ts.Empty(configurationTest.AccountPath())
}

func (ts *TSConfiguration) TestInvalidBaseURL1ReturnsError() {
	url, err := configurationTest.parseRawBaseURL(invalidURL1)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL2ReturnsError() {
	url, err := configurationTest.parseRawBaseURL(invalidURL2)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestInvalidBaseURL3ReturnsError() {
	url, err := configurationTest.parseRawBaseURL(invalidURL3)
	ts.ErrorContains(err, "failed parsing rawBaseURL:")
	ts.Nil(url)
}

func (ts *TSConfiguration) TestValidInitializeByEnvReturnsNoError() {
	_, ok1 := os.LookupEnv(baseURLEnvKey)
	if !ok1 {
		os.Setenv(baseURLEnvKey, rawBaseURL)
		defer os.Unsetenv(baseURLEnvKey)
	}
	_, ok2 := os.LookupEnv(accountPathEnvKey)
	if !ok2 {
		os.Setenv(accountPathEnvKey, accountPath)
		defer os.Unsetenv(accountPathEnvKey)
	}

	err := configurationTest.InitializeByEnv()
	ts.NoError(err)
	ts.NotEmpty(configurationTest.baseURL)
	ts.NotEmpty(configurationTest.accountPath)
}

func (ts *TSConfiguration) TestInvalidInitializeByEnvReturnsError() {
	err := configurationTest.InitializeByEnv()
	ts.ErrorContains(err, "failed to get BASE_URL from environment variables")
	ts.Empty(configurationTest.baseURL)
	ts.Empty(configurationTest.accountPath)
}

func (ts *TSConfiguration) TestInvalidInitializedByEnvReturnsErrorOnAccountPath() {
	_, ok1 := os.LookupEnv(baseURLEnvKey)
	if !ok1 {
		os.Setenv(baseURLEnvKey, rawBaseURL)
		defer os.Unsetenv(baseURLEnvKey)
	}
	err := configurationTest.InitializeByEnv()
	ts.Error(err)
}
