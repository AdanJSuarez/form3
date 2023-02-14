package handler

import (
	"fmt"
	"net/http"
)

const notAcceptableMessage = "trying to access content with an incorrect content type specific in the request header"

type notAcceptableHandler struct {
	next StatusHandler
}

func NewNotAcceptableHandler() StatusHandler {
	return &notAcceptableHandler{}
}

func (s *notAcceptableHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusNotAcceptable {
		return newError(response.StatusCode, fmt.Errorf(notAcceptableMessage))
	}
	return s.next.Execute(response)
}

func (s *notAcceptableHandler) SetNext(next StatusHandler) {
	s.next = next
}
