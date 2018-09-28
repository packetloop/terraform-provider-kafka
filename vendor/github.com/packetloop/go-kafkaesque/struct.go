package gokafkaesque

// Topic includes Kafka topic config, partitions, replication
// factor and name.
type Topic struct {
	*Config           `json:"config"`
	Partitions        string  `json:"partitions"`
	ReplicationFactor string  `json:"replicationFactor"`
	Name              *string `json:"topic"`
}

// Config contains a Kafka topic retention config in ms.
type Config struct {
	RetentionMs       string `json:"retention.ms"`
	SegmentBytes      string `json:"segment.bytes"`
	CleanupPolicy     string `json:"cleanup.policy"` // Accepted values are: "deleted", "compact"
	MinInsyncReplicas string `json:"min.insync.replicas"`
	RetentionBytes    string `json:"retention.bytes"`
	SegmentMs         string `json:"segment.ms"`
}

// AllTopics is a list of topic names.
type AllTopics struct {
	Topics []string `json:"topics"`
}

// Response returns a response of OK.
type Response struct {
	Message string `json:"message"`
}
