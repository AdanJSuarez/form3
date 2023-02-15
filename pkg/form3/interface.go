package form3

import (
	"github.com/AdanJSuarez/form3/pkg/model"
)

//go:generate mockery --inpackage --name=Account
//go:generate mockery --inpackage --name=f3Configuration

type Account interface {
	// Create creates an bank account and returns the account values.
	// It returns an error otherwise.
	Create(data model.DataModel) (model.DataModel, error)
	// Fetch retrieves the account information for the specific account ID.
	// It returns an error otherwise.
	Fetch(accountID string) (model.DataModel, error)
	// Delete deletes an account by its ID and version number.
	// It returns an error otherwise.
	Delete(accountID string, version int) error
}
