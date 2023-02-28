package handler

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	dataCodeMessage = `{
		"error_message": "Message parsing failed: Unexpected character (';' (code 34)): was expecting comma to separate Object entries ",
		"error_code": "d0a17902-63ed-4cb6-a8e8-fac5ca31b0b7"
	}`
	dataTypeDescription = `{
		"error": "invalid_grant",
		"error_description": "Wrong email or password."
	}`
	invalidData = "tenerifeLaMejorIsla8347349873er3"

	bodyCodeMessage     = io.NopCloser(bytes.NewBuffer([]byte(dataCodeMessage)))
	bodyTypeDescription = io.NopCloser(bytes.NewBuffer([]byte(dataTypeDescription)))
	invalidBody         = io.NopCloser(bytes.NewBuffer([]byte(invalidData)))
)

type TSError struct{ suite.Suite }

func TestRunErrorSuite(t *testing.T) {
	suite.Run(t, new(TSError))
}

func (ts *TSError) TestNewErrorReturnsCorrectly() {
	err := newError(777, fmt.Errorf("fake error in form3 because I am on the beach :)"))
	ts.ErrorContains(err, "status code 777:")
	ts.ErrorContains(err, "I am on the beach")
}

func (ts *TSError) TestTypeDescriptionErrorReturnsCorrectly() {
	err := newTypeDescriptionError(777, bodyTypeDescription)
	ts.ErrorContains(err, "status code 777:")
	ts.ErrorContains(err, "error: invalid_grant")
	ts.ErrorContains(err, "errorDescription: Wrong email or password.")
}

func (ts *TSError) TestInvalidBodyNotReturnsTypeDescriptionError() {
	err := newTypeDescriptionError(777, invalidBody)
	ts.ErrorContains(err, "status code 777:")
	ts.NotContains(err.Error(), "errorDescription:")
}

func (ts *TSError) TestCodeMessageErrorReturnsCorrectly() {
	err := newCodeMessageError(777, bodyCodeMessage)
	ts.ErrorContains(err, "status code 777")
	ts.ErrorContains(err, "errorCode: d0a17902-63ed-4cb6-a8e8-fac5ca31b0b7")
	ts.ErrorContains(err, "errorMessage: Message parsing failed: Unexpected character (';' (code 34)): was")
}

func (ts *TSError) TestInvalidBodyNotReturnsCodeMessageError() {
	err := newCodeMessageError(777, invalidBody)
	errText := err.Error()
	ts.ErrorContains(err, "status code 777:")
	ts.NotContains(errText, "errorCode:")
	ts.NotContains(errText, "errorMessage:")
}
