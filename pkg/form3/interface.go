package form3

import (
	"net/url"

	"github.com/AdanJSuarez/form3/pkg/model"
)

//go:generate mockery --inpackage --name=Account
//go:generate mockery --inpackage --name=f3Configuration

type Account interface {
	Create(data model.DataModel) (model.DataModel, error)
	Fetch(accountID string) (model.DataModel, error)
	Delete(accountID string, version int) error
}

type f3Configuration interface {
	BaseURL() *url.URL
	AccountPath() string
	InitializeByValue(rawBaseURL, accountPath string) error
	InitializeByYaml() error
	InitializeByEnv() error
}
