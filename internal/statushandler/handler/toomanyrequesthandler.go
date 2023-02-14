package handler

import (
	"fmt"
	"net/http"
)

const tooManyRequestsHandlerMessage = `the rate limit for requests per second has been exceeded.
Also used in the Form3 Multi-Cloud stack when an attempted change involves a resource whose state is still being synchronised across the stack.
Wait, then retry later`

type tooManyRequestsHandler struct {
	next StatusErrorHandler
}

func NewTooManyRequestsHandler() StatusErrorHandler {
	return &tooManyRequestsHandler{}
}

func (t *tooManyRequestsHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusTooManyRequests {
		return newError(response.StatusCode, fmt.Errorf(tooManyRequestsHandlerMessage))
	}
	return t.next.Execute(response)
}

func (t *tooManyRequestsHandler) SetNext(next StatusErrorHandler) {
	t.next = next
}
