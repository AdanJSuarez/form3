package handler

import (
	"fmt"
	"net/http"
)

const serverErrorMessage = "an internal error occurs or the request times out. This is safe to retry after waiting a short amount of time"

type serverErrorHandler struct {
	next StatusErrorHandler
}

func NewServerErrorHandler() StatusErrorHandler {
	return &serverErrorHandler{}
}

func (s *serverErrorHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusInternalServerError {
		return newError(response.StatusCode, fmt.Errorf(serverErrorMessage))
	}
	return s.next.Execute(response)
}

func (s *serverErrorHandler) SetNext(next StatusErrorHandler) {
	s.next = next
}
