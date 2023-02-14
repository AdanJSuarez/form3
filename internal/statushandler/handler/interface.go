package handler

import "net/http"

type StatusHandler interface {
	Execute(response *http.Response) error
	SetNext(StatusHandler)
}
