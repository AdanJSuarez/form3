package handler

import (
	"fmt"
	"net/http"
)

const notFoundMessage = "not found: trying to access a non-existent endpoint or resource. Returned in some APIs when a queried parameter cannot be found"

type notFoundHandler struct {
	next StatusErrorHandler
}

func NewNotFoundHandler() StatusErrorHandler {
	return &notFoundHandler{}
}

func (n *notFoundHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusNotFound {
		return newError(response.StatusCode, fmt.Errorf(notFoundMessage))
	}
	return n.next.Execute(response)
}

func (n *notFoundHandler) SetNext(next StatusErrorHandler) {
	n.next = next
}
