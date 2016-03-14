package main

import (
	"testing"
)

const (
	validDescribeResponse =
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

	shortDescribeResponse = "Topic:file-imported	PartitionCount:12	ReplicationFactor:3	Configs:"

	emptyDescribeResponse = ""

	invalidDescribeResponse = "some unknown stuff"

  noBrokersError =
`Error while executing topic command : replication factor: 1 larger than available brokers: 0
[2016-03-14 15:19:38,777] ERROR kafka.admin.AdminOperationException: replication factor: 1 larger than available brokers: 0
	at kafka.admin.AdminUtils$.assignReplicasToBrokers(AdminUtils.scala:77)
	at kafka.admin.AdminUtils$.createTopic(AdminUtils.scala:236)
	at kafka.admin.TopicCommand$.createTopic(TopicCommand.scala:105)
	at kafka.admin.TopicCommand$.main(TopicCommand.scala:60)
	at kafka.admin.TopicCommand.main(TopicCommand.scala)
 (kafka.admin.TopicCommand$)`

  topicExistsError =
`Error while executing topic command : Topic "test" already exists.
[2016-03-14 16:03:58,893] ERROR kafka.common.TopicExistsException: Topic "test" already exists.
	at kafka.admin.AdminUtils$.createOrUpdateTopicPartitionAssignmentPathInZK(AdminUtils.scala:253)
	at kafka.admin.AdminUtils$.createTopic(AdminUtils.scala:237)
	at kafka.admin.TopicCommand$.createTopic(TopicCommand.scala:105)
	at kafka.admin.TopicCommand$.main(TopicCommand.scala:60)
	at kafka.admin.TopicCommand.main(TopicCommand.scala)
 (kafka.admin.TopicCommand$)`
)

func TestKafkaManagingClient_topicInfo(t *testing.T) {
  res, err := readTopicInfo(validDescribeResponse)
  if err != nil { t.Fatal(err) }
  if res.PartitionsCount != 12 { t.Errorf("expected PartitionsCount to be %d, but got %d", 12, res.PartitionsCount) }
  if res.ReplicationFactor != 3 { t.Errorf("expected ReplicationFactor to be %d, but got %d", 3, res.ReplicationFactor) }
}

func TestKafkaManagingClient_shortTopicInfo(t *testing.T) {
  res, err := readTopicInfo(shortDescribeResponse)
  if err != nil { t.Fatal(err) }
  if res.PartitionsCount != 12 { t.Errorf("expected PartitionsCount to be %d, but got %d", 12, res.PartitionsCount) }
  if res.ReplicationFactor != 3 { t.Errorf("expected ReplicationFactor to be %d, but got %d", 3, res.ReplicationFactor) }
}

func TestKafkaManagingClient_emptyTopicInfo(t *testing.T) {
  _, err := readTopicInfo(emptyDescribeResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}

func TestKafkaManagingClient_invalidResponseTopicInfo(t *testing.T) {
  _, err := readTopicInfo(invalidDescribeResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}

func TestKafkaManagingClient_parseNoBrokersError(t *testing.T) {
  err := readError(noBrokersError)

  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
  if err.Error() != "Error while executing topic command : replication factor: 1 larger than available brokers: 0" {
    t.Errorf("Unexpected error message: '%s'", err.Error())
  }
}


func TestKafkaManagingClient_parseTopicExistsError(t *testing.T) {
  err := readError(topicExistsError)

  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
  if err.Error() != "Error while executing topic command : Topic \"test\" already exists." {
    t.Errorf("Unexpected error message: '%s'", err.Error())
  }
}