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
    },
  }
}

func resourceKafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*KafkaManagingClient)

  topicName  := d.Get("name").(string)
  partitions := d.Get("partitions").(int)
  replicas   := d.Get("replication_factor").(int)

  log.Printf("[DEBUG] Kafka to create topic '%s'", topicName)

  d.SetId(topicName)

  res, err := client.createTopic(topicName, partitions, replicas)

  if (err == nil) {
    log.Printf("[DEBUG] Kafka topic '%s:%d:%d' created ", topicName, res.PartitionsCount, res.ReplicationFactor)
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

  return nil
}


func resourceKafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Deleting Kafka On Demand '%s'", d.Id())
  
  // perform delete

  return nil
}
