package statushandler

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/statushandler/handler"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

type StatusHandler struct {
	next handler.StatusErrorHandler
}

func NewStatusHandler() *StatusHandler {
	sh := &StatusHandler{}
	uncoveredStatus := handler.NewUncoveredHandler()
	errs5XX := sh.errors5XXChained(uncoveredStatus)
	errs4XX := sh.errors4XXChained(errs5XX)
	sh.next = errs4XX
	return sh
}

func (s *StatusHandler) StatusCreated(response *http.Response) bool {
	return response.StatusCode == http.StatusCreated
}

func (s *StatusHandler) StatusOK(response *http.Response) bool {
	return response.StatusCode == http.StatusOK
}

func (s *StatusHandler) StatusNotContent(response *http.Response) bool {
	return response.StatusCode == http.StatusNoContent
}

func (s *StatusHandler) HandleError(response *http.Response) error {
	return s.next.Execute(response)
}

func (s *StatusHandler) errors4XXChained(otherHandler handler.StatusErrorHandler) handler.StatusErrorHandler {
	tooManyRequests := handler.NewTooManyRequestsHandler()
	conflict := handler.NewConflictHandler()
	notAcceptable := handler.NewNotAcceptableHandler()
	methodNotAllowed := handler.NewMethodNotAllowedHandler()
	notFound := handler.NewNotFoundHandler()
	forbidden := handler.NewForbiddenHandler()
	unauthorized := handler.NewUnauthorizedHandler()
	badRequest := handler.NewBadRequestHandler()

	tooManyRequests.SetNext(otherHandler)
	conflict.SetNext(tooManyRequests)
	notAcceptable.SetNext(conflict)
	methodNotAllowed.SetNext(notAcceptable)
	notFound.SetNext(methodNotAllowed)
	forbidden.SetNext(notFound)
	unauthorized.SetNext(forbidden)
	badRequest.SetNext(unauthorized)

	return badRequest
}

func (s *StatusHandler) errors5XXChained(otherHandler handler.StatusErrorHandler) handler.StatusErrorHandler {
	gatewayTimeout := handler.NewGatewayTimeoutHandler()
	serviceUnavailable := handler.NewServiceUnavailableHandler()
	badGateway := handler.NewBadGatewayHandler()
	serverError := handler.NewServerErrorHandler()

	gatewayTimeout.SetNext(otherHandler)
	serviceUnavailable.SetNext(gatewayTimeout)
	badGateway.SetNext(serviceUnavailable)
	serverError.SetNext(badGateway)

	return serverError
}
