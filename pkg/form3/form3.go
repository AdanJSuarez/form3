package form3

import (
	"github.com/AdanJSuarez/form3/internal/account"
	"github.com/AdanJSuarez/form3/internal/configuration"
)

type Form3 struct {
	configuration f3Configuration
	account       Account
}

func New() *Form3 {
	return &Form3{
		configuration: configuration.New(),
	}
}

func (f *Form3) ConfigurationByValue(rawBaseURL, accountPath string) error {
	if err := f.configuration.InitializeByValue(rawBaseURL, accountPath); err != nil {
		return err
	}
	f.initializeForm3()
	return nil
}

// ConfigurationByYaml returns error: Not implemented (for simplicity)
func (f *Form3) ConfigurationByYaml() error {
	if err := f.configuration.InitializeByYaml(); err != nil {
		return err
	}
	f.initializeForm3()
	return nil
}

// ConfigurationByEnv returns error: Not implemented (for simplicity)
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
	baseURL := f.configuration.BaseURL()
	accountPath := f.configuration.AccountPath()
	f.account = account.New(baseURL, accountPath)
}
