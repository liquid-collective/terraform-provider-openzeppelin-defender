package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	defenderprovider "github.com/liquid-collective/terraform-provider-openzeppelin-defender/src/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return defenderprovider.Provider()
		},
	})
}
