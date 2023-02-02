package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AdanJSuarez/form3/model"
)

var emptyData = model.AccountData{}

func Create(url URL, accountData model.AccountData) (model.AccountData, error) {
	connection := NewConnection(POST, url)
	requestBody := setBody(accountData)

	response, err := connection.Post(requestBody)
	if err != nil {
		return emptyData, err
	}
	defer response.Body.Close()

	dataReturned, err := decodeResponse(response)
	if err != nil {
		return emptyData, err
	}

	return dataReturned, nil
}

func setBody(accData model.AccountData) io.Reader {
	data, err := json.Marshal(accData)
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}

	return bytes.NewBuffer(data)
}

func decodeResponse(res *http.Response) (model.AccountData, error) {
	dataReturned := model.AccountData{}
	if err := json.NewDecoder(res.Body).Decode(&dataReturned); err != nil {
		return emptyData, err
	}

	if res.StatusCode != http.StatusCreated {
		return emptyData, fmt.Errorf("status code: %d", res.StatusCode)
	}

	return dataReturned, nil
}
