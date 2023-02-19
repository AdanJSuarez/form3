package form3

import (
	"net/url"

	"github.com/AdanJSuarez/form3/pkg/model"
)

//go:generate mockery --inpackage --name=Account
//go:generate mockery --inpackage --name=Configuration

type Account interface {
	/*
		Create creates an bank account and returns the account values.
		It returns an error otherwise.

		For more reference about model.DataModel, please check form3 API documentation.
	*/
	Create(data model.DataModel) (model.DataModel, error)
	/*
		Fetch retrieves the account information for the specific account ID.
		It returns an error otherwise.

		For more reference about model.DataModel and accountID, please check form3 API documentation.
	*/
	Fetch(accountID string) (model.DataModel, error)
	/*
		Delete deletes an account by its ID and version number.
		It returns an error otherwise.

		For more reference about accountID and version, please check form3 API documentation.
	*/
	Delete(accountID string, version int) error
}

type Configuration interface {
	BaseURL() *url.URL
	AccountPath() string
	InitializeByValue(rawBaseURL, accountPath string) error
	InitializeByEnv() error
}
