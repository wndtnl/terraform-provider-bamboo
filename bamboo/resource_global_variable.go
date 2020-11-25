package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"strconv"
)

func resourceGlobalVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlobalVariableCreate,
		ReadContext:   resourceGlobalVariableRead,
		UpdateContext: resourceGlobalVariableUpdate,
		DeleteContext: resourceGlobalVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGlobalVariableCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	key := data.Get("key").(string)
	value := data.Get("value").(string)

	variable, err := client.GlobalVariable.Create(key, value)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(strconv.FormatInt(variable.Id, 10))

	return resourceGlobalVariableRead(ctx, data, meta)
}

func resourceGlobalVariableRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	variableId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	variable, err := client.GlobalVariable.GetOne(variableId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("key", variable.Key); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("value", variable.Value); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGlobalVariableUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	variableId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	if data.HasChange("key") || data.HasChange("value") {

		err = client.GlobalVariable.Update(
			variableId, data.Get("key").(string), data.Get("value").(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceGlobalVariableRead(ctx, data, meta)
}

func resourceGlobalVariableDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	variableId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	err = client.GlobalVariable.Delete(variableId)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags
}