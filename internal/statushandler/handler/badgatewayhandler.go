package handler

import (
	"fmt"
	"net/http"
)

const badGatewayMessage = "here is a temporary internal networking problem. This is safe to retry after waiting a short amount of time"

type badGatewayHandler struct {
	next StatusHandler
}

func NewBadGatewayHandler() StatusHandler {
	return &badGatewayHandler{}
}

func (b *badGatewayHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusBadGateway {
		return newError(response.StatusCode, fmt.Errorf(badGatewayMessage))
	}
	return b.next.Execute(response)
}

func (b *badGatewayHandler) SetNext(next StatusHandler) {
	b.next = next
}
