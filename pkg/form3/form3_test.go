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
	mockConfiguration *MockConfiguration
	mockAccount       *MockAccount
)

type TSForm3 struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSForm3))
}

func (ts *TSForm3) BeforeTest(_, _ string) {
	mockConfiguration = NewMockConfiguration(ts.T())
	mockAccount = NewMockAccount(ts.T())
	form3Test = New()
	ts.IsType(new(Form3), form3Test)
	form3Test.configuration = mockConfiguration
	form3Test.account = mockAccount
}

func (ts *TSForm3) TestValidConfiguration() {
	ts.NoError(err)
}

func (ts *TSForm3) TestValidConfigurationNotNil() {
	ts.NotNil(form3Test)
}

func (ts *TSForm3) TestInvalidConfiguration() {
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"),
		mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration

	err = f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.Error(err)
}

func (ts *TSForm3) TestValidAccountObject() {
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)

	err = form3Test.ConfigurationByValue("fakeURL", accountPath)
	ts.Nil(err)
	account := form3Test.Account()
	ts.NotNil(account)
}

func (ts *TSForm3) TestInvalidAccountObject() {
	mockConfiguration.On("InitializeByValue", mock.AnythingOfType("string"),
		mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(fmt.Errorf("fake error"))
	f3Test := New()
	f3Test.configuration = mockConfiguration

	err = f3Test.ConfigurationByValue(rawBaseURLTest, accountPath)
	ts.NotNil(err)
	ts.Nil(f3Test.Account())
}

func (ts *TSForm3) TestInvalidConfigurationByEnv() {
	mockConfiguration = new(MockConfiguration)
	mockConfiguration.On("InitializeByEnv").Return(fmt.Errorf("not implemented"))
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)
	form3Test.configuration = mockConfiguration
	err := form3Test.ConfigurationByEnv()
	ts.NotNil(err)
}

func (ts *TSForm3) TestValidConfigurationByEnv() {
	mockConfiguration.On("InitializeByEnv").Return(nil)
	mockConfiguration.On("AccountPath").Return(accountPath)
	mockConfiguration.On("BaseURL").Return(baseURLTest)

	err := form3Test.ConfigurationByEnv()
	ts.Nil(err)
	ts.NotEmpty(form3Test.Account())
}

func (ts *TSForm3) TestCreateAccount() {
	mockAccount.On("Create", mock.Anything).Return(dataTest1, nil)
	account := form3Test.Account()
	data, err := account.Create(model.DataModel{})
	ts.NoError(err)
	ts.Equal(dataTest1, data)
}

func (ts *TSForm3) TestCreateAccountError() {
	mockAccount.On("Create", mock.Anything).Return(model.DataModel{}, fmt.Errorf("fakeError"))
	account := form3Test.Account()
	data, err := account.Create(model.DataModel{})
	ts.Error(err)
	ts.Empty(data)
}

func (ts *TSForm3) TestFetchAccount() {
	mockAccount.On("Fetch", mock.Anything).Return(dataTest1, nil)
	account := form3Test.Account()
	data, err := account.Fetch("FakeID")
	ts.NoError(err)
	ts.Equal(dataTest1, data)
}

func (ts *TSForm3) TestFetchAccountError() {
	mockAccount.On("Fetch", mock.Anything).Return(model.DataModel{}, fmt.Errorf("status code 404: fakeErrorOnFetch"))
	account := form3Test.Account()
	data, err := account.Fetch("FakeID")
	ts.ErrorContains(err, "status code 404:")
	ts.Empty(data)
}

func (ts *TSForm3) TestDeleteAccount() {
	mockAccount.On("Delete", mock.Anything, mock.Anything).Return(nil)
	account := form3Test.Account()
	err := account.Delete("fakeID", 0)
	ts.NoError(err)
}

func (ts *TSForm3) TestDeleteAccountError() {
	mockAccount.On("Delete", mock.Anything, mock.Anything).Return(fmt.Errorf("status code 404: fakeError"))
	account := form3Test.Account()
	err := account.Delete("fakeID", 7)
	ts.ErrorContains(err, "status code 404:")
}
