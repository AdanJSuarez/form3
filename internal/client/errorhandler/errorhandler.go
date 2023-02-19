package errorhandler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AdanJSuarez/form3/internal/client/errorhandler/handler"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

const nilResponseError = "http response is nil"

type ErrorHandler struct {
	next handler.StatusErrorHandler
}

// TODO: Implement Retry

func NewErrorHandler() *ErrorHandler {
	sh := &ErrorHandler{}
	uncoveredStatus := handler.NewUncoveredHandler()
	chainOfResponsibilityErrors := sh.chainOfResponsibilityErrors(uncoveredStatus)
	sh.next = chainOfResponsibilityErrors
	return sh
}

// func (s *ErrorHandler) StatusCreated(response *http.Response) bool {
// 	if response == nil {
// 		return false
// 	}
// 	return response.StatusCode == http.StatusCreated
// }

// func (s *ErrorHandler) StatusOK(response *http.Response) bool {
// 	if response == nil {
// 		return false
// 	}
// 	return response.StatusCode == http.StatusOK
// }

// func (s *ErrorHandler) StatusNoContent(response *http.Response) bool {
// 	if response == nil {
// 		return false
// 	}
// 	return response.StatusCode == http.StatusNoContent
// }

func (s *ErrorHandler) StatusError(request *http.Request, response *http.Response) (*http.Response, error) {
	if response == nil {
		return nil, fmt.Errorf(nilResponseError)
	}
	return nil, s.next.Execute(response)
}

func (s *ErrorHandler) Error(request *http.Request, err error) (*http.Response, error) {
	//TODO: Implement HandleError
	if os.IsTimeout(err) {
		// Implement retry mechanism
		return nil, fmt.Errorf("not implemented retry mechanism: %v", err)
	}
	return nil, err
}

func (s *ErrorHandler) chainOfResponsibilityErrors(otherHandler handler.StatusErrorHandler) handler.StatusErrorHandler {
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
