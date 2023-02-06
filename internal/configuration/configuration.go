package configuration

import (
	"fmt"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/configuration/validation"
)

const accountPath = "/v1/organisation/accounts"

type Configuration struct {
	baseURL        string
	accountURL     string
	organizationID string
}

func New(baseURL, organizationID string) (*Configuration, error) {
	_, err := validation.NewValidation(baseURL)
	if err != nil {
		return nil, err
	}

	accountURL, err := joinURLAndPath(baseURL, accountPath)
	if err != nil {
		return nil, err
	}

	return &Configuration{
		baseURL:        baseURL,
		accountURL:     accountURL,
		organizationID: organizationID,
	}, nil
}

func NewByYaml() (*Configuration, error) {
	//TODO: Implement config from a yaml file
	return nil, fmt.Errorf("not implemented NewByYaml")
}

func NewByEnv() (*Configuration, error) {
	// TODO: Implement config from environment variables
	return nil, fmt.Errorf("not implemented NewByEnv")
}

func (c *Configuration) AccountURL() string {
	return c.baseURL
}

func (c *Configuration) OrganizationID() string {
	return c.organizationID
}

// joinURLAndPath returns the joined URL. An error otherwise
func joinURLAndPath(baseURL, path string) (string, error) {
	url, err := url.JoinPath(baseURL, path)
	if err != nil {
		return "", err
	}
	return url, nil
}
