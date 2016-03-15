provider "kafka" {
  zookeeper = "localhost"
  kafka_bin_path = "/Users/alexey/Work/kafka/bin"
}

resource "kafka_topic" "my-test" {
  name = "test12"
  partitions = 3
  replication_factor = 1
  retention_ms = 86000
  cleanup_policy = "compact"
}
