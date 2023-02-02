package form3

import (
	"github.com/AdanJSuarez/form3/internal"
	"github.com/AdanJSuarez/form3/internal/validation"
	"github.com/AdanJSuarez/form3/model"
)

type Form3 struct {
	url string
}

// New returns a instance of Form3 client. Returns an error if the URL is wrong.
// Configuration should be set in this step in a real application.
func New(url string) (*Form3, error) {
	_, err := validation.NewValidation(url)
	if err != nil {
		return nil, err
	}

	f3 := &Form3{url: url}
	return f3, nil
}

func (f *Form3) Create(accData model.AccountData) (model.AccountData, error) {
	return internal.Create(f.url, accData)
}
