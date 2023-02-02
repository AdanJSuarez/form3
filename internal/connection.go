package internal

import (
	"net/http"
	"time"
)

const timeout = 5 * time.Second
const POST = "POST"

type URL string

type Connection struct {
	url    string
	client http.Client
}

func NewConnection(URL string) *Connection {
	client := http.Client{
		Timeout: timeout,
	}
	r, err := http.NewRequest(POST, "", nil)
	if err != nil {
		//
	}
	r.Header.Add("", "")
	// res, err := client.Do(r)

	return &Connection{url: URL, client: client}
}
