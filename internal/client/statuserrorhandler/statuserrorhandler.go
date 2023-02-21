package statuserrorhandler

import (
	"fmt"
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client/statuserrorhandler/handler"
)

// Ref: https://refactoring.guru/design-patterns/chain-of-responsibility

const nilResponseError = "http response is nil"

type StatusErrorHandler struct {
	next handler.StatusErrorHandler
}

func NewStatusErrorHandler() *StatusErrorHandler {
	sh := &StatusErrorHandler{}
	uncoveredStatus := handler.NewUncoveredHandler()
	chainOfResponsibilityErrors := sh.chainOfResponsibilityErrors(uncoveredStatus)
	sh.next = chainOfResponsibilityErrors
	return sh
}

func (s *StatusErrorHandler) StatusError(response *http.Response) (*http.Response, error) {
	if response == nil {
		return nil, fmt.Errorf(nilResponseError)
	}
	return nil, s.next.Execute(response)
}

func (s *StatusErrorHandler) chainOfResponsibilityErrors(
	otherHandler handler.StatusErrorHandler) handler.StatusErrorHandler {
	tooManyRequests := handler.NewTooManyRequestsHandler()
	conflict := handler.NewConflictHandler()
	methodNotAllowed := handler.NewMethodNotAllowedHandler()
	notFound := handler.NewNotFoundHandler()
	forbidden := handler.NewForbiddenHandler()
	unauthorized := handler.NewErrorStatusWithoutMessageHandler()
	badRequest := handler.NewBadRequestHandler()

	tooManyRequests.SetNext(otherHandler)
	conflict.SetNext(tooManyRequests)
	methodNotAllowed.SetNext(conflict)
	notFound.SetNext(methodNotAllowed)
	forbidden.SetNext(notFound)
	unauthorized.SetNext(forbidden)
	badRequest.SetNext(unauthorized)

	return badRequest
}
