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
