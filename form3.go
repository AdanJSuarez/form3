package form3

import "errors"

type Form3 struct {
	URL  string
	Port uint16
}

func New(URL string, Port uint16) (*Form3, error) {
	return nil, errors.New("not implemented")
}
