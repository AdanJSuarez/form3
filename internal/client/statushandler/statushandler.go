package statushandler

import (
	"fmt"
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client/statushandler/handler"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

const nilResponseError = "http response is nil"

type StatusHandler struct {
	next handler.StatusErrorHandler
}

// TODO: Implement Retry

func NewStatusHandler() *StatusHandler {
	sh := &StatusHandler{}
	uncoveredStatus := handler.NewUncoveredHandler()
	chainOfResponsibilityErrors := sh.chainOfResponsibilityErrors(uncoveredStatus)
	sh.next = chainOfResponsibilityErrors
	return sh
}

func (s *StatusHandler) StatusCreated(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusCreated
}

func (s *StatusHandler) StatusOK(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusOK
}

func (s *StatusHandler) StatusNoContent(response *http.Response) bool {
	if response == nil {
		return false
	}
	return response.StatusCode == http.StatusNoContent
}

func (s *StatusHandler) HandleError(response *http.Response) error {
	if response == nil {
		return fmt.Errorf(nilResponseError)
	}
	return s.next.Execute(response)
}

func (s *StatusHandler) chainOfResponsibilityErrors(otherHandler handler.StatusErrorHandler) handler.StatusErrorHandler {
	gatewayTimeout := handler.NewGatewayTimeoutHandler()
	serviceUnavailable := handler.NewServiceUnavailableHandler()
	badGateway := handler.NewBadGatewayHandler()
	serverError := handler.NewServerErrorHandler()

	tooManyRequests := handler.NewTooManyRequestsHandler()
	conflict := handler.NewConflictHandler()
	notAcceptable := handler.NewNotAcceptableHandler()
	methodNotAllowed := handler.NewMethodNotAllowedHandler()
	notFound := handler.NewNotFoundHandler()
	forbidden := handler.NewForbiddenHandler()
	unauthorized := handler.NewUnauthorizedHandler()
	badRequest := handler.NewBadRequestHandler()

	gatewayTimeout.SetNext(otherHandler)
	serviceUnavailable.SetNext(gatewayTimeout)
	badGateway.SetNext(serviceUnavailable)
	serverError.SetNext(badGateway)

	tooManyRequests.SetNext(serverError)
	conflict.SetNext(tooManyRequests)
	notAcceptable.SetNext(conflict)
	methodNotAllowed.SetNext(notAcceptable)
	notFound.SetNext(methodNotAllowed)
	forbidden.SetNext(notFound)
	unauthorized.SetNext(forbidden)
	badRequest.SetNext(unauthorized)

	return badRequest
}
