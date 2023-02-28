package integration

import (
	"log"
	"testing"

	"github.com/AdanJSuarez/form3/pkg/account"
	"github.com/AdanJSuarez/form3/pkg/form3"
	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

const (
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	baseAPIURL     = "http://accountapi:8080"
	accountPath    = "/v1/organisation/accounts"
	fakeIBAN       = "ES2317002001280000001200527600"
)

var (
	f3Test        *form3.Form3
	accountTest   *account.Account
	uuids         = make(map[string]struct{})
	dataModelTest model.DataModel
	attributeBE   = model.Attributes{
		Country:      "BE",
		BankID:       "ABC",
		BankIDCode:   "BE",
		BaseCurrency: "EUR",
		Name:         []string{"John Doe"},
	}
	dataBE = model.Data{
		OrganizationID: organizationID,
		Type:           "accounts",
		Version:        0,
		Attributes:     attributeBE,
	}
	attributeUK = model.Attributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "123456",
		BankIDCode:   "GBDSC",
		Bic:          "EXMPLGB2XXX",
		Name:         []string{"a", "b"},
	}
	dataUK = model.Data{
		OrganizationID: organizationID,
		Type:           "accounts",
		Version:        0,
		Attributes:     attributeUK,
	}
	dataModelUK = model.DataModel{
		Data: dataUK,
	}
	dataModelBE = model.DataModel{
		Data: dataBE,
	}
)

type TSIntegration struct{ suite.Suite }

func TestRunTSIntegration(t *testing.T) {
	suite.Run(t, new(TSIntegration))
}

func (ts *TSIntegration) BeforeTest(_, _ string) {
	dataModelTest = dataModelUK
	f3Test = form3.New()
	if err := f3Test.ConfigurationByEnv(); err != nil {
		log.Printf("Error in Configuration: %v", err)
		return
	}
	accountTest = f3Test.Account()
}

func (ts *TSIntegration) AfterTest(_, _ string) {
	for id := range uuids {
		log.Printf("Deleting %s", id)
		accountTest.Delete(id, 0)
		delete(uuids, id)
	}
}

// It should connect, and return an error because of the wrong account "ID".
func (ts *TSIntegration) TestConfigurationByValue() {
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(baseAPIURL, accountPath); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	accountTest = f3Test.Account()
	data, err := accountTest.Fetch("xxxxxx")
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

// It should connect but returns an error because of the wrong Account path.
func (ts *TSIntegration) TestInvalidConfigurationByValueWrongPath() {
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(baseAPIURL, "/organisation/account"); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	accountTest = f3Test.Account()
	data, err := accountTest.Fetch("xxxxxx")
	ts.ErrorContains(err, "status code 404")
	ts.Empty(data)
}

// It should not create an account with empty "DataModel".
func (ts *TSIntegration) TestFailToCreateAccountEmptyDataModel() {
	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account without "Data".
func (ts *TSIntegration) TestFailToCreateAccountEmptyData() {
	data, err := accountTest.Create(model.DataModel{Data: model.Data{}})
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account with an already used "ID".
func (ts *TSIntegration) TestFailToCreateAccountSameID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.Equal(dataModelTest, data)
	data, err = accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 409")
	ts.Empty(data)
}

// It should not create an account without "ID".
func (ts *TSIntegration) TestFailToCreateAccountWithoutID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.ErrorContains(err, "id in body is required")
	ts.Empty(data)
}

// It should not create an account without several fields not included: "ID", "Country" and "Name"
// and should return the list of requirements.
func (ts *TSIntegration) TestFailToCreateAccountWithoutIDCountryName() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = ""
	dataModelTest.Data.Attributes.Country = ""
	dataModelTest.Data.Attributes.Name = nil
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.ErrorContains(err, "id in body is required")
	ts.ErrorContains(err, "country in body is required")
	ts.ErrorContains(err, "name in body is required")
	ts.Empty(data)
}

