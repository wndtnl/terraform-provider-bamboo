package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BAMBOO_ADDR", bamboo.DefaultBaseURL),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BAMBOO_USER", bamboo.DefaultUsername),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BAMBOO_PASS", bamboo.DefaultPassword),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bamboo_deployment_project": resourceDeploymentProject(),
			"bamboo_global_permission":  resourceGlobalPermission(),
			"bamboo_global_variable":    resourceGlobalVariable(),
			"bamboo_group":              resourceGroup(),
			"bamboo_linked_repository":  resourceLinkedRepository(),
			"bamboo_local_agent":        resourceLocalAgent(),
			"bamboo_project":            resourceProject(),
			"bamboo_shared_credential":  resourceSharedCredential(),
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	address := data.Get("address").(string)
	username := data.Get("username").(string)
	password := data.Get("password").(string)

	client, err := bamboo.NewClient(address, username, password)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, diags
}
