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
	clientMock = NewMockClient(ts.T())
	accountTest.client = clientMock
}

func (ts *TSAccount) TestCreateValidDataModel() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)

	data, err := accountTest.Create(dataModelRequest)
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestCreateInvalidDataModel() {
	clientMock.On("Post", mock.Anything).Return(nil, fmt.Errorf(statusBadRequestMsg))

	data, err := accountTest.Create(model.DataModel{})
	ts.ErrorContains(err, "status code 400:")
	ts.Empty(data)
}

func (ts *TSAccount) TestCreateDecodeError() {
	res := &http.Response{
		StatusCode: 201,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Post", mock.Anything).Return(res, nil)

	data, err := accountTest.Create(dataModelRequest)
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchValidID() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(dataModelByte)),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)

	data, err := accountTest.Fetch("fakeID")
	ts.NoError(err)
	ts.Equal(dataModelResponse, data)
}

func (ts *TSAccount) TestFetchNotFoundID() {
	clientMock.On("Get", mock.AnythingOfType("string")).Return(nil,
		fmt.Errorf(statusNotFoundMsg))

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "status code 404:")
	ts.Empty(data)
}

func (ts *TSAccount) TestFetchDecodeError() {
	res := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer([]byte("fakeReturnedBodyError"))),
	}
	clientMock.On("Get", mock.Anything).Return(res, nil)

	data, err := accountTest.Fetch("fakeID")
	ts.ErrorContains(err, "invalid character")
	ts.Empty(data)
}

func (ts *TSAccount) TestDeleteValidAccount() {
	res := &http.Response{
		StatusCode: 204,
		Body:       nil,
	}
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(res, nil)

	err := accountTest.Delete("fakeID", 0)
	ts.NoError(err)
}

func (ts *TSAccount) TestDeleteNotFoundAccount() {
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf(statusNotFoundMsg))

	err := accountTest.Delete("fakeID", 0)
	ts.ErrorContains(err, "status code 404:")
}

func (ts *TSAccount) TestDeleteInvalidVersion() {
	clientMock.On("Delete", mock.Anything, mock.Anything,
		mock.Anything).Return(nil, fmt.Errorf(statusNotFoundMsg))

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
func (ts *TSAccount) TestDecodeResponseInvalidData() {
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
