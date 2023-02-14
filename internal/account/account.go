package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/client"
	"github.com/AdanJSuarez/form3/pkg/model"
)

var emptyDataModel = model.DataModel{}

type errorBadRequest struct {
	Message string `json:"error_message"`
	Code    string `json:"error_code"`
}

type Account struct {
	client Client
}

// New returns a pointer of Account initialized
func New(baseURL url.URL, accountPath string) *Account {
	baseURL.Path = accountPath
	return &Account{
		client: client.New(baseURL),
	}
}

// Create creates an bank account and returns the account values.
// It returns an error otherwise.
func (a *Account) Create(data model.DataModel) (model.DataModel, error) {
	dataBody := client.NewRequestBody(data)
	response, err := a.client.Post(dataBody)
	if err != nil {
		return emptyDataModel, err
	}
	defer a.closeBody(response)

	if !a.statusCreated(response) {
		return emptyDataModel, fmt.Errorf("status code: %d", response.StatusCode)
	}

	return a.decodeResponse(response)
}

// Fetch retrieves the account information for the specific account ID.
// It returns an error otherwise.
func (a *Account) Fetch(accountID string) (model.DataModel, error) {
	response, err := a.client.Get(accountID)
	if err != nil {
		return emptyDataModel, err
	}
	defer a.closeBody(response)

	if !a.statusOK(response) {
		return emptyDataModel, fmt.Errorf("status code: %d", response.StatusCode)
	}

	return a.decodeResponse(response)
}

// Delete deletes an account by its ID and version number.
// It returns an error otherwise.
func (a *Account) Delete(accountID string, version int) error {
	response, err := a.client.Delete(accountID, "version", fmt.Sprint(version))
	if err != nil {
		return err
	}
	defer a.closeBody(response)

	if !a.statusDeleted(response) {
		return fmt.Errorf("status code: %d", response.StatusCode)
	}

	return err
}

func (a *Account) decodeResponse(response *http.Response) (model.DataModel, error) {
	dataReturned := &model.DataModel{}
	if err := json.NewDecoder(response.Body).Decode(dataReturned); err != nil {
		return emptyDataModel, err
	}
	return *dataReturned, nil
}

func (a *Account) statusCreated(response *http.Response) bool {
	return response.StatusCode == http.StatusCreated
}

func (a *Account) statusOK(response *http.Response) bool {
	return response.StatusCode == http.StatusOK
}

func (a *Account) statusDeleted(response *http.Response) bool {
	return response.StatusCode == http.StatusNoContent
}

func (a *Account) closeBody(response *http.Response) {
	if response.Body != nil {
		response.Body.Close()
	}
}

// // TODO: Handle different Status code

// func (a *Account) handleBadRequest(response *http.Response) error {
// 	badRequest := "status code 400: %v"
// 	if response.StatusCode == http.StatusBadRequest {
// 		dataReturned := errorBadRequest{}
// 		if err := json.NewDecoder(response.Body).Decode(&dataReturned); err != nil {
// 			return fmt.Errorf(badRequest, err)
// 		}
// 		messageCode := fmt.Sprintf("%s:%s", dataReturned.Message, dataReturned.Code)
// 		return fmt.Errorf(badRequest, messageCode)
// 	}
// 	return nil
// }
