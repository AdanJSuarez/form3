package configuration

import (
	"fmt"
	"net/url"
)

type Configuration struct {
	baseURL        *url.URL
	accountPath    string
	organizationID string
}

func New() *Configuration {
	return &Configuration{}
}

func (c *Configuration) InitializeByValue(rawBaseURL, accountPath, organizationID string) error {
	baseURL, err := c.parseRawBaseURL(rawBaseURL)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	c.accountPath = accountPath
	c.organizationID = organizationID

	return nil
}

func (c *Configuration) BaseURL() *url.URL {
	return c.baseURL
}

func (c *Configuration) AccountPath() string {
	return c.accountPath
}

func (c *Configuration) OrganizationID() string {
	return c.organizationID
}

func (c *Configuration) InitializeByYaml() error {
	//TODO: Implement config from a yaml file
	return fmt.Errorf("not implemented")
}

func (c *Configuration) InitializeByEnv() error {
	// TODO: Implement config from environment variables
	return fmt.Errorf("not implemented")
}

func (c *Configuration) parseRawBaseURL(rawBaseURL string) (*url.URL, error) {
	url, err := url.ParseRequestURI(rawBaseURL)
	if err != nil || url == nil {
		return nil, fmt.Errorf("failed parsing rawBaseURL: %v", err)
	}
	return url, nil
}
