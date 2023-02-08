package client

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

const desireFmt = "sha-256=%s"

type RequestBody struct {
	data   []byte
	body   io.ReadCloser
	size   int
	digest string
}

func NewRequestBody(data interface{}) RequestBody {
	rb := RequestBody{}
	if data == nil {
		return rb
	}
	rb.data = rb.dataToBytes(data)
	rb.body = rb.dataToBody()
	rb.size = len(rb.data)
	rb.digest = rb.digestFormatted()
	return rb
}

func (b *RequestBody) Body() io.ReadCloser {
	return b.body
}

func (b *RequestBody) Size() int {
	return b.size
}

func (b *RequestBody) Digest() string {
	return b.digest
}

func (b *RequestBody) dataToBytes(data interface{}) []byte {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return dataBytes
}

func (b *RequestBody) dataToBody() io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer(b.data))
}

func (b *RequestBody) digestFormatted() string {
	hash := sha256.New()
	hash.Write(b.data)
	hashBytes := hash.Sum(nil)
	desire := fmt.Sprintf(desireFmt, base64.StdEncoding.EncodeToString(hashBytes))
	return desire
}
