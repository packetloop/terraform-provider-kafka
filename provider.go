package main

import (
  "github.com/hashicorp/terraform/helper/schema"
  "github.com/hashicorp/terraform/terraform"
  "os"
  "fmt"
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
  topicScript  := "kafka-topics.sh"
  configScript := "kafka-configs.sh"
  if v := d.Get("kafka_bin_path").(string); v != "" {
    topicScript = v + "/" + topicScript
    configScript = v + "/" + configScript
  }

  if err := ensureScriptExists(topicScript);  err != nil { return nil, err }
  if err := ensureScriptExists(configScript); err != nil { return nil, err }

  client.TopicScript  = topicScript
  client.ConfigScript = configScript
  return client, nil
}

func ensureScriptExists(path string) error {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return fmt.Errorf("Unable to find Kafka scripts: %s not found", path)
  }
  return nil
}
