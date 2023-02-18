package request

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	HOST_KEY              = "Host"
	DATE_KEY              = "Date"
	ACCEPT_KEY            = "Accept"
	ACCEPT_ENCODING_KEY   = "Accept-Encoding"
	CONTENT_TYPE_KEY      = "Content-Type"
	CONTENT_LENGTH_KEY    = "Content-Length"
	DIGEST_KEY            = "Digest"
	CONTENT_TYPE_VALUE    = "application/vnd.api+json"
	ACCEPT_ENCODING_VALUE = "gzip"
	desireFmt             = "sha-256=%s"
)

type RequestHandler struct {
	data []byte
	body io.ReadCloser
}

func NewRequestHandler(data interface{}) *RequestHandler {
	r := &RequestHandler{}
	if data == nil {
		return r
	}
	r.data = r.dataToBytes(data)
	r.body = r.dataToBody()
	return r
}

func (r *RequestHandler) Body() io.ReadCloser {
	return r.body
}

func (r *RequestHandler) Request(method, url, host string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, r.body)
	if err != nil {
		return nil, err
	}

	r.addRequiredHeader(host, request)

	if r.body != nil {
		r.addHeaderToRequestWithBody(request)
	}

	return request, nil
}

func (r *RequestHandler) SetQuery(request *http.Request, parameterKey, parameterValue string) {
	query := request.URL.Query()
	query.Add(parameterKey, parameterValue)
	request.URL.RawQuery = query.Encode()
}

func (r *RequestHandler) dataToBytes(data interface{}) []byte {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return dataBytes
}

func (r *RequestHandler) dataToBody() io.ReadCloser {
	return io.NopCloser(bytes.NewBuffer(r.data))
}

func (r *RequestHandler) addRequiredHeader(host string, request *http.Request) {
	request.Header.Add(HOST_KEY, host)
	request.Header.Add(DATE_KEY, time.Now().Format(time.RFC1123))
	request.Header.Add(ACCEPT_KEY, CONTENT_TYPE_VALUE)
	request.Header.Add(ACCEPT_ENCODING_KEY, ACCEPT_ENCODING_VALUE)
}

func (r *RequestHandler) addHeaderToRequestWithBody(request *http.Request) {
	request.Header.Add(CONTENT_TYPE_KEY, CONTENT_TYPE_VALUE)
	request.Header.Add(CONTENT_LENGTH_KEY, fmt.Sprint(len(r.data)))
	request.Header.Add(DIGEST_KEY, r.digestFormatted())
}

func (r *RequestHandler) digestFormatted() string {
	hash := sha256.New()
	hash.Write(r.data)
	hashBytes := hash.Sum(nil)
	desire := fmt.Sprintf(desireFmt, base64.StdEncoding.EncodeToString(hashBytes))
	return desire
}
