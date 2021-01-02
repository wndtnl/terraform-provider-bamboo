package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceProjectImportState,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func unmarshalProject(s *schema.ResourceData) *bamboo.Project {

	p := new(bamboo.Project)

	p.Id = s.Get("id").(string)
	p.Name = s.Get("name").(string)
	p.Key = s.Get("key").(string)
	p.Description = s.Get("description").(string)

	return p
}

func marshalProject(s *schema.ResourceData, p *bamboo.Project) error {

	if err := s.Set("id", p.Id); err != nil {
		return err
	}

	if err := s.Set("name", p.Name); err != nil {
		return err
	}

	if err := s.Set("key", p.Key); err != nil {
		return err
	}

	if err := s.Set("description", p.Description); err != nil {
		return err
	}

	return nil
}

func resourceProjectCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	project := unmarshalProject(data)

	newProject, err := client.Project.Create(project)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newProject.Id)

	return resourceProjectRead(ctx, data, meta)
}

func resourceProjectRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	project, err := client.Project.GetOne(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalProject(data, project); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProjectUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	project := unmarshalProject(data)

	err := client.Project.Update(id, project)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProjectRead(ctx, data, meta)
}

func resourceProjectDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.Project.Delete(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceProjectImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	key := data.Id()

	client := meta.(*bamboo.Client)

	project, err := client.Project.Search(key)
	if err != nil {
		return nil, err
	}

	if err = marshalProject(data, project); err != nil {
		return nil, err
	}

	data.SetId(project.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
