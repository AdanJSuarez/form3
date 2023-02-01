package form3

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	validURLTest   = "https://api.fakeaddress/fake:8080"
	invalidURLTest = ""
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
	form3Test, err = New(validURLTest)
}

func (ts *TSForm3) TestValidURLNoError() {
	ts.NoError(err)
}

func (ts *TSForm3) TestValidURLForm3NotNil() {
	ts.NotNil(form3Test)
}

func (ts *TSForm3) TestInvalidURLErrorandNilForm3() {
	f3Test, err := New(invalidURLTest)
	ts.Nil(f3Test)
	ts.Error(err)
}
