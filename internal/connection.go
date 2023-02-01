package internal

import (
	"net/http"
	"time"
)

const timeout = 5 * time.Second

type URL string

type Connection struct {
	url    string
	client http.Client
}

func NewConnection(URL string) *Connection {
	client := http.Client{
		Timeout: timeout,
	}

	return &Connection{url: URL, client: client}
}
