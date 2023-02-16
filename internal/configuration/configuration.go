package configuration

import (
	"fmt"
	"net/url"
	"os"
)

const (
	baseURLEnvKey     = "BASE_URL"
	accountPathEnvKey = "ACCOUNT_PATH"
	errorEnvFmt       = "failed to get %s from environment variables"
)

type Configuration struct {
	baseURL     *url.URL
	accountPath string
}

func New() *Configuration {
	return &Configuration{}
}

func (c *Configuration) InitializeByValue(rawBaseURL, accountPath string) error {
	baseURL, err := c.parseRawBaseURL(rawBaseURL)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	c.accountPath = accountPath

	return nil
}

func (c *Configuration) BaseURL() *url.URL {
	return c.baseURL
}

func (c *Configuration) AccountPath() string {
	return c.accountPath
}

func (c *Configuration) InitializeByEnv() error {
	rawBaseURL, ok := os.LookupEnv(baseURLEnvKey)
	if !ok {
		return fmt.Errorf(errorEnvFmt, baseURLEnvKey)
	}
	accountPath, ok := os.LookupEnv(accountPathEnvKey)
	if !ok {
		return fmt.Errorf(errorEnvFmt, accountPathEnvKey)
	}

	baseURL, err := c.parseRawBaseURL(rawBaseURL)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	c.accountPath = accountPath
	return nil
}

func (c *Configuration) parseRawBaseURL(rawBaseURL string) (*url.URL, error) {
	url, err := url.ParseRequestURI(rawBaseURL)
	if err != nil || url == nil {
		return nil, fmt.Errorf("failed parsing rawBaseURL: %v", err)
	}
	return url, nil
}
