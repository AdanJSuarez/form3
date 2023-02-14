package handler

import (
	"fmt"
	"net/http"
)

const conflictHandlerMessage = `resource has already been created. It is safe ignore this error message and continue processing.
Returned for DELETE calls when an incorrect version has been specified`

type conflictHandler struct {
	next StatusHandler
}

func NewConflictHandler() StatusHandler {
	return &conflictHandler{}
}

func (c *conflictHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusConflict {
		return newError(response.StatusCode, fmt.Errorf(conflictHandlerMessage))
	}
	return c.next.Execute(response)
}

func (s *conflictHandler) SetNext(next StatusHandler) {
	s.next = next
}
