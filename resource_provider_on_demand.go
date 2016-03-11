package main

import (
  "log"

  "github.com/hashicorp/terraform/helper/schema"
)

func resourceProviderOnDemand() *schema.Resource {
  return &schema.Resource{
    Create: resourceProviderOnDemandCreate,
    Read:   resourceProviderOnDemandRead,
    Update: resourceProviderOnDemandUpdate,
    Delete: resourceProviderOnDemandDelete,

    Schema: map[string]*schema.Schema{
      "id": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        ForceNew:    true,
        Description: "the id of the " + resourceTypeName,
      },
      "retries": &schema.Schema{
        Type:        schema.TypeInt,
        Optional:    true,
        Default:     3,
        Description: "retries on failure",
      },
      "cpus": &schema.Schema{
        Type:        schema.TypeFloat,
        Optional:    true,
        Default:     0.1,
        Description: "cpus to use",
      },
      "memory": &schema.Schema{
        Type:        schema.TypeInt,
        Optional:    true,
        Default:     512,
        Description: "memory in megabytes",
      },
      "image": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: "tag for the docker image",
      },
    },
  }
}

func resourceProviderOnDemandCreate(d *schema.ResourceData, meta interface{}) error {

  id := d.Get("id").(string)

  log.Printf("[DEBUG] Kafka On Demand create: %s", id)

  // create request

  // create deploy

  d.SetId(id)
      
  return resourceProviderOnDemandRead(d, meta)
}

func resourceProviderOnDemandUpdate(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Updating Kafka On Demand '%s'", d.Id())
  
  // if request changed, recreate request

  // if deploy changed, do a new deploy

  return nil
}


func resourceProviderOnDemandRead(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Loading data for Kafka On Demand '%s'", d.Id())

  // get remove state, save to schema

  return nil
}


func resourceProviderOnDemandDelete(d *schema.ResourceData, meta interface{}) error {
  
  log.Printf("[DEBUG] Deleting Kafka On Demand '%s'", d.Id())
  
  // perform delete

  return nil
}

