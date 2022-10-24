package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	defenderclient "github.com/liquid-collective/terraform-provider-openzeppelin-defender/pkg/defender/client/http"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(defenderclient.APIKeyEnv, nil),
			},
			"api_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(defenderclient.APISecretEnv, nil),
			},
			"user_pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(defenderclient.UserPoolIDEnv, defenderclient.UserPoolIDDefault),
			},
			"client_pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(defenderclient.ClientPoolIDEnv, defenderclient.ClientPoolIDDefault),
			},
			"api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(defenderclient.APIURLEnv, defenderclient.APIURLDefault),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"defender_proposal": resourceProposal(),
		},
	}

	provider.ConfigureContextFunc = ConfigureProvider()

	return provider
}

func ConfigureProvider() schema.ConfigureContextFunc {
	return func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		cfg := &defenderclient.Config{
			APIURL:       data.Get("api_url").(string),
			APIKey:       data.Get("api_key").(string),
			APISecret:    data.Get("api_secret").(string),
			UserPoolID:   data.Get("user_pool_id").(string),
			ClientPoolID: data.Get("client_pool_id").(string),
		}

		client, err := defenderclient.New(cfg.SetDefault())
		if err != nil {
			return nil, diag.FromErr(err)
		}

		err = client.Init(ctx)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return client, nil
	}
}
