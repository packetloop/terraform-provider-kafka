package kafkaadmin

import (
	"fmt"
	"testing"
)

func TestValidateCleanupPolicy(t *testing.T) {
	var data = []struct {
		topic         string
		cleanupPolicy string
	}{
		{"kafka_topic.foo", "compact"},
		{"kafka_topic.foo", "delete"},
	}

	for _, tt := range data {
		_, err := validateCleanupPolicy(tt.cleanupPolicy, tt.topic)
		if err != nil {
			fmt.Printf("got %v expected nil \n", err)
			t.FailNow()
		}
	}
}

func TestInvalidCleanupPolicy(t *testing.T) {
	var data = []struct {
		topic         string
		cleanupPolicy string
	}{
		{"kafka_topic.foo", "xxxxxxx"},
	}

	for _, tt := range data {
		_, err := validateCleanupPolicy(tt.cleanupPolicy, tt.topic)
		if err == nil {
			fmt.Printf("got %v expected nil \n", err)
			t.FailNow()
		}
	}
}
