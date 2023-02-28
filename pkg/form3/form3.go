package form3

import (
	"github.com/AdanJSuarez/form3/internal/configuration"
	"github.com/AdanJSuarez/form3/pkg/account"
)

type Form3 struct {
	configuration Configuration
	account       *account.Account
}

/*
New returns a initialized pointer of Form3

This is the entry point of the library.
*/
func New() *Form3 {
	return &Form3{
		configuration: configuration.New(),
	}
}

/*
ConfigurationByValue initializes form3 with the parameters passed, it returns
nil if success, but an error otherwise.

Example: form3.ConfigurationByValue("https://api.form3.tech", "/v1/organisation/accounts")

For baseURL and accountPath consult Form3 API documentation.
*/
func (f *Form3) ConfigurationByValue(baseURL, accountPath string) error {
	if err := f.configuration.InitializeByValue(baseURL, accountPath); err != nil {
		return err
	}

	f.initializeForm3()

	return nil
}

/*
ConfigurationByEnv initializes form3 by reading the configuration values from the
environmental variables.

Example: BASE_URL=https://api.form3.tech ACCOUNT_PATH=/v1/organisation/accounts

For baseURL and accountPath information consult Form3 API documentation.
*/
func (f *Form3) ConfigurationByEnv() error {
	if err := f.configuration.InitializeByEnv(); err != nil {
		return err
	}

	f.initializeForm3()

	return nil
}

/*
Account returns a pointer of account.Account. It requires the configuration to
be previously set, either by value or by env. It will return nil otherwise.

For account.Account consult its documentation, and Form3 API documentation.
*/
func (f *Form3) Account() *account.Account {
	return f.account
}

func (f *Form3) initializeForm3() {
	f.account = account.New(f.configuration)
}