// It should not create an account with an "ID" that is not a UUID well formatted.
func (ts *TSIntegration) TestFailToCreateAccountWithoutCorrectID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = "XXXXX-XXXXX-333"
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account without "organizationID"
func (ts *TSIntegration) TestFailToCreateAccountWithoutOrgID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.OrganizationID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account with an incorrect "organizationID"
func (ts *TSIntegration) TestFailToCreateAccountWithoutCorrectOrgID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.OrganizationID = "ZZZZ-ZZZZZ"
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account without "type"
func (ts *TSIntegration) TestFailCreateAccountWithoutType() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Type = ""
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with an incorrect "type"
func (ts *TSIntegration) TestFailCreateAccountWithIncorrectType() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Type = "AdanJSuarez"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with any "version" number
func (ts *TSIntegration) TestCreateAccountWithAnyVersion() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Version = 1000000000
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(int64(0), data.Data.Version)
}

// Attributes ////

// It should not create an account without "attributes".
func (ts *TSIntegration) TestFailToCreateAccountWithoutAttributes() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes = model.Attributes{}
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account with wrong "AccountClassification"
func (ts *TSIntegration) TestFailCreateAccountWithWrongAccountClassification() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountClassification = "ssss"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with "AccountClassification"
func (ts *TSIntegration) TestCreateAccountWithAccountClassification() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountClassification = "Personal"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.AccountClassification, data.Data.Attributes.AccountClassification)
}

// It should create an account with true "AccountMatchingOptOut"
func (ts *TSIntegration) TestCreateAccountWithTrueAccountMatching() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountMatchingOptOut = true
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should create an account when "Account Number" is included.
func (ts *TSIntegration) TestAccountCreateBEWithAccountNumber() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountNumber = "1234567"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.AccountNumber, data.Data.Attributes.AccountNumber)
}

// It should not create an account when "Account Number" is invalid.
func (ts *TSIntegration) TestFailAccountCreateBEWithWrongAccountNumber() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountNumber = "abcdefg"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with an "AlternativeName"
func (ts *TSIntegration) TestCreateAccountWithAlternativeName() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AlternativeNames = []string{"Jane Doe the Third"}
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.Name, data.Data.Attributes.Name)
}

// It should create an account with an empty "AlternativeName"
func (ts *TSIntegration) TestCreateAccountWithEmptyAlternativeName() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AlternativeNames = []string{}
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Nil(data.Data.Attributes.AlternativeNames)
}

// It should create an account with an nil "AlternativeName"
func (ts *TSIntegration) TestCreateAccountWithNilAlternativeName() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AlternativeNames = nil
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Nil(data.Data.Attributes.AlternativeNames)
}

// It should create an account with no "BankID"
func (ts *TSIntegration) TestCreateAccountWithoutBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BankID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.BankID, data.Data.Attributes.BankID)
}

// It should not create an account with an incorrect "BankID"
func (ts *TSIntegration) TestFailCreateAccountWithIncorrectBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BankID = "sssssss"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account when "BankIDCode" isn't included
func (ts *TSIntegration) TestFailCreateAccountBEWithoutBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BankIDCode = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.BankIDCode, data.Data.Attributes.BankIDCode)
}

// It should not create an account when "BankIDCode" isn't correct
func (ts *TSIntegration) TestFailCreateAccountBEWithIncorrectBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BankIDCode = "11111"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account without "BaseCurrency" for a none EUR country.
func (ts *TSIntegration) TestCreateAccountWithoutBaseCurrency() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.BaseCurrency, data.Data.Attributes.BaseCurrency)
}

// It should create an account without "BaseCurrency" for a EUR country
func (ts *TSIntegration) TestFailCreateAccountWithoutBaseCurrency() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.BaseCurrency, data.Data.Attributes.BaseCurrency)
}

