package main

import (
	"github.com/hashicorp/terraform/plugin"
)

const providerName = "kafka"
const resourceTypeName = "kafka job"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
