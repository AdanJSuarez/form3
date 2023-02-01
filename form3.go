package form3

import (
	"errors"
	"net/url"

	"github.com/AdanJSuarez/form3/internal"
)

type Form3 struct {
	URL *url.URL
}

// New returns a instance of Form3 client. Returns an error if the URL is wrong.
func New(URL string) (*Form3, error) {
	url, err := internal.NewValidation(URL)
	if err != nil {
		return nil, err
	}

	f3 := &Form3{URL: url}
	return f3, nil
}

func Create() error {
	return errors.New("not implemented")
}