// It should not create an account with the invalid "BaseCurrency" for a EUR country.
func (ts *TSIntegration) TestFailCreateAccountWithInvalidBaseCurrency() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = "333"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with the invalid "BaseCurrency" for a not EUR country.
func (ts *TSIntegration) TestFailCreateAccountWithInvalidBaseCurrencyUK() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = "333"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account when "BIC" is included.
func (ts *TSIntegration) TestAccountCreateBEWithBIC() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Bic = "EBAXBEBB"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.Bic, data.Data.Attributes.Bic)
}

// It should not create an account when the "BIC" doesn't meet the requirements
func (ts *TSIntegration) TestFailCreateAccountBEWithWrongBIC() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Bic = "12345"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account without "Country"
func (ts *TSIntegration) TestFailCreateAccountWithoutCountry() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Country = ""
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with wrong "Country" formatted
func (ts *TSIntegration) TestFailCreateAccountWithWrongCountryFormat() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Country = "xxxx"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account when "IBAN" is included.
func (ts *TSIntegration) TestAccountCreateBEWithIBAN() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Iban = fakeIBAN
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(fakeIBAN, data.Data.Attributes.Iban)
}

// It should not create an account when "IBAN" is not correct.
func (ts *TSIntegration) TestFailAccountCreateBEWithIncorrectIBAN() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Iban = "AABB00000000"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with true "JoinAccount"
func (ts *TSIntegration) TestCreateAccountJoinAccountTrue() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.JointAccount = true
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.JointAccount, data.Data.Attributes.JointAccount)
}

// It should not create an account without "Name" in the attributes.
func (ts *TSIntegration) TestCreateAccountWithoutName() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Name = nil
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with empty "Name" in the attributes.
func (ts *TSIntegration) TestCreateAccountWithEmptyName() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Name = []string{}
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with "SecondaryIdentification"
func (ts *TSIntegration) TestCreateAccountWithSecondaryIdentification() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.SecondaryIdentification = "xxxxxx"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.SecondaryIdentification, data.Data.Attributes.SecondaryIdentification)
}

// It should not create an account with wrong "Status"
func (ts *TSIntegration) TestFailCreateAccountWithWrongStatus() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Status = "xxxx"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account with pending "Status"
func (ts *TSIntegration) TestCreateAccountWithStatus() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Status = "pending"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.Status, data.Data.Attributes.Status)
}

// It should create an account with "Switched"
func (ts *TSIntegration) TestCreateAccountWithSwitched() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Switched = true
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
	ts.Equal(dataModelTest.Data.Attributes.Switched, data.Data.Attributes.Switched)
}

// It should create an account
func (ts *TSIntegration) TestAccountCreateBE() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest.Data.ID, data.Data.ID)
}

// It should create more than one account
func (ts *TSIntegration) TestAccountsCreateBE() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	data1, err1 := accountTest.Create(dataModelTest)
	ts.NoError(err1)
	ts.NotEmpty(data1)
	ts.Equal(dataModelTest.Data.ID, data1.Data.ID)

	dataModelTest.Data.ID = generateAccountUUID()
	data2, err2 := accountTest.Create(dataModelTest)
	ts.NoError(err2)
	ts.NotEmpty(data2)
	ts.Equal(dataModelTest.Data.ID, data2.Data.ID)

	dataModelTest.Data.ID = generateAccountUUID()
	data3, err3 := accountTest.Create(dataModelTest)
	ts.NoError(err3)
	ts.NotEmpty(data3)
	ts.Equal(dataModelTest.Data.ID, data3.Data.ID)
}

// It should fetch an existing account
func (ts *TSIntegration) TestFetchExistingAccount() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	accountTest.Create(dataModelTest)

	data, err := accountTest.Fetch(dataModelTest.Data.ID)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(dataModelTest, data)
}

// It should delete an existing account
func (ts *TSIntegration) TestDeleteExistingAccount() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	accountTest.Create(dataModelTest)

	err := accountTest.Delete(dataModelTest.Data.ID, 0)
	ts.NoError(err)
}

func generateAccountUUID() string {
	id := uuid.New()
	uuidString := id.String()
	uuids[uuidString] = struct{}{}
	return uuidString
}
