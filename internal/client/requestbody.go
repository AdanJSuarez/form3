package client

import "io"

type RequestBody struct {
	body io.Reader
	size int
}

func NewRequestBody(body io.Reader, size int) RequestBody {
	return RequestBody{
		body: body,
		size: size,
	}
}

func (b RequestBody) Body() io.Reader {
	return b.body
}

func (b RequestBody) Size() int {
	return b.size
}
