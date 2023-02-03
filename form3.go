package form3

import (
	"path"

	"github.com/AdanJSuarez/form3/internal/account"
	"github.com/AdanJSuarez/form3/internal/validation"
	"github.com/AdanJSuarez/form3/model"
)

type Account interface {
	Create(data model.Data) (model.Data, error)
	Fetch(accountID string) (model.Data, error)
	// Delete(accountID string) (model.Data, error)
}

type Form3 struct {
	url     string
	account Account
}

// New returns a instance of Form3 client. Returns an error if the URL is wrong.
// Configuration should be set in this step in a real application.
func New(form3URL string) (*Form3, error) {
	_, err := validation.NewValidation(form3URL)
	if err != nil {
		return nil, err
	}

	f3 := &Form3{url: form3URL}
	return f3, nil
}

func (f *Form3) Account(accountURL string) Account {
	url := path.Join(f.url, accountURL)
	f.account = account.New(url)
	return f.account
}

// func (f *Form3) Create(data model.Data) (model.Data, error) {
// 	return internal.Create(f.url, data)
// }

// func (f *Form3) Get(accountID string) (model.Data, error) {
// 	return internal.Fetch(f.url, accountID)
// }
