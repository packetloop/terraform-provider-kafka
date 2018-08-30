package gokafkaesque

// TopicResponse have information about a Kafka topic.
type TopicResponse struct {
	Topic `json:"response"`
}

// Topic includes Kafka topic config, partitions, replication
// factor and name.
type Topic struct {
	*Config           `json:"config"`
	Partitions        int64   `json:"partitions"`
	ReplicationFactor int64   `json:"replicationFactor"`
	Name              *string `json:"name"`
}

// Config contains a Kafka topic retention config in ms.
type Config struct {
	RetentionMs       string `json:"retention.ms"`
	SegmentBytes      string `json:"segment.bytes"`
	CleanupPolicy     string `json:"cleanup.policy"`
	MinInsyncReplicas string `json:"min.insync.replicas"`
	RetentionBytes    string `json:"retention.bytes"`
	SegmentMs         string `json:"segment.ms"`
}

// Topics is a list of topic names.
type Topics struct {
	Response struct {
		Topics []string `json:"topics"`
	} `json:"response"`
}

// GenericResponse returns a response of OK.
type GenericResponse struct {
	Response string `json:"response"`
}
