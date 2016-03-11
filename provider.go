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
      "endpoint": &schema.Schema{
        Type:        schema.TypeString,
        Required:    true,
        Description: providerName + " endpoint",
      },
    },
    
    ResourcesMap: map[string]*schema.Resource{
      "kafka_on_demand": resourceProviderOnDemand(),
    },

    ConfigureFunc: providerConfigure,
  }
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  var client = new(ProviderClient)
  client.endpoint = d.Get("endpoint").(string)
  return client, nil
}
