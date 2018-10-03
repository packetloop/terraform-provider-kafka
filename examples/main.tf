provider "kafkaadmin" {
  HOST_URL = "http://localhost:8080"
}

resource "kafka_topic" "foobar" {
  name                = "mytopicconfig"
  partitions          = 2
  replication_factor  = 3
  retention_ms        = 300000
  cleanup_policy      = "delete"
  segment_bytes       = 10737418
  min_insync_replicas = 2
  retention_bytes     = 100000
  segment_ms          = 600000
}
