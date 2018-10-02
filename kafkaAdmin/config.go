package kafkaadmin

import (
	kafka "github.com/comozo/go-kafkaesque"
)

// Conn is the client connection manager for singularity provider.
// It holds the connection information such as API endpoint to interface with.
type Conn struct {
	sclient *kafka.Client
}

// Config holds the provider configuration, and delivers a populated
// singularity connection based off the contained settings.
type Config struct {
	HostURL string
	Retry   int
}

// Client returns a new client for accessing Singularity Rest API.
// We don't do any authorisation as of the moment. Hence, this block
// is simple.
func (c *Config) Client() (*Conn, error) {
	cf := kafka.NewConfig().
		SetURL(c.HostURL).
		Build()

	client := kafka.NewClient(cf)

	return &Conn{
		sclient: client,
	}, nil
}
