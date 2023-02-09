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

const (
	healthCheckNumOfTries = 5
	healthCheckInterval   = 5 * time.Second
	organizationID        = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	localhostURL          = "http://host.docker.internal:8080"
	accountPath           = "/v1/organisation/accounts"
)

var (
	f3Test        *form3.Form3
	accountTest   form3.Account
	uuids         = []string{}
	jitUUID       = generateUUID()
	dataModelTest = model.DataModel{
		Data: model.Data{
			ID:             jitUUID,
			OrganizationID: organizationID,
			Type:           "accounts",
			Version:        0,
			Attributes: model.Attributes{
				Country: "GB",
				// BaseCurrency: "GBP",
				BankID:     "123456",
				BankIDCode: "GBDSC",
				Bic:        "EXMPLGB2XXX",
				Name:       []string{"a", "b"},
			},
		},
	}
	linksTest = model.Links{
		Self: accountPath + "/" + jitUUID,
	}
)

type TSIntegration struct{ suite.Suite }

func TestRunTSIntegration(t *testing.T) {
	suite.Run(t, new(TSIntegration))
}

func (ts *TSIntegration) SetupSuite() {
	ts.startHealthCheck()
}

func (ts *TSIntegration) BeforeTest(_, _ string) {
	f3Test = form3.New()
	if err := f3Test.ConfigurationByValue(localhostURL, accountPath); err != nil {
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

func (ts *TSIntegration) TestCreateAccount() {
	data, err := accountTest.Create(dataModelTest)
	dataModelTest.Links = linksTest
	ts.NoError(err)
	ts.Equal(dataModelTest, data)
}

func (ts *TSIntegration) TestThrottling() {
	// TODO: Implement TestThrottling
}

func (ts *TSIntegration) startHealthCheck() {
	for idx := 0; idx < healthCheckNumOfTries; idx++ {
		log.Printf("Starting health-check num. %d", idx+1)
		if ts.getRequest() {
			log.Printf("Health-check num. %d success", idx+1)
			return
		}
	}
	log.Fatal("==> Server not ready. Integration tests cannot run! <==")
}

func (ts *TSIntegration) getRequest() bool {
	stringConnection := localhostURL + "/v1/health"
	_, err := http.Get(stringConnection)
	if err != nil {
		log.Printf("error on health-check: %v", err)
		time.Sleep(healthCheckInterval)
		return false
	}
	return true
}

func generateUUID() string {
	id := uuid.New()
	uuidString := id.String()
	uuids = append(uuids, uuidString)
	return uuidString
}
