package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/client"
	"github.com/AdanJSuarez/form3/pkg/model"
)

const (
	httpResponseNilError = "http response is nil"
	versionParam         = "version"
)

var emptyDataModel = model.DataModel{}

type Account struct {
	client Client
}

// New returns a pointer of "Account" initialized.
func New(config Configuration) *Account {
	baseURL := *config.BaseURL()
	accountPath := config.AccountPath()

	account := &Account{}
	accountURL := account.accountURL(baseURL, accountPath)
	account.client = client.New(accountURL)
	return account
}

/*
Create creates an bank account and returns the account values (model.DataModel).
It returns an error otherwise.

For more reference about model.DataModel values, please check form3 API documentation.
*/
func (a *Account) Create(data model.DataModel) (model.DataModel, error) {
	response, err := a.client.Post(data)
	if err != nil {
		return emptyDataModel, err
	}

	defer a.closeBody(response)

	return a.decodeResponse(response)
}

/*
Fetch retrieves the account information (model.DataModel) for the specific account ID.
It returns an error otherwise.

For more reference about model.DataModel values and accountID, please check form3 API documentation.
*/
func (a *Account) Fetch(accountID string) (model.DataModel, error) {
	response, err := a.client.Get(accountID)
	if err != nil {
		return emptyDataModel, err
	}

	defer a.closeBody(response)

	return a.decodeResponse(response)
}

/*
Delete deletes an account by its ID and version number.
It returns an error otherwise.

For more reference about accountID and version, please check form3 API documentation.
*/
func (a *Account) Delete(accountID string, version int) error {
	response, err := a.client.Delete(accountID, versionParam, fmt.Sprint(version))
	if err != nil {
		return err
	}

	defer a.closeBody(response)

	return nil
}

func (a *Account) accountURL(baseURL url.URL, accountPath string) url.URL {
	baseURL.Path = accountPath
	return baseURL
}

func (a *Account) decodeResponse(response *http.Response) (model.DataModel, error) {
	dataModel := model.DataModel{}
	if response == nil {
		return dataModel, fmt.Errorf(httpResponseNilError)
	}

	if err := json.NewDecoder(response.Body).Decode(&dataModel); err != nil {
		return dataModel, err
	}

	return dataModel, nil
}

func (a *Account) closeBody(response *http.Response) {
	if response != nil && response.Body != nil {
		response.Body.Close()
	}
}
