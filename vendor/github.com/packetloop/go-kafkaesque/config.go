package gokafkaesque

import (
	"github.com/go-resty/resty"
)

// Client contains Singularity endpoint for http requests
type Client struct {
	Rest *resty.Client
}

// serverConfig contains Kafka-admin-service HTTP endpoint and serverConfiguration for
// retryablehttp client's retry options
type serverConfig struct {
	URL   string
	Retry int
}

// ServerConfigBuilder sets port, host, http retry count config to
// be passed to create a NewClient.
type ServerConfigBuilder interface {
	SetURL(string) ServerConfigBuilder
	SetRetry(int) ServerConfigBuilder
	Build() serverConfig
}

// NewConfig returns an empty ServerConfigBuilder.
func NewConfig() ServerConfigBuilder {
	return &serverConfig{}
}

// SetHost accepts a string in the form of http://url and sets
// this as URL in serverConfig.
func (co *serverConfig) SetURL(URL string) ServerConfigBuilder {
	co.URL = URL
	return co
}

// SetHost accepts an int and sets the retry count.
func (co *serverConfig) SetRetry(r int) ServerConfigBuilder {
	co.Retry = r
	return co
}

// Build method returns a serverConfig struct.
func (co *serverConfig) Build() serverConfig {

	return serverConfig{
		URL:   co.URL,
		Retry: co.Retry,
	}
}

// NewClient returns Singularity HTTP endpoint.
func NewClient(c serverConfig) *Client {
	r := resty.New().
		SetRESTMode().
		SetRetryCount(c.Retry).
		SetHostURL(c.URL)
	return &Client{
		Rest: r,
	}
}
