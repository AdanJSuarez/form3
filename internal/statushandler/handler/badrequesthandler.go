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

func (b *badRequestHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusBadRequest {
		return newCodeMessageError(response.StatusCode, response.Body)
	}
	return b.next.Execute(response)
}

func (b *badRequestHandler) SetNext(next StatusHandler) {
	b.next = next
}
