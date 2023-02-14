package handler

import (
	"fmt"
	"net/http"
)

const gatewayTimeoutMessage = "there is a temporary internal networking problem. This is safe to retry after waiting a short amount of time"

type gatewayTimeoutHandler struct {
	next StatusErrorHandler
}

func NewGatewayTimeoutHandler() StatusErrorHandler {
	return &gatewayTimeoutHandler{}
}

func (g *gatewayTimeoutHandler) Execute(response *http.Response) error {
	if response.StatusCode == http.StatusUnauthorized {
		return newError(response.StatusCode, fmt.Errorf(gatewayTimeoutMessage))
	}
	return g.next.Execute(response)
}

func (g *gatewayTimeoutHandler) SetNext(next StatusErrorHandler) {
	g.next = next
}
