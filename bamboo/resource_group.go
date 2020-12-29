package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGroupCreate,
		ReadContext:   resourceGroupRead,
		UpdateContext: resourceGroupUpdate,
		DeleteContext: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGroupImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"members": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func unmarshalGroup(s *schema.ResourceData) *bamboo.Group {

	g := new(bamboo.Group)

	g.Name = s.Get("name").(string)
	rawMembers := s.Get("members").(*schema.Set).List()
	g.Members = *ExpandStringSlice(rawMembers)

	return g
}

func marshalGroup(s *schema.ResourceData, g *bamboo.Group) error {

	if err := s.Set("name", g.Name); err != nil {
		return err
	}

	memberSet := schema.NewSet(schema.HashString, FlattenStringSlice(&g.Members))
	if err := s.Set("members", memberSet); err != nil {
		return err
	}

	return nil
}

func resourceGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	group := unmarshalGroup(data)

	newGroup, err := client.Group.Create(group)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newGroup.Name)

	return resourceGroupRead(ctx, data, meta)
}

func resourceGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	name := data.Id()
	group, err := client.Group.GetOne(name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = marshalGroup(data, group); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGroupUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	name := data.Id()
	group := unmarshalGroup(data)

	err := client.Group.Update(name, group)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGroupRead(ctx, data, meta)
}

func resourceGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	name := data.Id()

	err := client.Group.Delete(name)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags
}

func resourceGroupImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := data.Id()

	client := meta.(*bamboo.Client)

	group, err := client.Group.GetOne(name)
	if err != nil {
		return nil, err
	}

	if err = marshalGroup(data, group); err != nil {
		return nil, err
	}

	data.SetId(name)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
