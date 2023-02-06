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
	AccountURL() string
	OrganizationID() string
}

type Form3 struct {
	configuration configurationForm3
	account       Account
}

// New returns a instance of Form3 client. Returns an error if the URL is wrong.
// Configuration should be set in this step in a real application.
func New(baseURL, organizationID string) (*Form3, error) {
	configuration, err := configuration.New(baseURL, organizationID)
	if err != nil {
		return nil, err
	}

	return initializeForm3(configuration)
}

func NewByYaml() (*Form3, error) {
	configuration, err := configuration.NewByYaml()
	if err != nil {
		return nil, err
	}

	return initializeForm3(configuration)
}

func NewByEnv() (*Form3, error) {
	configuration, err := configuration.NewByEnv()
	if err != nil {
		return nil, err
	}

	return initializeForm3(configuration)
}

// Account returns an initialized pointer to an object the implement SetAccountConfiguration interface
func (f *Form3) Account() Account {
	return f.account
}

func initializeForm3(configuration *configuration.Configuration) (*Form3, error) {
	f3 := &Form3{
		configuration: configuration,
	}
	accountURL := configuration.AccountURL()
	f3.account = account.New(accountURL)
	return f3, nil
}
