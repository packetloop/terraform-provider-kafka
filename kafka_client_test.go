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

	shortDescribeResponse  = "Topic:file-imported	PartitionCount:12	ReplicationFactor:3	Configs:retention.ms=1457999337,cleanup.policy=compact"
	shortDescribeResponse2 = "Topic:file-imported	PartitionCount:12	ReplicationFactor:3	Configs:retention.bytes=1023"

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

  errorWithWarnings =
`WARNING: If partitions are increased for a topic that has a key, the partition logic or ordering of the messages will be affected
Error while executing topic command : The number of partitions for a topic can only be increased
[2016-03-15 15:20:05,571] ERROR kafka.admin.AdminOperationException: The number of partitions for a topic can only be increased
	at kafka.admin.AdminUtils$.addPartitions(AdminUtils.scala:119)
	at kafka.admin.TopicCommand$$anonfun$alterTopic$1.apply(TopicCommand.scala:139)
	at kafka.admin.TopicCommand$$anonfun$alterTopic$1.apply(TopicCommand.scala:116)
	at scala.collection.mutable.ResizableArray$class.foreach(ResizableArray.scala:59)
	at scala.collection.mutable.ArrayBuffer.foreach(ArrayBuffer.scala:48)
	at kafka.admin.TopicCommand$.alterTopic(TopicCommand.scala:116)
	at kafka.admin.TopicCommand$.main(TopicCommand.scala:62)
	at kafka.admin.TopicCommand.main(TopicCommand.scala)
 (kafka.admin.TopicCommand$)`
)

func TestKafkaManagingClient_topicInfo(t *testing.T) {
  res, err := readTopicInfo(validDescribeResponse)
  if err != nil { t.Fatal(err) }
  assertInt   (t, "PartitionsCount",   res.PartitionsCount,   12)
  assertInt   (t, "ReplicationFactor", res.ReplicationFactor, 3 )
  assertInt64 (t, "RetentionBytes",    res.RetentionBytes,    -1)
  assertInt64 (t, "RetentionMs",       res.RetentionMs,       -1)
  assertString(t, "CleanupPolicy",     res.CleanupPolicy,     "")
}

func TestKafkaManagingClient_shortTopicInfo(t *testing.T) {
  var res, err = readTopicInfo(shortDescribeResponse)
  if err != nil { t.Fatal(err) }
  assertInt   (t, "PartitionsCount",   res.PartitionsCount,   12)
  assertInt   (t, "ReplicationFactor", res.ReplicationFactor, 3 )
  assertInt64 (t, "RetentionBytes",    res.RetentionBytes,    -1)
  assertInt64 (t, "RetentionMs",       res.RetentionMs,       1457999337)
  assertString(t, "CleanupPolicy",     res.CleanupPolicy,     "compact")

  res, err = readTopicInfo(shortDescribeResponse2)
  if err != nil { t.Fatal(err) }
  assertInt   (t, "PartitionsCount",   res.PartitionsCount,   12)
  assertInt   (t, "ReplicationFactor", res.ReplicationFactor, 3 )
  assertInt64 (t, "RetentionBytes",    res.RetentionBytes,    1023)
  assertInt64 (t, "RetentionMs",       res.RetentionMs,       -1)
  assertString(t, "CleanupPolicy",     res.CleanupPolicy,     "")
}

func assertInt(t *testing.T, name string, value int, expected int) {
  if (expected != value) {
    t.Errorf("expected %s to be %d, but got %d", name, expected, value)
  }
}

func assertInt64(t *testing.T, name string, value int64, expected int64) {
  if (expected != value) {
    t.Errorf("expected %s to be %d, but got %d", name, expected, value)
  }
}

func assertString(t *testing.T, name string, value string, expected string) {
  if (expected != value) {
    t.Errorf("expected %s to be %s, but got %s", name, expected, value)
  }
}

func TestKafkaManagingClient_emptyTopicInfo(t *testing.T) {
  _, err := readTopicInfo(emptyDescribeResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}

func TestKafkaManagingClient_invalidResponseTopicInfo(t *testing.T) {
  _, err := readTopicInfo(invalidDescribeResponse)
  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
}

func TestKafkaManagingClient_errorWithWarnings(t *testing.T) {
  err := readError(errorWithWarnings)

  if err == nil { t.Fatal("Error is expected, but success found. Sometimes success is not what you are after.") }
  if err.Error() != "Error while executing topic command : The number of partitions for a topic can only be increased" {
    t.Errorf("Unexpected error message: '%s'", err.Error())
  }
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