package handler

import (
	"net/http"
)

type badRequestHandler struct {
	next StatusHandler
}

func NewBadRequestHandler() StatusHandler {
	return &badRequestHandler{}
}

func (s *badRequestHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusBadRequest {
		return newCodeMessageError(response.StatusCode, response.Body)
	}
	return s.next.Execute(response)
}

func (s *badRequestHandler) SetNext(next StatusHandler) {
	s.next = next
}
