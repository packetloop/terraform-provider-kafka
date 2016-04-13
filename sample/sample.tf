provider "kafka" {
  zookeeper = "192.168.99.100"
}

resource "kafka_topic" "my-topic" {
  name = "my-topic"
  partitions = 2
  replication_factor = 1
  retention_ms = 300000
  cleanup_policy = "compact"
}
