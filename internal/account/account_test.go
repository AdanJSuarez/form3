package account

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/AdanJSuarez/form3/internal/client/request"
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
	returnedBodyError = io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError")))
)

type TSAccount struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSAccount))
}

func (ts *TSAccount) BeforeTest(_, _ string) {
	accountTest = New(*baseURLTest, accountPath)
	ts.IsType(new(Account), accountTest)
	clientMock = new(MockClient)
}

func (ts *TSAccount) TestCreateValidDataModel() {
	req := request.NewRequestHandler(dataModelResponse)
	res := &http.Response{
		StatusCode: 201,
		Body:       req.Body(),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock

	data, err := accountTest.Create(dataModelRequest)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestCreateInvalidDataModel() {
	res := &http.Response{
		StatusCode: 400,
		Body:       returnedBodyError,
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.AnythingOfType("*http.Response")).Return(false)
	clientMock.On("HandleError", mock.AnythingOfType("*http.Response")).Return(fmt.Errorf("status code 400: fakeError"))
	accountTest.client = clientMock

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

func (ts *TSAccount) TestCreateDecodeError() {
	res := &http.Response{
		StatusCode: 201,
		Body:       returnedBodyError,
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock

	data, err := accountTest.Create(dataModelRequest)
	ts.Error(err)
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchValidAccount() {
	req := request.NewRequestHandler(dataModelResponse)
	res := &http.Response{
		StatusCode: 200,
		Body:       req.Body(),
	}
	clientMock.On("Get", mock.AnythingOfType("string")).Return(res, nil)
	clientMock.On("StatusOK", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock

	data, err := accountTest.Fetch("fakeID")
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestFetchInvalidAccount() {
	res := &http.Response{
		StatusCode: 404,
		Body:       returnedBodyError,
	}
	clientMock.On("Get", mock.AnythingOfType("string")).Return(res, nil)
	clientMock.On("StatusOK", mock.AnythingOfType("*http.Response")).Return(false)
	clientMock.On("HandleError", mock.AnythingOfType("*http.Response")).Return(fmt.Errorf("status code 404: fakeError"))
	accountTest.client = clientMock

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "status code 404:")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchErrorReturned() {
	clientMock.On("Get", mock.AnythingOfType("string")).Return(nil, fmt.Errorf("fakeErrorReturned"))
	accountTest.client = clientMock

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "fakeErrorReturned")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchDecodeError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Get", mock.AnythingOfType("string")).Return(res, nil)
	clientMock.On("StatusOK", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock

	data, err := accountTest.Fetch("fakeID")
	ts.Error(err)
	ts.Empty(data)
}

func (ts *TSAccount) TestDeleteValidAccount() {
	res := &http.Response{
		StatusCode: 204,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.AnythingOfType("*http.Response")).Return(true)
	accountTest.client = clientMock

	err := accountTest.Delete("fakeID", 0)
	ts.NoError(err)
}

func (ts *TSAccount) TestDeleteInvalidAccount() {
	res := &http.Response{
		StatusCode: 404,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.AnythingOfType("*http.Response")).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 404: fakeError"))
	accountTest.client = clientMock

	err := accountTest.Delete("fakeID", 0)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDeleteInvalidVersion() {
	res := &http.Response{
		StatusCode: 404,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.AnythingOfType("*http.Response")).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 404: fakeError"))
	accountTest.client = clientMock

	err := accountTest.Delete("fakeID", 7)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDecodeResponse() {
	reqBody := request.NewRequestHandler(dataModelResponse)
	res := &http.Response{
		StatusCode: 200,
		Body:       reqBody.Body(),
	}
	data, err := accountTest.decodeResponse(res)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}
func (ts *TSAccount) TestDecodeResponseInvalid() {
	reqBody := request.NewRequestHandler("SantaUrsulaCapitalDelUniverso")
	res := &http.Response{
		StatusCode: 200,
		Body:       reqBody.Body(),
	}
	data, err := accountTest.decodeResponse(res)
	ts.Error(err)
	ts.Empty(data)
}

func (ts *TSAccount) TestDecodeResponseNilResponse() {
	data, err := accountTest.decodeResponse(nil)
	ts.Error(err)
	ts.Empty(data)
}

func (ts *TSAccount) TestCloseBody() {
	reqBody := request.NewRequestHandler(dataModelResponse)
	req := &http.Response{
		StatusCode: 200,
		Body:       reqBody.Body(),
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}
func (ts *TSAccount) TestCloseBodyNilBody() {
	req := &http.Response{
		StatusCode: 200,
		Body:       nil,
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}

func (ts *TSAccount) TestCloseBodyNilResponse() {
	ts.NotPanics(func() { accountTest.closeBody(nil) })
}
