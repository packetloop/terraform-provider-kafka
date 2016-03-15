provider "kafka" {
  zookeeper = "localhost"
  kafka_bin_path = "/Users/alexey/Work/kafka/bin"
}

resource "kafka_topic" "my-test" {
  name = "my-test"
  partitions = 2
  replication_factor = 1
  retention_ms = 300000
  cleanup_policy = "compact"
}
