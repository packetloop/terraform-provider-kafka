
resource "kafka_on_demand" "TestOnDemand" {
  id = "something_cool"
  retries = 5
  cpus = 0.5
  memory = 1024
  image = "arbornetworks-docker-docker.bintray.io/aws-cli_0.2.0:18da34d"
}
