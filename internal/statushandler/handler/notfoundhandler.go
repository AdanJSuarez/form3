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

func (n *notFoundHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusNotFound {
		// TODO: Check message returned and the right and set the right message returned. It looks like it returns html message.
		b := []byte{}
		response.Body.Read(b)
		return fmt.Errorf(notFoundFmt, b)
	}
	return n.next.Execute(response)
}

func (n *notFoundHandler) SetNext(next StatusHandler) {
	n.next = next
}
