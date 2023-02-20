package integration

import "github.com/AdanJSuarez/form3/pkg/model"

var (
	attributeUK = model.Attributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "123456",
		BankIDCode:   "GBDSC",
		Bic:          "EXMPLGB2XXX",
		Name:         []string{"a", "b"},
	}
	dataUK = model.Data{
		ID:             generateAccountUUID(),
		OrganizationID: organizationID,
		Type:           "accounts",
		Version:        0,
		Attributes:     attributeUK,
	}
	dataModelUK = model.DataModel{
		Data: dataUK,
	}
)

// It should create a valid account with no errors
func (ts *TSIntegration) TestCreateAccountUK() {
	dataModelTest = dataModelUK
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
}
