package form3

import (
	"github.com/AdanJSuarez/form3/internal/account"
	"github.com/AdanJSuarez/form3/internal/configuration"
	"github.com/AdanJSuarez/form3/pkg/model"
)

//go:generate mockery --inpackage --name=Account
//go:generate mockery --inpackage --name=configurationForm3

type Account interface {
	Create(data model.DataModel) (model.DataModel, error)
	Fetch(accountID string) (model.DataModel, error)
	Delete(accountID string, version int) error
}

type configurationForm3 interface {
	InitializeByValue(baseURL, accountPath, organizationID string) error
	InitializeByYaml() error
	InitializeByEnv() error
	AccountURL() string
}

type Form3 struct {
	configuration configurationForm3
	account       Account
}

func New() *Form3 {
	return &Form3{
		configuration: configuration.New(),
	}
}

func (f *Form3) ConfigurationByValue(baseURL, accountPath, organizationID string) error {
	if err := f.configuration.InitializeByValue(baseURL, accountPath, organizationID); err != nil {
		return err
	}
	f.initializeForm3()
	return nil
}

func (f *Form3) ConfigurationByYaml() error {
	if err := f.configuration.InitializeByYaml(); err != nil {
		return err
	}
	f.initializeForm3()
	return nil
}

func (f *Form3) ConfigurationByEnv() error {
	if err := f.configuration.InitializeByEnv(); err != nil {
		return err
	}
	f.initializeForm3()
	return nil
}

func (f *Form3) Account() Account {
	return f.account
}

func (f *Form3) initializeForm3() {
	accountURL := f.configuration.AccountURL()
	f.account = account.New(accountURL)
}
