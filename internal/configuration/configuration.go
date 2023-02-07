package configuration

import (
	"fmt"
	"net/url"

	"github.com/AdanJSuarez/form3/internal/configuration/validation"
)

type Configuration struct {
	baseURL        string
	accountURL     string
	organizationID string
}

func New() *Configuration {
	return &Configuration{}
}

func (c *Configuration) InitializeByValue(baseURL, accountPath, organizationID string) error {
	_, err := validation.NewValidation(baseURL)
	if err != nil {
		return err
	}

	accountURL, err := joinURLAndPath(baseURL, accountPath)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	c.accountURL = accountURL
	c.organizationID = organizationID

	return nil
}

func (c *Configuration) InitializeByYaml() error {
	//TODO: Implement config from a yaml file
	return fmt.Errorf("not implemented NewByYaml")
}

func (c *Configuration) InitializeByEnv() error {
	// TODO: Implement config from environment variables
	return fmt.Errorf("not implemented NewByEnv")
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
