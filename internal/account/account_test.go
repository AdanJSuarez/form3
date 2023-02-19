package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

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
	dataModelByte, _ = json.Marshal(dataModelResponse)
)

type TSAccount struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSAccount))
}

func (ts *TSAccount) BeforeTest(_, _ string) {
	accountTest = New(*baseURLTest, accountPath)
	ts.IsType(new(Account), accountTest)
	clientMock = new(MockClient)
	accountTest.client = clientMock
}

func (ts *TSAccount) TestCreateValidDataModel() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.Anything).Return(true)

	data, err := accountTest.Create(dataModelRequest)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestCreateInvalidDataModel() {
	res := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.Anything).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 400: fakeError"))

	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

func (ts *TSAccount) TestCreateNilResponseWithError() {
	clientMock.On("Post", mock.Anything).Return(nil, fmt.Errorf("fakeError"))
	data, err := accountTest.Create(dataModelRequest)
	ts.ErrorContains(err, "fakeError")
	ts.Empty(data)
}

func (ts *TSAccount) TestCreateDecodeError() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)
	clientMock.On("StatusCreated", mock.Anything).Return(true)

	data, err := accountTest.Create(dataModelRequest)
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchValidAccount() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)
	clientMock.On("StatusOK", mock.Anything).Return(true)

	data, err := accountTest.Fetch("fakeID")
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestFetchInvalidAccount() {
	res := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Get", mock.AnythingOfType("string")).Return(res, nil)
	clientMock.On("StatusOK", mock.Anything).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 404: fakeErrorFetch"))

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "status code 404: fakeErrorFetch")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchErrorReturned() {
	clientMock.On("Get", mock.Anything).Return(nil, fmt.Errorf("fakeErrorReturnedOnFetch"))

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "fakeErrorReturnedOnFetch")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchDecodeError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)
	clientMock.On("StatusOK", mock.Anything).Return(true)

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestDeleteValidAccount() {
	res := &http.Response{
		StatusCode: 204,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.Anything).Return(true)

	err := accountTest.Delete("fakeID", 0)
	ts.NoError(err)
}

func (ts *TSAccount) TestDeleteInvalidAccount() {
	res := &http.Response{
		StatusCode: 404,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.Anything).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 404: fakeError"))

	err := accountTest.Delete("fakeID", 0)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDeleteInvalidVersion() {
	res := &http.Response{
		StatusCode: 404,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
	clientMock.On("StatusNoContent", mock.Anything).Return(false)
	clientMock.On("HandleError", mock.Anything).Return(fmt.Errorf("status code 404: fakeError"))

	err := accountTest.Delete("fakeID", 7)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDecodeResponse() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	data, err := accountTest.decodeResponse(res)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}
func (ts *TSAccount) TestDecodeResponseInvalid() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	data, err := accountTest.decodeResponse(res)
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestDecodeResponseNilResponse() {
	data, err := accountTest.decodeResponse(nil)
	ts.ErrorContains(err, "http response is nil")
	ts.Empty(data)
}

func (ts *TSAccount) TestCloseBody() {
	req := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}
func (ts *TSAccount) TestCloseBodyNilBodyNoPanic() {
	req := &http.Response{
		StatusCode: 200,
		Body:       nil,
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}

func (ts *TSAccount) TestCloseBodyNilResponseNoPanic() {
	ts.NotPanics(func() { accountTest.closeBody(nil) })
}
