package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	errorFmt                = "status code %d: %v"
	errorCodeMessageFmt     = "errorCode: %s - errorMessage: %s"
	errorTypeDescriptionFmt = "error: %s - errorDescription: %s"
)

type errorCodeMessage struct {
	Message string `json:"error_message"`
	Code    string `json:"error_code"`
}

type errorTypeDescription struct {
	ErrorType        string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func newError(statusCode int, err error) error {
	return fmt.Errorf(errorFmt, statusCode, err)
}

func newTypeDescriptionError(statusCode int, body io.ReadCloser) error {
	dataReturned := errorTypeDescription{}
	if err := json.NewDecoder(body).Decode(&dataReturned); err != nil {
		return fmt.Errorf(errorFmt, statusCode, err)
	}
	messageCode := fmt.Sprintf(errorTypeDescriptionFmt, dataReturned.ErrorType, dataReturned.ErrorDescription)
	return fmt.Errorf(errorFmt, statusCode, messageCode)
}

func newCodeMessageError(statusCode int, body io.ReadCloser) error {
	dataReturned := errorCodeMessage{}
	if err := json.NewDecoder(body).Decode(&dataReturned); err != nil {
		return fmt.Errorf(errorFmt, statusCode, err)
	}
	dataReturned.Message = strings.ReplaceAll(dataReturned.Message, "validation failure list:\n", "")
	messageCode := fmt.Sprintf(errorCodeMessageFmt, dataReturned.Code, dataReturned.Message)
	return fmt.Errorf(errorFmt, statusCode, messageCode)
}
