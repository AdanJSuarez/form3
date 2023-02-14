package statushandler

import (
	"net/http"

	"github.com/AdanJSuarez/form3/internal/statushandler/handler"
)

type StatusErrorHandler struct {
	next handler.StatusHandler
}

func NewStatusError() *StatusErrorHandler {
	sh := &StatusErrorHandler{}
	// hForbidden := handler.newForbiddenHandler()
	hBadRequest := handler.NewBadRequestHandler()
	// hBadRequest.setNext(hForbidden)
	sh.next = hBadRequest
	return sh
}

func (s *StatusErrorHandler) HandleError(response *http.Response) error {
	// TODO: Implement handleError
	return s.next.Execute(response)
}
