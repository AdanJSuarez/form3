package form3

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURLTest = "https://api.fakeaddress/fake:8080"
	accountPath    = "/v1/organisation/accounts"
)

var (
	form3Test         *Form3
	baseURLTest, _    = url.ParseRequestURI(rawBaseURLTest)
	mockConfiguration *MockConfiguration
)

type TSForm3 struct{ suite.Suite }

func TestRunForm3Suite(t *testing.T) {
	suite.Run(t, new(TSForm3))
}

func (ts *TSForm3) BeforeTest(_, _ string) {
	mockConfiguration = NewMockConfiguration(ts.T())
	form3Test = New()
	ts.IsType(new(Form3), form3Test)
	form3Test.configuration = mockConfiguration
}

func (ts *TSForm3) TestValidConfigurationReturnsNoError() {
	mockConfiguration.On("InitializeByValue", mock.Anything, mock.Anything).Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)
	err := form3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.NoError(err)
}

func (ts *TSForm3) TestInvalidConfigurationReturnsError() {
	mockConfiguration.On("InitializeByValue", mock.Anything,
		mock.Anything).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration

	err := f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.Error(err)
}

func (ts *TSForm3) TestOnValidConfigReturnsAccountObject() {
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)

	err := form3Test.ConfigurationByValue("fakeURL", accountPath)
	ts.NoError(err)
	account := form3Test.Account()
	ts.NotNil(account)
}

func (ts *TSForm3) TestOnInvalidConfigReturnsNilAccount() {
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration

	err := f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.Error(err)
	ts.Nil(f3Test.Account())
}

func (ts *TSForm3) TestInvalidConfigurationByEnvReturnsError() {
	mockConfiguration = new(MockConfiguration)
	mockConfiguration.On("InitializeByEnv").Return(fmt.Errorf("not implemented"))
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByEnv()
	ts.Error(err)
}

func (ts *TSForm3) TestValidConfigurationByEnvReturnsAccount() {
	mockConfiguration.On("InitializeByEnv").Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)

	err := form3Test.ConfigurationByEnv()
	ts.NoError(err)
	ts.NotNil(form3Test.Account())
}
