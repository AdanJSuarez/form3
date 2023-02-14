package handler

import (
	"fmt"
	"net/http"
)

const methodNotAllowedMessage = "trying to access an endpoint that exists using a method that is not supported by the target resource"

type methodNotAllowedHandler struct {
	next StatusHandler
}

func NewMethodNotAllowedHandler() StatusHandler {
	return &methodNotAllowedHandler{}
}

func (s *methodNotAllowedHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusMethodNotAllowed {
		return newError(response.StatusCode, fmt.Errorf(methodNotAllowedMessage))
	}
	return s.next.Execute(response)
}

func (s *methodNotAllowedHandler) SetNext(next StatusHandler) {
	s.next = next
}
