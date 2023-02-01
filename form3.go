package form3

import (
	"errors"

	"github.com/AdanJSuarez/form3/internal"
	"github.com/AdanJSuarez/form3/internal/validation"
)

type Form3 struct {
	url internal.URL
}

// New returns a instance of Form3 client. Returns an error if the URL is wrong.
func New(URL string) (*Form3, error) {
	_, err := validation.NewValidation(URL)
	if err != nil {
		return nil, err
	}

	f3 := &Form3{url: internal.URL(URL)}
	return f3, nil
}

func (f *Form3) Create() error {
	return errors.New("not implemented")
}
