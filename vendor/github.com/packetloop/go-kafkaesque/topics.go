package gokafkaesque

import (
	"fmt"
)

func callError(msg string) func(error) error {
	return func(err error) error {
		return fmt.Errorf("%s ERROR: %v", msg, err.Error())
	}
}

// GetTopics is a method that returns all Kafka topics.
func (client *Client) GetTopics() (Topics, error) {
	e := callError("LIST TOPICS")

	resp, err := client.Rest.R().Get("/topics")
	if err != nil {
		return Topics{}, e(err)
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data Topics
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return Topics{}, e(err)
		}
		return data, nil
	}
	return Topics{}, e(fmt.Errorf("%v", resp.Status()))
}

// Count is a method that returns total size of topics.
func (t Topics) Count() int {
	return len(t.Response.Topics)
}

// Topics is a method that returns a slice of topics.
func (t Topics) Topics() []string {
	return t.Response.Topics
}

// GetTopic is a method that return a Kafka topic
func (client *Client) GetTopic(t string) (TopicResponse, error) {
	e := callError(fmt.Sprintf("GET TOPIC %s", t))

	resp, err := client.Rest.R().Get(uriPath("/topics", t))
	if err != nil {
		return TopicResponse{}, e(err)
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data TopicResponse
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return TopicResponse{}, e(err)
		}
		return data, nil
	}
	return TopicResponse{}, e(fmt.Errorf("%v", resp.Status()))
}

// GetPartitions is a method that returns partitions of a topic.
func (t Topic) GetPartitions() int64 {
	return t.Partitions
}

// GetReplicationFactor is a method that returns partitions of a topic.
func (t Topic) GetReplicationFactor() int64 {
	return t.ReplicationFactor
}

// GetRetentionMs is a method that returns partitions of a topic.
func (c *Config) GetRetentionMs() string {
	return fmt.Sprintf("%s", c.RetentionMs)

}

// GetSegmentBytes is a method that returns partitions of a topic.
func (c *Config) GetSegmentBytes() string {
	return c.SegmentBytes

}

// GetCleanupPolicy is a method that returns cleanup policy of a topic.
func (c *Config) GetCleanupPolicy() string {
	return fmt.Sprintf("%s", c.CleanupPolicy)
}

// GetSegmentMs is a method that returns the time after which Kafka will
// force the log to rollof a topic.
func (c *Config) GetSegmentMs() string {
	return fmt.Sprintf("%s", c.SegmentMs)
}

// GetRetentionBytes is a method that returns the retention bytes for the topic.
// force the log to rollof a topic.
func (c *Config) GetRetentionBytes() string {
	return fmt.Sprintf("%s", c.RetentionBytes)
}

// GetMinInSyncReplicas is a method that returns the minimum number of insync replicas.
// force the log to rollof a topic.
func (c *Config) GetMinInSyncReplicas() string {
	return fmt.Sprintf("%s", c.MinInsyncReplicas)
}

// TopicBuilder is an interface that builds a Kafka Topic
// Config.
type TopicBuilder interface {
	SetPartitions(int64) TopicBuilder
	SetReplicationFactor(int64) TopicBuilder
	SetConfig(Config) TopicBuilder
	BuildTopic() Topic
}

// NewTopic accepts a string topic name and returns a TopicBuilder interface.
func NewTopic(name string) TopicBuilder {
	return &Topic{
		Name: &name,
	}
}

// SetPartitions is a method that accepts an int64 and sets Topic
// partition.
func (t *Topic) SetPartitions(p int64) TopicBuilder {
	t.Partitions = p
	return t
}

// SetReplicationFactor is a method that accepts an int64 and sets Topic
// replication factor.
func (t *Topic) SetReplicationFactor(r int64) TopicBuilder {
	t.ReplicationFactor = r
	return t
}

// SetConfig is a method that accepts a Config struct and
//  sets Topic config such as retention periond in ms.
func (t *Topic) SetConfig(c Config) TopicBuilder {
	t.Config = &c
	return t
}

// BuildTopic is a method that builds a Topic parameters and pass
// this as an argument when calling CreateTopic method.
func (t *Topic) BuildTopic() Topic {
	return Topic{
		Name:              t.Name,
		ReplicationFactor: t.ReplicationFactor,
		Partitions:        t.Partitions,
	}
}

// CreateTopic accepts a Topic and returns an "Ok" response
// or error.
func (client *Client) CreateTopic(t Topic) (GenericResponse, error) {
	e := callError(fmt.Sprintf("CREATE TOPIC %+v", t))

	resp, err := client.Rest.R().SetBody(t).Post("/topics")
	if err != nil {
		return GenericResponse{}, e(err)
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data GenericResponse
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return GenericResponse{}, e(err)
		}
		return data, nil
	}
	return GenericResponse{}, e(fmt.Errorf("%v", resp.Status()))
}

// DeleteTopic method accepts a string topic, deletes this Kafka topic and
// returns a string response or error.
func (client *Client) DeleteTopic(t string) (GenericResponse, error) {
	e := callError(fmt.Sprintf("DELETE TOPIC %v", t))

	resp, err := client.Rest.R().Delete(uriPath("/topics", t))
	if err != nil {
		return GenericResponse{}, e(err)
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data GenericResponse
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return GenericResponse{}, e(err)
		}
		return data, nil
	}
	return GenericResponse{}, e(fmt.Errorf("%v", resp.Status()))
}

// uriPath function accepts path and topic string and returns a
// valid uri path of /topics/<topic_name>.
func uriPath(p, t string) string {
	return p + "/" + t
}
