package provider

import (
	client "terraform-provider-python-app/api"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ENDPOINT", "http://127.0.0.1:5000"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"numbers": resourceList(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	return client.NewClient(endpoint), nil
}
