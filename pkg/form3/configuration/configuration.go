package configuration

import (
	"fmt"
	"net/url"
)

type configuration struct {
	baseURL     *url.URL
	accountPath string
}

func New() Configuration {
	return &configuration{}
}

func (c *configuration) InitializeByValue(rawBaseURL, accountPath string) error {
	baseURL, err := c.parseRawBaseURL(rawBaseURL)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	c.accountPath = accountPath

	return nil
}

func (c *configuration) BaseURL() *url.URL {
	return c.baseURL
}

func (c *configuration) AccountPath() string {
	return c.accountPath
}

func (c *configuration) InitializeByYaml() error {
	//TODO: Implement config from a yaml file
	return fmt.Errorf("not implemented")
}

func (c *configuration) InitializeByEnv() error {
	// TODO: Implement config from environment variables
	return fmt.Errorf("not implemented")
}

func (c *configuration) parseRawBaseURL(rawBaseURL string) (*url.URL, error) {
	url, err := url.ParseRequestURI(rawBaseURL)
	if err != nil || url == nil {
		return nil, fmt.Errorf("failed parsing rawBaseURL: %v", err)
	}
	return url, nil
}
