package account

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/AdanJSuarez/form3/internal/client"
	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	baseURL        = "http://fakeurl.com"
	accountPath    = "/v1/origanisation/"
	uuidTest       = "123e4567-e89b-12d3-a456-426614174000"
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
)

var (
	accountTest        *Account
	clientMock         *MockClient
	statusHandlerMock  *MockStatusHandler
	baseURLTest, _     = url.Parse(baseURL)
	dataAttributesTest = model.Attributes{
		Country:      "GB",
		BaseCurrency: "GBP",
		BankID:       "123456",
		BankIDCode:   "GBDSC",
		Bic:          "EXMPLGB2XXX",
	}

	dataTest = model.Data{
		ID:             uuidTest,
		OrganizationID: organizationID,
		Type:           "accounts",
		Attributes:     dataAttributesTest,
	}
	dataModelRequest = model.DataModel{
		Data: dataTest,
	}

	dataModelResponse = model.DataModel{
		Data: dataTest,
	}
)

type TSAccount struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSAccount))
}

func (ts *TSAccount) BeforeTest(_, _ string) {
	accountTest = New(*baseURLTest, accountPath)
	ts.IsType(new(Account), accountTest)
	clientMock = new(MockClient)
	statusHandlerMock = new(MockStatusHandler)
}

func (ts *TSAccount) TestCreateValidDataModel() {
	req := client.NewRequestBody(dataModelResponse)
	res := &http.Response{
		StatusCode: 201,
		Body:       req.Body(),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	statusHandlerMock.On("StatusCreated", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock
	accountTest.statusHandler = statusHandlerMock

	data, err := accountTest.Create(dataModelRequest)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestCreateInvalidDataModel() {
	res := &http.Response{
		StatusCode: 400,
		Body:       nil,
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	statusHandlerMock.On("StatusCreated", mock.AnythingOfType("*http.Response")).Return(false)
	statusHandlerMock.On("HandleError", mock.AnythingOfType("*http.Response")).Return(fmt.Errorf("status code 400: fakeError"))
	accountTest.client = clientMock
	accountTest.statusHandler = statusHandlerMock

	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

func (ts *TSAccount) TestCreateNilResponse() {
	clientMock.On("Post", mock.Anything).Return(nil, fmt.Errorf("fakeError"))
	accountTest.client = clientMock
	data, err := accountTest.Create(dataModelRequest)
	ts.ErrorContains(err, "fakeError")
	ts.Empty(data)
}
