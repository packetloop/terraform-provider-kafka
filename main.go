package main

import (
	kafkaadmin "github.com/comozo/terraform-provider-kafka/kafkaAdmin"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kafkaadmin.Provider})
}
