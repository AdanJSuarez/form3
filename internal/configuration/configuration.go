package configuration

import (
	"fmt"
	"net/url"
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
	if err := c.validateBaseURL(baseURL); err != nil {
		return err
	}

	accountURL, err := c.joinPathToBaseURL(baseURL, accountPath)
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
	return fmt.Errorf("not implemented")
}

func (c *Configuration) InitializeByEnv() error {
	// TODO: Implement config from environment variables
	return fmt.Errorf("not implemented")
}

func (c *Configuration) AccountURL() string {
	return c.accountURL
}

func (c *Configuration) OrganizationID() string {
	return c.organizationID
}

func (c *Configuration) validateBaseURL(baseURL string) error {
	_, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) joinPathToBaseURL(baseURL, path string) (string, error) {
	url, err := url.JoinPath(baseURL, path)
	if err != nil {
		return "", err
	}
	return url, nil
}
