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
	baseURL             = "http://fakeurl.com"
	accountPath         = "/v1/origanisation/"
	uuidTest            = "123e4567-e89b-12d3-a456-426614174000"
	organizationID      = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	statusBadRequestMsg = "status code 400: errorCode: 12345 - errorMessage: fake message"
	statusNotFoundMsg   = "status code 404: not found"
)

var (
	accountTest        *Account
	clientMock         *MockClient
	configurationMock  *MockConfiguration
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
	configurationMock = NewMockConfiguration(ts.T())
	configurationMock.On("BaseURL").Return(baseURLTest)
	configurationMock.On("AccountPath").Return(accountPath)
	clientMock = NewMockClient(ts.T())

	accountTest = New(configurationMock)
	ts.IsType(new(Account), accountTest)

	accountTest.client = clientMock
}

func (ts *TSAccount) TestCreateValidDataModelReturnsNoError() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)

	data, err := accountTest.Create(dataModelRequest)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestCreateInvalidDataModelReturnsError() {
	clientMock.On("Post", mock.Anything).Return(nil, fmt.Errorf(statusBadRequestMsg))

	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

func (ts *TSAccount) TestCreateWithDecodeErrorReturnsError() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)

	data, err := accountTest.Create(dataModelRequest)
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchValidIDReturnsNoError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)

	data, err := accountTest.Fetch("fakeID")
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestFetchNotFoundIDReturnsError() {
	clientMock.On("Get", mock.AnythingOfType("string")).Return(nil,
		fmt.Errorf(statusNotFoundMsg))

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "status code 404:")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchDecodeErrorReturnsError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestDeleteValidAccountReturnsNoError() {
	res := &http.Response{
		StatusCode: 204,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(res, nil)

	err := accountTest.Delete("fakeID", 0)
	ts.NoError(err)
}

func (ts *TSAccount) TestDeleteNotFoundAccountReturnsError() {
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf(statusNotFoundMsg))

	err := accountTest.Delete("fakeID", 0)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDeleteInvalidVersionReturnsError() {
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf(statusNotFoundMsg))

	err := accountTest.Delete("fakeID", 7)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestValidResponseReturnsNoError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	data, err := accountTest.decodeResponse(res)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}
func (ts *TSAccount) TestInvalidDataResponseReturnsError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	data, err := accountTest.decodeResponse(res)
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestNilResponseReturnsError() {
	data, err := accountTest.decodeResponse(nil)
	ts.ErrorContains(err, "http response is nil")
	ts.Empty(data)
}

func (ts *TSAccount) TestCloseBodyNotPanicValidResponse() {
	req := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}
func (ts *TSAccount) TestCloseBodyNotPanicNilBody() {
	req := &http.Response{
		StatusCode: 200,
		Body:       nil,
	}
	ts.NotPanics(func() { accountTest.closeBody(req) })
}

func (ts *TSAccount) TestCloseBodyNotNoPanicNilResponse() {
	ts.NotPanics(func() { accountTest.closeBody(nil) })
}
