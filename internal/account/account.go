package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/client"
	"github.com/AdanJSuarez/form3/internal/statushandler"
	"github.com/AdanJSuarez/form3/pkg/model"
)

const httpResponseNilError = "http response is nil"

var emptyDataModel = model.DataModel{}

type Account struct {
	client        Client
	statusHandler StatusHandler
}

// New returns a pointer of Account initialized
func New(baseURL url.URL, accountPath string) *Account {
	baseURL.Path = accountPath
	return &Account{
		client:        client.New(baseURL),
		statusHandler: statushandler.NewStatusHandler(),
	}
}

func (a *Account) Create(data model.DataModel) (model.DataModel, error) {
	dataBody := client.NewRequestBody(data)
	response, err := a.client.Post(dataBody)
	if err != nil {
		return emptyDataModel, err
	}
	defer a.closeBody(response)

	if !a.isCreated(response) {
		return emptyDataModel, a.statusHandler.HandleError(response)
	}

	return a.decodeResponse(response)
}

func (a *Account) Fetch(accountID string) (model.DataModel, error) {
	response, err := a.client.Get(accountID)
	if err != nil {
		return emptyDataModel, err
	}
	defer a.closeBody(response)

	if !a.isFetched(response) {
		return emptyDataModel, a.statusHandler.HandleError(response)
	}

	return a.decodeResponse(response)
}

func (a *Account) Delete(accountID string, version int) error {
	response, err := a.client.Delete(accountID, "version", fmt.Sprint(version))
	if err != nil {
		return err
	}
	defer a.closeBody(response)

	if !a.isDeleted(response) {
		return a.statusHandler.HandleError(response)
	}

	return err
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

func (a *Account) isCreated(response *http.Response) bool {
	return a.statusHandler.StatusCreated(response)
}

func (a *Account) isFetched(response *http.Response) bool {
	return a.statusHandler.StatusOK(response)
}

func (a *Account) isDeleted(response *http.Response) bool {
	return a.statusHandler.StatusNoContent(response)
}

func (a *Account) closeBody(response *http.Response) {
	if response != nil && response.Body != nil {
		response.Body.Close()
	}
}
