package handler

import (
	"fmt"
	"net/http"
)

const methodNotAllowedMessage = `trying to access an endpoint that exists using
	a method that is not supported by the target resource`

type methodNotAllowedHandler struct {
	next StatusErrorHandler
}

func NewMethodNotAllowedHandler() StatusErrorHandler {
	return &methodNotAllowedHandler{}
}

func (m *methodNotAllowedHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusMethodNotAllowed {
		return newError(response.StatusCode, fmt.Errorf(methodNotAllowedMessage))
	}
	return m.next.Execute(response)
}

func (m *methodNotAllowedHandler) SetNext(next StatusErrorHandler) {
	m.next = next
}
