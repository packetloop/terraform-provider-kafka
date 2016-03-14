package main

import (
	"testing"
)

const (
	validResponse =
`Topic:file-imported	PartitionCount:12	ReplicationFactor:3	Configs:
Topic: file-imported	Partition: 0	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 1	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 2	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 3	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 4	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 5	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 6	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 7	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 8	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 9	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 10	Leader: -1	Replicas: 0	Isr:
Topic: file-imported	Partition: 11	Leader: -1	Replicas: 0	Isr:`

	shortResponse = "Topic:file-imported	PartitionCount:12	ReplicationFactor:3	Configs:"

	emptyResponse = ""

	invalidResponse = "some unknown stuff"
)

func TestKafkaManagingClient_topicInfo(t *testing.T) {
  res, err := readTopicInfo(validResponse)
  if err != nil { t.Fatal(err) }
  if res.PartitionsCount != 12 { t.Errorf("expected PartitionsCount to be %d, but got %d", 12, res.PartitionsCount) }
  if res.ReplicationFactor != 3 { t.Errorf("expected ReplicationFactor to be %d, but got %d", 3, res.ReplicationFactor) }
}

func TestKafkaManagingClient_shortTopicInfo(t *testing.T) {
  res, err := readTopicInfo(shortResponse)
  if err != nil { t.Fatal(err) }
  if res.PartitionsCount != 12 { t.Errorf("expected PartitionsCount to be %d, but got %d", 12, res.PartitionsCount) }
  if res.ReplicationFactor != 3 { t.Errorf("expected ReplicationFactor to be %d, but got %d", 3, res.ReplicationFactor) }
}

func TestKafkaManagingClient_emptyTopicInfo(t *testing.T) {
  _, err := readTopicInfo(emptyResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}

func TestKafkaManagingClient_invalidResponseTopicInfo(t *testing.T) {
  _, err := readTopicInfo(invalidResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}