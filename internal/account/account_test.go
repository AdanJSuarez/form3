package account

import (
	"net/url"
	"testing"

	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/stretchr/testify/suite"
)

const (
	baseURL        = "http://fakeurl.com"
	accountPath    = "/v1/origanisation/"
	uuidTest       = "123e4567-e89b-12d3-a456-426614174000"
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
)

var (
	accountTest    *Account
	baseURLTest, _ = url.Parse(baseURL)
	dataModelTest  = model.DataModel{
		Data: model.Data{
			ID:             uuidTest,
			OrganizationID: organizationID,
			Type:           "accounts",
			Attributes:     dataAttributesTest,
		},
	}
	dataAttributesTest = model.Attributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "123456",
		BankIDCode:   "GBDSC",
		Bic:          "EXMPLGB2XXX",
	}
)

type TSAccount struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSAccount))
}

func (ts *TSAccount) BeforeTest(_, _ string) {
	accountTest = New(*baseURLTest, accountPath)
	ts.IsType(new(Account), accountTest)
}

func (ts *TSAccount) TestCreateValidDataModel() {

}
