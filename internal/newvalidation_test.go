package internal

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	validURL    = "https://api.fakeaddress/fake"
	invalidURL1 = "https//api.fakeaddress/"
	invalidURL2 = ""
)

type TSForm3 struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSForm3))
}

func (ts *TSForm3) BeforeTest(_, _ string) {

}

func (ts *TSForm3) TestValidURL() {
	_, err := NewValidation(validURL)
	ts.NoError(err)
}

func (ts *TSForm3) TestInvalidValidURL1() {
	URL := ""
	_, err := NewValidation(URL)
	ts.Error(err)
}

func (ts *TSForm3) TestInvalidValidURL3() {
	_, err := NewValidation(invalidURL1)
	ts.Error(err)
}
