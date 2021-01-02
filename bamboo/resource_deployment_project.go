package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceDeploymentProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeploymentProjectCreate,
		ReadContext:   resourceDeploymentProjectRead,
		UpdateContext: resourceDeploymentProjectUpdate,
		DeleteContext: resourceDeploymentProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceDeploymentProjectImportState,
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
			"plan_key": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func unmarshalDeploymentProject(s *schema.ResourceData) *bamboo.DeploymentProject {

	p := new(bamboo.DeploymentProject)

	p.Id = s.Get("id").(string)
	p.Name = s.Get("name").(string)
	p.Description = s.Get("description").(string)
	p.PlanKey = s.Get("plan_key").(string)

	return p
}

func marshalDeploymentProject(s *schema.ResourceData, p *bamboo.DeploymentProject) error {

	if err := s.Set("id", p.Id); err != nil {
		return err
	}

	if err := s.Set("name", p.Name); err != nil {
		return err
	}

	if err := s.Set("description", p.Description); err != nil {
		return err
	}

	if err := s.Set("plan_key", p.PlanKey); err != nil {
		return err
	}

	return nil
}

func resourceDeploymentProjectCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	deploymentProject := unmarshalDeploymentProject(data)

	newDeploymentProject, err := client.DeploymentProject.Create(deploymentProject)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newDeploymentProject.Id)

	return resourceDeploymentProjectRead(ctx, data, meta)
}

func resourceDeploymentProjectRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	deploymentProject, err := client.DeploymentProject.GetOne(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalDeploymentProject(data, deploymentProject); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDeploymentProjectUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	deploymentProject := unmarshalDeploymentProject(data)

	err := client.DeploymentProject.Update(id, deploymentProject)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDeploymentProjectRead(ctx, data, meta)
}

func resourceDeploymentProjectDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.DeploymentProject.Delete(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceDeploymentProjectImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	id := data.Id()

	client := meta.(*bamboo.Client)

	deploymentProject, err := client.DeploymentProject.GetOne(id)
	if err != nil {
		return nil, err
	}

	if err = marshalDeploymentProject(data, deploymentProject); err != nil {
		return nil, err
	}

	data.SetId(deploymentProject.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
