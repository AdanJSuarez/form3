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
	healthCheckNumOfTries = 5
	healthCheckInterval   = 5 * time.Second
	organizationID        = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	baseAPIURL            = "http://localhost:8080" //"http://accountapi:8080"
	accountPath           = "/v1/organisation/accounts"
)

var (
	f3Test      *form3.Form3
	accountTest form3.Account
	uuids       = []string{}
	jitUUID     = generateUUID()
	attribute   = model.Attributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "123456",
		BankIDCode:   "GBDSC",
		Bic:          "EXMPLGB2XXX",
		Name:         []string{"a", "b"},
	}
	dataTest = model.Data{
		ID:             jitUUID,
		OrganizationID: organizationID,
		Type:           "accounts",
		Version:        0,
		Attributes:     attribute,
	}
	dataModel = model.DataModel{
		Data: dataTest,
	}
	dataModelTest model.DataModel
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
		log.Printf("error on health-check: %v", err)
		time.Sleep(healthCheckInterval)
		return false
	}
	return true
}

// TODO: Change ConfigurationByValue for ConfigurationByEnv
func (ts *TSIntegration) BeforeTest(_, _ string) {
	dataModelTest = dataModel
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(baseAPIURL, accountPath); err != nil { //if err := f3Test.ConfigurationByEnv(); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	accountTest = f3Test.Account()
}

func (ts *TSIntegration) AfterTest(_, _ string) {
	for _, id := range uuids {
		log.Printf("Deleting %s", id)
		accountTest.Delete(id, 0)
	}
}

// It should connect, and return an error 400 because of the wrong account ID.
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
func (ts *TSIntegration) TestInvalidConfigurationByValue2() {
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

// It should not connect at all with the wrong baseURL.
func (ts *TSIntegration) TestInvalidConfigurationByValue1() {
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue("http://localhost:5999", accountPath); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	accountTest = f3Test.Account()
	data, err := accountTest.Fetch("xxxxxx")
	ts.ErrorContains(err, "connect: connection refused")
	ts.Empty(data)
}

// It should create a valid account with no errors
func (ts *TSIntegration) TestCreateAccount() {
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.Equal(dataModelTest, data)
}

// It should returns a 400 when trying to create an account with incomplete info.
func (ts *TSIntegration) TestFailToCreateAccountEmptyData() {
	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

// It should fail trying to create an account with an already used UUID.
func (ts *TSIntegration) TestFailToCreateAccountSameUUID() {
	dataModelTest.Data.ID = generateUUID()
	_, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	_, err = accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 409: resource has already been created")
}

// It should fail if we try to create an account without ID
func (ts *TSIntegration) TestFailToCreateAccountWithoutID() {
	dataModelTest.Data.ID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It should fail if we try to create an account without organizationID
func (ts *TSIntegration) TestFailToCreateAccountWithoutOrgID() {
	dataModelTest.Data.OrganizationID = ""
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// TODO: Do the same for none account creation
// It shouldn't fail if we don't provide "type" in account creation
func (ts *TSIntegration) TestCreateAccountWithoutType() {
	dataModelTest.Data.Type = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should fail if we don't pass the "attributes" in account creation
func (ts *TSIntegration) TestFailToCreateAccountWithoutAttributes() {
	dataModelTest.Data.Attributes = model.Attributes{}
	data, err := accountTest.Create(dataModelTest)
	ts.ErrorContains(err, "status code 400")
	ts.Empty(data)
}

// It shouldn't fail if we don't pass "name" in the attributes.
func (ts *TSIntegration) TestCreateAccountWithoutName() {
	dataModelTest.Data.Attributes.Name = nil
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should not fail without base_currency for none EUR country
func (ts *TSIntegration) TestCreateAccountWithoutBaseCurrency() {
	dataModelTest.Data.Attributes.BaseCurrency = ""
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}

// It should fail without base_currency for a EUR country
func (ts *TSIntegration) TestFailCreateAccountWithoutBaseCurrency() {
	attribute := model.Attributes{
		Country:    "ES",
		BankID:     "12345678",
		BankIDCode: "ESNC",
		Name:       []string{"a", "b"},
	}
	dataModelTest.Data.Attributes = attribute
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

func generateUUID() string {
	id := uuid.New()
	uuidString := id.String()
	uuids = append(uuids, uuidString)
	return uuidString
}
