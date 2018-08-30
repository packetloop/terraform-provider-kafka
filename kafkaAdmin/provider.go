package kafkaadmin

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider function returns a custom TF provider called kafkaadmin that manages Kafka
// topics lifecycle.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOST_URL", nil),
				Description: "Kafka Admin Service API endpoint to manage Kafka topic lifecycle.",
			},
			"retry": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("retry", 3),
				Description: "Number of times to retry when http requests fails. Defaults to 3.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kafka_topic": resourceKafkaTopic(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		HostURL: d.Get("host_url").(string),
		Retry:   d.Get("retry").(int),
	}

	return config.Client()
}
