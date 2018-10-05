[![Build status](https://circleci.com/gh/packetloop/terraform-provider-kafka.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/packetloop/terraform-provider-kafka)
[![GitHub release](https://img.shields.io/github/release/packetloop/terraform-provider-kafka.svg)](https://github.com/packetloop/terraform-provider-kafka/releases/)
[![All Contributors](https://img.shields.io/github/contributors/packetloop/terraform-provider-kafka.svg?longCache=true&style=flat-square&colorB=orange&label=all%20contributors)](#contributors)
[![Github All Releases](https://img.shields.io/github/downloads/packetloop/terraform-provider-kafka/total.svg)]()


# terraform-provider-kafka

A Terraform provider to manage Kafka topics lifecycle and compatible with latest
Terraform version. 

We couldn't use Confluent's Kafka REST service due to the way we create a Kafka topic.
Hence, we build our own REST service https://github.com/packetloop/kafka-admin-service
with Go bindings https://github.com/packetloop/go-kafkaesque.


#### [NOTE]

There's a old version of this provider in develop branch which is not compatible with
latest Terraform.

## Usage:

Download this provider, pick a version you'd like from releases from
[Binary Releases](https://github.com/packetloop/terraform-provider-kafka/releases)

```bash
curl -L \
  https://github.com/packetloop/terraform-provider-kafka/releases/download/v0.1.0/terraform-provider-kafka_v0.1.0_darwin_x86_64 \
  -o ~/.terraform.d/plugins/terraform-provider-kafka_v0.1.0 && \
  chmod +x ~/.terraform.d/plugins/terraform-provider-kafka_v0.1.0
```

```bash
provider "kafka" {
  host_url     = "http://localhost:8080"
  version = "~> 0.1.0"
}

resource "kafka_topic" "my-topic" {
  name = "my-topic"
  partitions = 2
  replication_factor = 1
  retention_ms = 300000
  cleanup_policy = "compact"
  segment_ms = 1440000
  segment_bytes = 1073741824
}
```

## Development and testing:

1. Clone kafka-admin-service ( REST SERVICE )

```bash
  git clone git@github.com:packetloop/kafka-admin-service.git && \
    cd kafka-admin-service && ./run.sh

    TF_ACC=1 make test
```

2. From this project:

```
  TF_ACC=0  go test -race -cover -v ./...
```

More examples can be found in examples/main.tf.

## TODO:

### Import Resources:

Syntax

```
  terraform import kafka_request.lenfree-run <resource ID>
  terraform import kafka_docker_deploy.test-deploy-2 <resource ID>
```

### Data resource:
