package handler

import (
	"fmt"
	"net/http"
)

const uncoveredMessage = "uncovered status code"

type uncoveredHandler struct {
}

func NewUncoveredHandler() StatusHandler {
	return &uncoveredHandler{}
}

func (u *uncoveredHandler) Execute(response *http.Response) error {
	return newError(response.StatusCode, fmt.Errorf(uncoveredMessage))
}

// SetNext is supposed to not be called in this struct.
// It is here to implement the StatusHandler interface.
func (u *uncoveredHandler) SetNext(next StatusHandler) {
	// Do nothing.
}
