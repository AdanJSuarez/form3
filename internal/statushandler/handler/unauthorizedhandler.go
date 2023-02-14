package handler

import (
	"fmt"
	"net/http"
)

const unauthorizedMessage = "invalid request signature or access token"

type unauthorizedHandler struct {
	next StatusHandler
}

func NewUnauthorizedHandler() StatusHandler {
	return &unauthorizedHandler{}
}

func (u *unauthorizedHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusUnauthorized {
		return newError(response.StatusCode, fmt.Errorf(unauthorizedMessage))
	}
	return u.next.Execute(response)
}

func (u *unauthorizedHandler) SetNext(next StatusHandler) {
	u.next = next
}
