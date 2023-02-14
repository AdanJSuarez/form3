package handler

import (
	"net/http"
)

type forbiddenHandler struct {
	next StatusHandler
}

func NewForbiddenHandler() StatusHandler {
	return &forbiddenHandler{}
}

func (s *forbiddenHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusForbidden {
		return newTypeDescriptionError(response.StatusCode, response.Body)
	}
	return s.next.Execute(response)
}

func (s *forbiddenHandler) SetNext(next StatusHandler) {
	s.next = next
}
