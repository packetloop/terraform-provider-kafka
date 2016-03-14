package main

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
  // The actual provider
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "kafka_bin_path": &schema.Schema{
        Type:        schema.TypeString,
        Required:    false,
        Optional:    true,
        Default:     "",
        Description: providerName + " Custom path to /bin/kafka-topics.sh",
      },
      "zookeeper": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: providerName + " Zookeeper address (<host>:[<port>])",
      },
    },
    
    ResourcesMap: map[string]*schema.Resource{
      "kafka_topic": resourceKafkaTopic(),
    },

    ConfigureFunc: providerConfigure,
  }
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var client = new(KafkaManagingClient)
  client.Zookeeper = d.Get("zookeeper").(string)
  script := "kafka-topics.sh"
  if v := d.Get("kafka_bin_path").(string); v != "" {
    script = v + "/" + script
  }
  client.TopicScript = script
  return client, nil
}
