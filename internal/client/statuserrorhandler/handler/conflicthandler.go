package handler

import (
	"net/http"
)

type conflictHandler struct {
	next StatusErrorHandler
}

func NewConflictHandler() StatusErrorHandler {
	return &conflictHandler{}
}

func (c *conflictHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusConflict {
		return newCodeMessageError(response.StatusCode, response.Body)
	}
	return c.next.Execute(response)
}

func (c *conflictHandler) SetNext(next StatusErrorHandler) {
	c.next = next
}
