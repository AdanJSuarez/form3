package handler

import (
	"fmt"
	"net/http"
)

const unauthorizedMessage = "invalid request signature or access token"

type unauthorizedHandler struct {
	next StatusErrorHandler
}

func NewUnauthorizedHandler() StatusErrorHandler {
	return &unauthorizedHandler{}
}

func (u *unauthorizedHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusUnauthorized {
		return newError(response.StatusCode, fmt.Errorf(unauthorizedMessage))
	}
	return u.next.Execute(response)
}

func (u *unauthorizedHandler) SetNext(next StatusErrorHandler) {
	u.next = next
}
