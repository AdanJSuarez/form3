package handler

import "net/http"

type StatusErrorHandler interface {
	Execute(response *http.Response) error
	SetNext(StatusErrorHandler)
}
