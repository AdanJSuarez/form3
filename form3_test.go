package form3

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	urlTest = "https://api.fakeaddress/fake:8080"
)

var (
	form3Test *Form3
	err       error
)

type TSForm3 struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSForm3))
}

func (ts *TSForm3) BeforeTest(_, _ string) {
	form3Test, err = New(urlTest)
}

func (ts *TSForm3) TestValidURL() {
	ts.NoError(err)
}
