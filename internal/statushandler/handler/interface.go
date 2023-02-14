package handler

import "net/http"

//go:generate mockery --inpackage --name=statusHandler

type StatusErrorHandler interface {
	Execute(response *http.Response) error
	SetNext(StatusErrorHandler)
}
