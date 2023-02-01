package internal

import (
	"net/url"
)

// NewValidation returns an error if URL or port are not valid.
func NewValidation(URL string) (*url.URL, error) {
	url, err := url.ParseRequestURI(URL)
	if err != nil {
		return nil, err
	}

	return url, nil
}
