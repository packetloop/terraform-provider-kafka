package main

import (
  "os/exec"
  "strings"
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
        Description: providerName + " Custom path to Kafka executables",
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
  client := new(KafkaManagingClient)
  prefixPath := d.Get("kafka_bin_path").(string)
  var err error

  client.TopicScript, err  = scriptPath(prefixPath, "kafka-topics", "kafka-topics.sh")
  if err != nil { return nil, err }

  client.ConfigScript, err = scriptPath(prefixPath, "kafka-configs", "kafka-configs.sh")
  if err != nil { return nil, err }

  client.Zookeeper = d.Get("zookeeper").(string)

  return client, nil
}

func ensureScriptExists(path string) error {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return fmt.Errorf("Unable to find Kafka scripts: %s not found", path)
  }
  return nil
}

func scriptPath(prefixPath string, scriptNames ... string) (string, error) {
  if prefixPath != "" {
    for _, scriptName := range scriptNames {
      scriptPath, err := which(prefixPath + "/" + scriptName)
      if err == nil { return scriptPath, nil }
    }
    
    return "", fmt.Errorf("None of these scripts exist %s in %s", scriptNames, prefixPath)
  }

  for _, scriptName := range scriptNames {
    scriptPath, err := which(scriptName)
    if err == nil { return scriptPath, nil }
  }

  return "", fmt.Errorf("None of these scripts exist %s", scriptNames)
}

func which(executables ... string) (string, error) {
  for _, executable := range executables {
    cmd := exec.Command("which", executable)
    
    buffer, err := cmd.Output(); if err != nil { continue }
    fullPath := strings.TrimSpace(fmt.Sprintf("%s", buffer)); if fullPath == "" { continue }
    
    return fullPath, nil
  }
  
  return "", fmt.Errorf("None of these executables exist %s", executables)
}
