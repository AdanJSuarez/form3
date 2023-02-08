package form3

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURLTest = "https://api.fakeaddress/fake:8080"
	invalidURLTest = ""
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	accountPath    = "/v1/organisation/accounts"
)

var (
	form3Test      *Form3
	err            error
	baseURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	dataTest1      = model.DataModel{
		Data: model.Data{
			ID:             "020cf7d8-01b9-461d-89d4-89d57fd0d998",
			OrganizationID: organizationID,
			Type:           "accounts",
			Version:        1,
			Attributes: model.Attributes{
				AlternativeNames: nil,
				BankID:           "123456",
				BankIDCode:       "GBDSC",
				Bic:              "EXMPLGB2XXX",
				Country:          "GB",
				Name:             []string{"Adan"},
			},
		},
	}
)

type TSForm3 struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSForm3))
}

func (ts *TSForm3) BeforeTest(_, _ string) {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)
	form3Test = New()
	ts.IsType(new(Form3), form3Test)
	form3Test.configuration = mockConfiguration
}

func (ts *TSForm3) TestValidConfiguration() {
	ts.NoError(err)
}

func (ts *TSForm3) TestValidConfigurationNotNil() {
	ts.NotNil(form3Test)
}

func (ts *TSForm3) TestInvalidConfiguration() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"),
		mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration
	err = f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.Error(err)
}

func (ts *TSForm3) TestValidAccountObject() {
	err = form3Test.ConfigurationByValue("fakeURL", accountPath)
	ts.Nil(err)
	account := form3Test.Account()
	ts.NotNil(account)
}

func (ts *TSForm3) TestInvalidAccountObject() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"),
		mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration
	err = f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.NotNil(err)
	ts.Nil(f3Test.Account())
}

func (ts *TSForm3) TestInvalidConfigurationByYaml() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByYaml").Return(fmt.Errorf("not implemented"))
	mockConfiguration.On("AccountPath", mock.Anything).Return(accountPath)
	mockConfiguration.On("BaseURL", mock.Anything).Return(baseURLTest)
	form3Test = New()
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByYaml()
	ts.NotNil(err)
}

func (ts *TSForm3) TestValidConfigurationByYaml() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByYaml").Return(nil)
	mockConfiguration.On("AccountPath", mock.Anything).Return(accountPath)
	mockConfiguration.On("BaseURL", mock.Anything).Return(baseURLTest)
	form3Test = New()
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByYaml()
	ts.Nil(err)
	ts.NotEmpty(form3Test.Account())
}

func (ts *TSForm3) TestInvalidConfigurationByEnv() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByEnv").Return(fmt.Errorf("not implemented"))
	mockConfiguration.On("AccountPath", mock.Anything).Return(accountPath)
	mockConfiguration.On("BaseURL", mock.Anything).Return(baseURLTest)
	form3Test = New()
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByEnv()
	ts.NotNil(err)
}

func (ts *TSForm3) TestValidConfigurationByEnv() {
	mockConfiguration := new(mockF3Configuration)
	mockConfiguration.On("InitializeByEnv").Return(nil)
	mockConfiguration.On("AccountPath", mock.Anything).Return(accountPath)
	mockConfiguration.On("BaseURL", mock.Anything).Return(baseURLTest)
	form3Test = New()
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByEnv()
	ts.Nil(err)
	ts.NotEmpty(form3Test.Account())
}
