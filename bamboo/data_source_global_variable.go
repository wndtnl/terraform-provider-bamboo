package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"strconv"
	"time"
)

func dataSourceGlobalVariables() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGlobalVariablesRead,
		Schema: map[string]*schema.Schema{
			"global_variables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGlobalVariablesRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	// Fetch
	client := meta.(*bamboo.Client)
	variables, err := client.GlobalVariable.GetAll()
	if err != nil {
		return diag.FromErr(err)
	}

	// Store
	mv := mapGlobalVariables(variables)
	if err := data.Set("global_variables", mv); err != nil {
		return diag.FromErr(err)
	}

	// Always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func mapGlobalVariables(variables []*bamboo.GlobalVariable) []interface{} {
	if variables != nil {
		vs := make([]interface{}, len(variables), len(variables))
		for i, variable := range variables {
			vs[i] = mapGlobalVariable(*variable)
		}
		return vs
	}
	return make([]interface{}, 0)
}

func mapGlobalVariable(variable bamboo.GlobalVariable) interface{} {
	v := make(map[string]interface{})
	v["id"] = variable.Id
	v["key"] = variable.Key
	v["value"] = variable.Value
	return v
}
