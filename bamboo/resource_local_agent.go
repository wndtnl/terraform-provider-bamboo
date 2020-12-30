package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceLocalAgent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLocalAgentCreate,
		ReadContext:   resourceLocalAgentRead,
		UpdateContext: resourceLocalAgentUpdate,
		DeleteContext: resourceLocalAgentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceLocalAgentImportState,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func unmarshalLocalAgent(s *schema.ResourceData) *bamboo.LocalAgent {

	a := new(bamboo.LocalAgent)

	a.Id = s.Get("id").(string)
	a.Name = s.Get("name").(string)
	a.Description = s.Get("description").(string)
	a.Enabled = s.Get("enabled").(bool)

	return a
}

func marshalLocalAgent(s *schema.ResourceData, a *bamboo.LocalAgent) error {

	if err := s.Set("id", a.Id); err != nil {
		return err
	}

	if err := s.Set("name", a.Name); err != nil {
		return err
	}

	if err := s.Set("description", a.Description); err != nil {
		return err
	}

	if err := s.Set("enabled", a.Enabled); err != nil {
		return err
	}

	return nil
}

func resourceLocalAgentCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	localAgent := unmarshalLocalAgent(data)

	newLocalAgent, err := client.LocalAgent.Create(localAgent)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newLocalAgent.Id)

	return resourceLocalAgentRead(ctx, data, meta)
}

func resourceLocalAgentRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	localAgent, err := client.LocalAgent.GetOne(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalLocalAgent(data, localAgent); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceLocalAgentUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	localAgent := unmarshalLocalAgent(data)

	err := client.LocalAgent.Update(id, localAgent)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceLocalAgentRead(ctx, data, meta)
}

func resourceLocalAgentDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.LocalAgent.Delete(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceLocalAgentImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := data.Id()

	client := meta.(*bamboo.Client)

	localAgent, err := client.LocalAgent.Search(name)
	if err != nil {
		return nil, err
	}

	if err = marshalLocalAgent(data, localAgent); err != nil {
		return nil, err
	}

	data.SetId(localAgent.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
