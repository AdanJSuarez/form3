package handler

import (
	"fmt"
	"net/http"
)

const serviceUnavailableMessage = "service is temporarily overloaded. This is safe to retry after waiting a short amount of time"

type serviceUnavailableHandler struct {
	next StatusHandler
}

func NewServiceUnavailableHandler() StatusHandler {
	return &serviceUnavailableHandler{}
}

func (s *serviceUnavailableHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusServiceUnavailable {
		return newError(response.StatusCode, fmt.Errorf(serviceUnavailableMessage))
	}
	return s.next.Execute(response)
}

func (s *serviceUnavailableHandler) SetNext(next StatusHandler) {
	s.next = next
}
