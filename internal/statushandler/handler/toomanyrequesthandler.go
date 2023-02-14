package handler

import (
	"fmt"
	"net/http"
)

const tooManyRequestsHandlerMessage = `the rate limit for requests per second has been exceeded.
Also used in the Form3 Multi-Cloud stack when an attempted change involves a resource whose state is still being synchronised across the stack.
Wait, then retry later`

type tooManyRequestsHandler struct {
	next StatusHandler
}

func NewTooManyRequestsHandler() StatusHandler {
	return &tooManyRequestsHandler{}
}

func (c *tooManyRequestsHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusTooManyRequests {
		return newError(response.StatusCode, fmt.Errorf(tooManyRequestsHandlerMessage))
	}
	return c.next.Execute(response)
}

func (s *tooManyRequestsHandler) SetNext(next StatusHandler) {
	s.next = next
}
