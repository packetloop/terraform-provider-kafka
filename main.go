package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	kafkaadmin "github.com/packetloop/terraform-provider-kafka/kafkaAdmin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kafkaadmin.Provider})
}
