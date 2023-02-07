package form3

import "github.com/AdanJSuarez/form3/pkg/model"

//go:generate mockery --inpackage --name=Account
//go:generate mockery --inpackage --name=configurationForm3

type Account interface {
	Create(data model.DataModel) (model.DataModel, error)
	Fetch(accountID string) (model.DataModel, error)
	Delete(accountID string, version int) error
}

type f3Configuration interface {
	InitializeByValue(baseURL, accountPath, organizationID string) error
	InitializeByYaml() error
	InitializeByEnv() error
	AccountURL() string
}
