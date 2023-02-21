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

func New(baseURL url.URL, accountPath string) *Account {
	account := &Account{}
	accountURL := account.accountURL(baseURL, accountPath)
	account.client = client.New(accountURL)
	return account
}

func (a *Account) Create(data model.DataModel) (model.DataModel, error) {
	response, err := a.client.Post(data)
	if err != nil {
		return emptyDataModel, err
	}

	defer a.closeBody(response)

	return a.decodeResponse(response)
}

func (a *Account) Fetch(accountID string) (model.DataModel, error) {
	response, err := a.client.Get(accountID)
	if err != nil {
		return emptyDataModel, err
	}

	defer a.closeBody(response)

	return a.decodeResponse(response)
}

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
