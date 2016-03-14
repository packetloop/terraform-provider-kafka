package main

import (
  "log"
  "github.com/hashicorp/terraform/helper/schema"
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
      "replicas": &schema.Schema{
        Type:        schema.TypeFloat,
        Required:    true,
        Description: "replication factor",
      },
    },
  }
}

func resourceKafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {

  id := d.Get("name").(string)

  log.Printf("[DEBUG] Kafka On Demand create: %s", id)

  // create request

  // create deploy

  d.SetId(id)
      
  return resourceKafkaTopicRead(d, meta)
}

func resourceKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Updating Kafka On Demand '%s'", d.Id())
  
  // if request changed, recreate request

  // if deploy changed, do a new deploy

  return nil
}


func resourceKafkaTopicRead(d *schema.ResourceData, meta interface{}) error {
  log.Printf("[DEBUG] Loading data for Kafka On Demand '%s'", d.Id())

  return nil
}


func resourceKafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Deleting Kafka On Demand '%s'", d.Id())
  
  // perform delete

  return nil
}
