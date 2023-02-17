package integration

import (
	"github.com/AdanJSuarez/form3/pkg/model"
)

var (
	attributeBE = model.Attributes{
		Country:      "BE",
		BankID:       "ABC",
		BankIDCode:   "BE",
		BaseCurrency: "EUR",
		Name:         []string{"John Doe"},
	}
	dataBE = model.Data{
		ID:             generateAccountUUID(),
		OrganizationID: organizationID,
		Type:           "accounts",
		Version:        0,
		Attributes:     attributeBE,
	}
	dataModelBE = model.DataModel{
		Data: dataBE,
	}
)

// It should create an BE account
func (ts *TSIntegration) TestAccountCreateBE1() {
	dataModelTest = dataModelBE
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.NotEmpty(data.Data.Attributes.Iban)
	ts.NotEmpty(data.Data.Attributes.AccountNumber)
}

// It should create an BE account when we include BIC (Optional)
func (ts *TSIntegration) TestAccountCreateBE2() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.Bic = "EBAXBEBB"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.NotEmpty(data.Data.Attributes.Iban)
	ts.NotEmpty(data.Data.Attributes.AccountNumber)
}

// It should not create an account when the BIC doesn't meet the requirements
func (ts *TSIntegration) TestFailCreateAccountBE1() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.Bic = "12345"
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should not create an account when Bank ID code isn't included
func (ts *TSIntegration) TestFailCreateAccountBE2() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.BankIDCode = ""
	data, err := accountTest.Create(dataModelTest)
	ts.Error(err)
	ts.Empty(data)
}

// It should create an account when Account Number is included.
func (ts *TSIntegration) TestAccountCreateBE3() {
	dataModelTest = dataModelBE
	dataModelTest.Data.Attributes.AccountNumber = "1234567"
	data, err := accountTest.Create(dataModelTest)
	ts.NoError(err)
	ts.NotEmpty(data)
	ts.NotEmpty(data.Data.Attributes.Iban)
	ts.NotEmpty(data.Data.Attributes.AccountNumber)
}
