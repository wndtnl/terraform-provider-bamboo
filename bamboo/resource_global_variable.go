package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceGlobalVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlobalVariableCreate,
		ReadContext:   resourceGlobalVariableRead,
		UpdateContext: resourceGlobalVariableUpdate,
		DeleteContext: resourceGlobalVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGlobalVariableImportState,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

func unmarshalGlobalVariable(s *schema.ResourceData) *bamboo.GlobalVariable {

	v := new(bamboo.GlobalVariable)

	v.Id = s.Get("id").(string)
	v.Key = s.Get("key").(string)
	v.Value = s.Get("value").(string)

	return v
}

func marshalGlobalVariable(s *schema.ResourceData, v *bamboo.GlobalVariable) error {

	if err := s.Set("id", v.Id); err != nil {
		return err
	}
	if err := s.Set("key", v.Key); err != nil {
		return err
	}
	if err := s.Set("value", v.Value); err != nil {
		return err
	}

	return nil
}

func resourceGlobalVariableCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	variable := unmarshalGlobalVariable(data)

	newVariable, err := client.GlobalVariable.Create(variable)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newVariable.Id)

	return resourceGlobalVariableRead(ctx, data, meta)
}

func resourceGlobalVariableRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	variable, err := client.GlobalVariable.GetOne(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = marshalGlobalVariable(data, variable); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGlobalVariableUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	variable := unmarshalGlobalVariable(data)

	err := client.GlobalVariable.Update(id, variable)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGlobalVariableRead(ctx, data, meta)
}

func resourceGlobalVariableDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.GlobalVariable.Delete(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags
}

func resourceGlobalVariableImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	key := data.Id()

	client := meta.(*bamboo.Client)

	variable, err := client.GlobalVariable.Search(key)
	if err != nil {
		return nil, err
	}

	if err = marshalGlobalVariable(data, variable); err != nil {
		return nil, err
	}

	data.SetId(variable.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
