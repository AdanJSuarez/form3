package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AdanJSuarez/form3/internal/connection"
	"github.com/AdanJSuarez/form3/model"
)

var emptyData = model.Data{}

type Account struct {
	url        string
	connection connection.Connection
}

func New(url string) *Account {
	return &Account{
		url:        url,
		connection: *connection.New(url),
	}
}

func (a *Account) Create(data model.Data) (model.Data, error) {
	requestBody := a.requestBody(data)
	response, err := a.connection.Post(requestBody)
	if err != nil {
		return emptyData, err
	}
	defer response.Body.Close()

	dataReturned, err := a.decodeResponse(response)
	if err != nil {
		return emptyData, err
	}

	return dataReturned, nil
}

func (a *Account) Fetch(accountID string) (model.Data, error) {
	response, err := a.connection.Get(accountID)
	if err != nil {
		return emptyData, err
	}
	return a.decodeResponse(response)
}

func (a *Account) requestBody(data model.Data) io.Reader {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}

	return bytes.NewBuffer(dataBytes)
}

func (a *Account) decodeResponse(response *http.Response) (model.Data, error) {
	dataReturned := &model.Data{}
	if err := json.NewDecoder(response.Body).Decode(dataReturned); err != nil {
		return emptyData, err
	}

	if response.StatusCode != http.StatusCreated {
		return emptyData, fmt.Errorf("status code: %d", response.StatusCode)
	}

	return *dataReturned, nil
}
