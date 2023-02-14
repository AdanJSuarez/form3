package handler

import (
	"fmt"
	"net/http"
)

const notFoundFmt = "not found: %s"

type notFoundHandler struct {
	next StatusHandler
}

func NewNotFoundHandler() StatusHandler {
	return &notFoundHandler{}
}

func (s *notFoundHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusNotFound {
		// TODO: Check message returned and the right and set the right message returned. It looks like it returns html message.
		b := []byte{}
		response.Body.Read(b)
		return fmt.Errorf(notFoundFmt, b)
	}
	return s.next.Execute(response)
}

func (s *notFoundHandler) SetNext(next StatusHandler) {
	s.next = next
}
