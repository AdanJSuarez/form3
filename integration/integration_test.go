package integration

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/AdanJSuarez/form3/pkg/form3"
	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// TODO: Change baseAPIURL for accountapi
const (
	healthCheckNumOfTries = 10
	healthCheckInterval   = 5 * time.Second
	organizationID        = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	baseAPIURL            = "http://localhost:8080" //"http://accountapi:8080"
	accountPath           = "/v1/organisation/accounts"
	fakeIBAN              = "ES2317002001280000001200527600"
)

var (
	f3Test        *form3.Form3
	accountTest   form3.Account
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

func (ts *TSIntegration) SetupSuite() {
	ts.startHealthCheck()
}

func (ts *TSIntegration) startHealthCheck() {
	for idx := 0; idx < healthCheckNumOfTries; idx++ {
		log.Printf("Starting health-check num. %d", idx+1)
		if ts.getHealthCheck() {
			log.Printf("Health-check num. %d success", idx+1)
			return
		}
	}
	log.Fatal("==> Server not ready. Integration tests cannot run! <==")
}

func (ts *TSIntegration) getHealthCheck() bool {
	stringConnection := baseAPIURL + "/v1/health"
	_, err := http.Get(stringConnection)
	if err != nil {
		log.Printf("Error during health-check: %v", err)
		time.Sleep(healthCheckInterval)
		return false
	}
	return true
}

// TODO: Change ConfigurationByValue for ConfigurationByEnv
func (ts *TSIntegration) BeforeTest(_, _ string) {
	dataModelTest = dataModelUK
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(baseAPIURL, accountPath); err != nil { //if err := f3Test.ConfigurationByEnv(); err != nil {
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

// It should connect, and return an error 400 because of the wrong account "ID".
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

// It should connect but returns a 404 because the wrong Account path.
func (ts *TSIntegration) TestInvalidConfigurationByValueWrongPath() {
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(baseAPIURL, "/organisation/account"); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	accountTest = f3Test.Account()
	data, err := accountTest.Fetch("xxxxxx")
	ts.ErrorContains(err, "status code 404: not found: trying to access a non-existent endpoint or resource")
	ts.Empty(data)
}

// It should not create an account with incomplete info.
func (ts *TSIntegration) TestFailToCreateAccountEmptyData() {
	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

// It should not create an account with an already used "ID".
func (ts *TSIntegration) TestFailToCreateAccountSameUUID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	_, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	_, err = accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 409:")
}

// It should not create an account without "ID".
func (ts *TSIntegration) TestFailToCreateAccountWithoutID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
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
	dataModelTest.Data.OrganizationID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account with an incorrect "organizationID"
func (ts *TSIntegration) TestFailToCreateAccountWithoutCorrectOrgID() {
	dataModelTest = dataModelUK
	dataModelTest.Data.OrganizationID = "ZZZZ-ZZZZZ"
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account without "type"
func (ts *TSIntegration) TestCreateAccountWithoutType() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Type = ""
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with an incorrect "type"
func (ts *TSIntegration) TestFailCreateAccountWithType1() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Type = "AdanJSuarez"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account without "attributes".
func (ts *TSIntegration) TestFailToCreateAccountWithoutAttributes() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Attributes = model.Attributes{}
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should not create an account without "name" in the attributes.
func (ts *TSIntegration) TestCreateAccountWithoutName() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Attributes.Name = nil
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with empty "name" in the attributes.
func (ts *TSIntegration) TestCreateAccountWithEmptyName() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Attributes.Name = []string{}
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account without base_currency for a none EUR country.
func (ts *TSIntegration) TestCreateAccountWithoutBaseCurrency() {
	dataModelTest = dataModelUK
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should create an account without base_currency for a EUR country
func (ts *TSIntegration) TestFailCreateAccountWithoutBaseCurrency() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BaseCurrency = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should not create an account with the invalid base_currency for a EUR country.
func (ts *TSIntegration) TestFailCreateAccountWithInvalidBaseCurrency() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.BaseCurrency = "333"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account with the invalid base_currency for a not EUR country.
func (ts *TSIntegration) TestFailCreateAccountWithInvalidBaseCurrencyUK() {
	dataModelTest = dataModelUK
	dataModelTest.Data.Attributes.BaseCurrency = "333"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account
func (ts *TSIntegration) TestAccountCreateBE() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should create an account with no "BankID"
func (ts *TSIntegration) TestCreateAccountWithoutBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.BankID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
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

// It should create an account when "BIC" is included.
func (ts *TSIntegration) TestAccountCreateBEWithBIC() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Bic = "EBAXBEBB"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should not create an account when the "BIC" doesn't meet the requirements
func (ts *TSIntegration) TestFailCreateAccountBEWithWrongBIC() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.Bic = "12345"
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
}

// It should not create an account when "BankIDCode" isn't correct
func (ts *TSIntegration) TestFailCreateAccountBEWithIncorrectBankID() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.BankIDCode = "11111"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account when "Account Number" is included.
func (ts *TSIntegration) TestAccountCreateBEWithAccountNumber() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.AccountNumber = "1234567"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.NotEmpty(data.Data.Attributes.AccountNumber)
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

// It should create an account when "IBAN" is included.
func (ts *TSIntegration) TestAccountCreateBEWithIBAN() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	dataModelTest.Data.Attributes.Iban = fakeIBAN
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.Equal(fakeIBAN, data.Data.Attributes.Iban)
}

// It should not create an account when "IBAN" is not correct.
func (ts *TSIntegration) TestFailAccountCreateBEWithIncorrectIBAN() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.Iban = "AABB00000000"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should fetch an existing account
func (ts *TSIntegration) TestFetchExistingAccount() {
	dataModelTest = dataModelBE
	dataModelTest.Data.ID = generateAccountUUID()
	accountTest.Create(dataModelTest)
	data, err := accountTest.Fetch(dataModelTest.Data.ID)
	ts.NoError(err)
	ts.NotEmpty(data)
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
