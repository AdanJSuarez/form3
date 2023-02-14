package handler

import (
	"fmt"
	"net/http"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

const unauthorizedMessage = "invalid request signature or access token"

type unauthorizedHandler struct {
	next StatusHandler
}

func NewUnauthorizedHandler() StatusHandler {
	return &unauthorizedHandler{}
}

func (s *unauthorizedHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusUnauthorized {
		return newError(response.StatusCode, fmt.Errorf(unauthorizedMessage))
	}
	return s.next.Execute(response)
}

func (s *unauthorizedHandler) SetNext(next StatusHandler) {
	s.next = next
}
