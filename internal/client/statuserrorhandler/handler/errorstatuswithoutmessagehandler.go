package handler

import (
	"fmt"
	"net/http"
)

const (
	unauthorizedMessage  = "invalid request signature or access token"
	notAcceptableMessage = `trying to access content with an incorrect
		content type specific in the request header`
	serverErrorMessage = `an internal error occurs or the request
		times out. This is safe to retry after waiting a short amount of time`
	badGatewayMessage = `there is a temporary internal networking problem.
		This is safe to retry after waiting a short amount of time`
	serviceUnavailableMessage = `service is temporarily overloaded.
		This is safe to retry after waiting a short amount of time`
	gatewayTimeoutMessage = `there is a temporary internal networking
		problem. This is safe to retry after waiting a short amount of time`
)

type errorStatusWithoutMessageHandler struct {
	next StatusErrorHandler
}

func NewErrorStatusWithoutMessageHandler() StatusErrorHandler {
	return &errorStatusWithoutMessageHandler{}
}

func (u *errorStatusWithoutMessageHandler) Execute(response *http.Response) error {
	switch {
	case response.StatusCode == http.StatusUnauthorized:
		return newError(response.StatusCode, fmt.Errorf(unauthorizedMessage))
	case response.StatusCode == http.StatusNotAcceptable:
		return newError(response.StatusCode, fmt.Errorf(notAcceptableMessage))
	case response.StatusCode == http.StatusInternalServerError:
		return newError(response.StatusCode, fmt.Errorf(serverErrorMessage))
	case response.StatusCode == http.StatusBadGateway:
		return newError(response.StatusCode, fmt.Errorf(badGatewayMessage))
	case response.StatusCode == http.StatusServiceUnavailable:
		return newError(response.StatusCode, fmt.Errorf(serviceUnavailableMessage))
	case response.StatusCode == http.StatusGatewayTimeout:
		return newError(response.StatusCode, fmt.Errorf(gatewayTimeoutMessage))
	}
	return u.next.Execute(response)
}

func (u *errorStatusWithoutMessageHandler) SetNext(next StatusErrorHandler) {
	u.next = next
}
