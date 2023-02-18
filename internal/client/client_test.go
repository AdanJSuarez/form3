package client

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	rawBaseURLTest = "https://api.fakeaddress.tech/fake"
	valueURL       = "/v1/organisation/accounts"
	idTest         = "020cf7d8-01b9-461d-89d4-89d57fd0d998"
)

var (
	clientURLTest *url.URL
	clientTest    *Client
)

type TSClient struct{ suite.Suite }

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(TSClient))
}

func (ts *TSClient) BeforeTest(_, _ string) {
	clientURLTest, _ = url.ParseRequestURI(rawBaseURLTest)
	clientTest = New(*clientURLTest)
	ts.IsType(new(Client), clientTest)
}

func (ts *TSClient) TestJoinValueToURL() {
	url, err := clientTest.joinValuesToURL(idTest)
	ts.NoError(err)
	ts.NotEmpty(url)
}
