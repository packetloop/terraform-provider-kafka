provider "kafka" {
  zookeeper = "localhost"
  kafka_bin_path = "/Users/alexey/Work/kafka/bin"
}

resource "kafka_topic" "test7" {
  name = "test8"
  partitions = 3
  replication_factor = 1
}
