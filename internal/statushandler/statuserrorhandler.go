package statushandler

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/statushandler/handler"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

type StatusErrorHandler struct {
	next handler.StatusHandler
}

func NewStatusError() *StatusErrorHandler {
	// gatewayTimeout := handler.NewGatewayTimeoutHandler()
	// serviceUnavailable := handler.NewServiceUnavailableHandler()
	// badGateway := handler.NewBadGatewayHandler()
	// serverError := handler.NewServerErrorHandler()
	// tooManyRequests := handler.NewTooManyRequestsHandler()
	// conflict := handler.NewConflictHandler()
	// notAcceptable := handler.NewNotAcceptableHandler()
	// methodNotAllowed := handler.NewMethodNotAllowedHandler()
	// notFound := handler.NewNotFoundHandler()
	// forbidden := handler.NewForbiddenHandler()
	// unauthorized := handler.NewUnauthorizedHandler()
	// badRequest := handler.NewBadRequestHandler()

	// serviceUnavailable.SetNext(gatewayTimeout)
	// badGateway.SetNext(serviceUnavailable)
	// serverError.SetNext(badGateway)
	// tooManyRequests.SetNext(serverError)
	// conflict.SetNext(tooManyRequests)
	// notAcceptable.SetNext(conflict)
	// methodNotAllowed.SetNext(notAcceptable)
	// notFound.SetNext(methodNotAllowed)
	// forbidden.SetNext(notFound)
	// unauthorized.SetNext(forbidden)
	// badRequest.SetNext(unauthorized)

	sh := &StatusErrorHandler{}
	sh.next = sh.error4XX()
	return sh
}

func (s *StatusErrorHandler) HandleError(response *http.Response) error {
	return s.next.Execute(response)
}

func (s *StatusErrorHandler) error4XX() handler.StatusHandler {
	tooManyRequests := handler.NewTooManyRequestsHandler()
	conflict := handler.NewConflictHandler()
	notAcceptable := handler.NewNotAcceptableHandler()
	methodNotAllowed := handler.NewMethodNotAllowedHandler()
	notFound := handler.NewNotFoundHandler()
	forbidden := handler.NewForbiddenHandler()
	unauthorized := handler.NewUnauthorizedHandler()
	badRequest := handler.NewBadRequestHandler()

	error5XX := s.error5XX()

	tooManyRequests.SetNext(error5XX)
	conflict.SetNext(tooManyRequests)
	notAcceptable.SetNext(conflict)
	methodNotAllowed.SetNext(notAcceptable)
	notFound.SetNext(methodNotAllowed)
	forbidden.SetNext(notFound)
	unauthorized.SetNext(forbidden)
	badRequest.SetNext(unauthorized)

	return badRequest
}

func (s *StatusErrorHandler) error5XX() handler.StatusHandler {
	gatewayTimeout := handler.NewGatewayTimeoutHandler()
	serviceUnavailable := handler.NewServiceUnavailableHandler()
	badGateway := handler.NewBadGatewayHandler()
	serverError := handler.NewServerErrorHandler()

	serviceUnavailable.SetNext(gatewayTimeout)
	badGateway.SetNext(serviceUnavailable)
	serverError.SetNext(badGateway)

	return serverError
}
