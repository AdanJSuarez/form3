package handler

import (
	"net/http"
)

type forbiddenHandler struct {
	next StatusErrorHandler
}

func NewForbiddenHandler() StatusErrorHandler {
	return &forbiddenHandler{}
}

func (f *forbiddenHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusForbidden {
		return newTypeDescriptionError(response.StatusCode, response.Body)
	}
	return f.next.Execute(response)
}

func (f *forbiddenHandler) SetNext(next StatusErrorHandler) {
	f.next = next
}
