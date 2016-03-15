package main

import (
  "log"
  "github.com/hashicorp/terraform/helper/schema"
  "fmt"
)

func resourceKafkaTopic() *schema.Resource {
  return &schema.Resource{
    Create: resourceKafkaTopicCreate,
    Read:   resourceKafkaTopicRead,
    Update: resourceKafkaTopicUpdate,
    Delete: resourceKafkaTopicDelete,

    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        ForceNew:    true,
        Description: "topic name",
      },
      "partitions": &schema.Schema{
        Type:        schema.TypeInt,
        Required:    true,
        Description: "number of partitions",
      },
      "replication_factor": &schema.Schema{
        Type:        schema.TypeInt,
        Required:    true,
        Description: "replication factor",
      },
      "retention_bytes": &schema.Schema{
        Type:        schema.TypeInt,
        Optional:    true,
        Description: "log.retention.bytes",
        Default:     -1,
      },
      "retention_ms": &schema.Schema{
        Type:        schema.TypeInt,
        Optional:    true,
        Description: "log.retention.ms",
        Default:     -1,
      },
      "cleanup_policy": &schema.Schema{
        Type:        schema.TypeString,
        Optional:    true,
        Description: "cleanup.policy",
        Default:     "",
      },
    },
  }
}

func resourceKafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*KafkaManagingClient)

  topicName  := d.Get("name").(string)

  log.Printf("[DEBUG] Kafka to create topic '%s'", topicName)

  conf := buildKafkaConfig(d)

  d.SetId(topicName)

  err := client.createTopic(topicName, conf)

  if (err == nil) {
    log.Printf("[DEBUG] Kafka topic '%s:%d:%d' created ", topicName, conf.PartitionsCount, conf.ReplicationFactor)
  } else {
    log.Printf("[DEBUG] Kafka - unable to create topic: %v", err)
  }

  return err
}

func resourceKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Updating Kafka On Demand '%s'", d.Id())
  
  // if request changed, recreate request

  // if deploy changed, do a new deploy

  return nil
}


func resourceKafkaTopicRead(d *schema.ResourceData, meta interface{}) error {
  topicName := d.Get("name").(string)
  log.Printf("[DEBUG] Loading data for Kafka topic '%s' ['%s']", topicName, d.Id())

  client := meta.(*KafkaManagingClient)
  info, err := client.describeTopic(topicName)

  if (err != nil) {
    return fmt.Errorf("Error while looking for a topic '%s'", topicName)
  }

  if (!info.exists()) {
    d.SetId("")
    return nil
  }

  d.Set("name", topicName)
  d.Set("partitions", info.PartitionsCount)
  d.Set("replication_factor", info.ReplicationFactor)
  d.Set("cleanup_policy", info.CleanupPolicy)
  d.Set("retention_bytes", info.RetentionBytes)
  d.Set("retention_ms", info.RetentionMs)

  return nil
}


func resourceKafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {
  topicName := d.Get("name").(string)
  log.Printf("[DEBUG] Kafka to delete topic '%s' [%s]", topicName, d.Id())

  client := meta.(*KafkaManagingClient)

  return client.deleteTopic(topicName)
}

func buildKafkaConfig(d *schema.ResourceData) *KafkaTopicInfo {
  return &KafkaTopicInfo{
    PartitionsCount:   d.Get("partitions").(int),
    ReplicationFactor: d.Get("replication_factor").(int),
    CleanupPolicy:     d.Get("cleanup_policy").(string),
    RetentionBytes:    int64(d.Get("retention_bytes").(int)),
    RetentionMs:       int64(d.Get("retention_ms").(int)),
  }
}
