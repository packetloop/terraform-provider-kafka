# Terraform-Kafka Provider

This is a plugin for HashiCorp [Terraform](https://terraform.io/), which helps creates, configures and deletes topics on  on [Kafka](http://kafka.apache.org/).

## Usage

- Download the plugin from [Releases](https://github.com/packetloop/terraform-provider-kafka/releases) page.
- [Install](https://terraform.io/docs/plugins/basics.html) it, or put into a directory with configuration files.
- Create a minimal terraform template file.  There is an example in `sample/sample.tf`.
- Modify zookeeper settings in the provider and the topic settings in the `kafka_topic` resource.
- Run:
```
$ terraform apply
```

## `kafka` Provider Parameters

### Mandatory Parameters
- `kafka.zookeeper` - address to a node in the zookeeper cluster in `hostname[:port]` format

### Optional Parameters
- `kafka.kafka_bin_path` - specify the path to the Kafka command line tools if they are not on your path

## `kafka_topic` Resource Parameters

### Mandatory Parameters
- `kafka_topic.name` - name of the topic

### Optional Parameters
- `partitions` - number of partitions for the topic
- `replication_factor` - the replication factor for the topic
- `retention_bytes` - the retention bytes for the topic
- `retention_ms` - the retention period in milliseconds for the topic
- `cleanup_policy` - the clean up policy for the topic, for example compaction
- `segment_bytes` - the segment file size for the log
- `segement_ms` - the time after which Kafka will force the log to roll

## Building

This project uses the [glide](https://github.com/Masterminds/glide) package manager.
The package manger allows control of versions of dependencies used, including terraform.

Glide can be installed with the homebrew package manager.

The project dependencies can be installed with

```
glide install
```

## Manual testing in docker-compose

### Prerequisites

This requires that `docker-machine` is installed.  On OS X, this can be done with:

```
$ brew install docker-machine
$ brew install docker-compose
```

### Test steps

In one terminal run:

```
1$ cd sample/
1$ docker-compose up
```

Then in another terminal run:

```
2$ cd sample/
2$ docker exec -ti sample_kafka_1 /usr/bin/kafka-topics --list --zookeeper 192.168.99.100:2181
```

There should be no topics

```
2$ terraform apply
2$ docker exec -ti sample_kafka_1 /usr/bin/kafka-topics --list --zookeeper 192.168.99.100:2181
```

There should be a new topic `my-topic`.
