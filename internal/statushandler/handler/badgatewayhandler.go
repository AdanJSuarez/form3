package handler

import (
	"fmt"
	"net/http"
)

const badGatewayMessage = "there is a temporary internal networking problem. This is safe to retry after waiting a short amount of time"

type badGatewayHandler struct {
	next StatusErrorHandler
}

func NewBadGatewayHandler() StatusErrorHandler {
	return &badGatewayHandler{}
}

func (b *badGatewayHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusBadGateway {
		return newError(response.StatusCode, fmt.Errorf(badGatewayMessage))
	}
	return b.next.Execute(response)
}

func (b *badGatewayHandler) SetNext(next StatusErrorHandler) {
	b.next = next
}
