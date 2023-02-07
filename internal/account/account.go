package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AdanJSuarez/form3/internal/client"
	"github.com/AdanJSuarez/form3/pkg/model"
)

var emptyData = model.DataModel{}

type Account struct {
	client client.Client
}

// New returns a pointer of Account initialized
func New(url string) *Account {
	return &Account{
		client: *client.New(url),
	}
}

// Create creates an bank account and returns the account values.
// It returns an error otherwise.
func (a *Account) Create(data model.DataModel) (model.DataModel, error) {
	requestBody := a.requestBody(data)
	response, err := a.client.Post(requestBody)
	if err != nil {
		return emptyData, err
	}
	defer response.Body.Close()

	if !a.statusCreated(response) {
		return emptyData, fmt.Errorf("status code: %d", response.StatusCode)
	}

	return a.decodeResponse(response)
}

// Fetch retrieves the account information for the specific account ID.
// It returns an error otherwise.
func (a *Account) Fetch(accountID string) (model.DataModel, error) {
	response, err := a.client.Get(accountID)
	if err != nil {
		return model.DataModel{}, err
	}
	defer response.Body.Close()

	if !a.statusSuccess(response) {
		return emptyData, fmt.Errorf("status code: %d", response.StatusCode)
	}

	return a.decodeResponse(response)
}

// Delete deletes an account by its ID and version number.
// It returns an error otherwise.
func (a *Account) Delete(accountID string, version int) error {
	accountIDVersion := fmt.Sprintf("%s?version=%d", accountID, version)
	response, err := a.client.Delete(accountIDVersion)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if !a.statusDeleted(response) {
		return fmt.Errorf("status code: %d", response.StatusCode)
	}
	// TODO: FINISH DELETE
	return err
}

func (a *Account) requestBody(data model.DataModel) io.Reader {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return bytes.NewBuffer([]byte{})
	}
	return bytes.NewBuffer(dataBytes)
}

func (a *Account) decodeResponse(response *http.Response) (model.DataModel, error) {
	dataReturned := &model.DataModel{}
	if err := json.NewDecoder(response.Body).Decode(dataReturned); err != nil {
		return emptyData, err
	}
	return *dataReturned, nil
}

func (a *Account) statusCreated(response *http.Response) bool {
	return response.StatusCode == http.StatusCreated
}

func (a *Account) statusSuccess(response *http.Response) bool {
	return response.StatusCode == http.StatusOK
}

func (a *Account) statusDeleted(respose *http.Response) bool {
	if respose.StatusCode == http.StatusNoContent {
		// TODO: Implement this
	}
	return false
}
