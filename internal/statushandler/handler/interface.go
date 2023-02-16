package handler

import "net/http"

//go:generate mockery --inpackage --name=StatusErrorHandler

type StatusErrorHandler interface {
	Execute(response *http.Response) error
	SetNext(StatusErrorHandler)
}
