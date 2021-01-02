package bamboo

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"strings"
)

func resourceLinkedRepositoryPermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLinkedRepositoryPermissionCreate,
		ReadContext:   resourceLinkedRepositoryPermissionRead,
		UpdateContext: resourceLinkedRepositoryPermissionUpdate,
		DeleteContext: resourceLinkedRepositoryPermissionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceLinkedRepositoryPermissionImportState,
		},
		Schema: map[string]*schema.Schema{
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
					"USER",
					"GROUP",
				}, false)),
			},
			"permissions": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
						"READ",
						"ADMINISTRATION",
					}, false)),
				},
				Set: schema.HashString,
			},
		},
	}
}

func unmarshalLinkedRepositoryPermission(s *schema.ResourceData) *bamboo.LinkedRepositoryPermission {

	r := new(bamboo.LinkedRepositoryPermission)

	r.RepositoryId = s.Get("repository_id").(string)
	r.Name = s.Get("name").(string)
	r.Type = s.Get("type").(string)
	rawPermissions := s.Get("permissions").(*schema.Set).List()
	r.Permissions = *expandStringSlice(rawPermissions)

	return r
}

func marshalLinkedRepositoryPermission(s *schema.ResourceData, r *bamboo.LinkedRepositoryPermission) error {

	if err := s.Set("repository_id", r.RepositoryId); err != nil {
		return err
	}

	if err := s.Set("name", r.Name); err != nil {
		return err
	}

	if err := s.Set("type", r.Type); err != nil {
		return err
	}

	permissionSet := schema.NewSet(schema.HashString, flattenStringSlice(&r.Permissions))
	if err := s.Set("permissions", permissionSet); err != nil {
		return err
	}

	return nil
}

func resourceLinkedRepositoryPermissionCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	linkedRepositoryPermission := unmarshalLinkedRepositoryPermission(data)

	newLinkedRepositoryPermission, err := client.LinkedRepositoryPermission.Upsert(linkedRepositoryPermission)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(getLinkedRepositoryPermissionId(newLinkedRepositoryPermission))

	return resourceLinkedRepositoryPermissionRead(ctx, data, meta)
}

func resourceLinkedRepositoryPermissionRead(_ context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	repositoryId := data.Get("repository_id").(string)
	permissionType := data.Get("type").(string)
	name := data.Get("name").(string)

	linkedRepositoryPermission, err := client.LinkedRepositoryPermission.GetOne(repositoryId, permissionType, name)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalLinkedRepositoryPermission(data, linkedRepositoryPermission); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceLinkedRepositoryPermissionUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	linkedRepositoryPermission := unmarshalLinkedRepositoryPermission(data)

	_, err := client.LinkedRepositoryPermission.Upsert(linkedRepositoryPermission)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceLinkedRepositoryPermissionRead(ctx, data, meta)
}

func resourceLinkedRepositoryPermissionDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	repositoryId := data.Get("repository_id").(string)
	permissionType := data.Get("type").(string)
	name := data.Get("name").(string)

	err := client.LinkedRepositoryPermission.Delete(repositoryId, permissionType, name)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceLinkedRepositoryPermissionImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	client := meta.(*bamboo.Client)

	id := data.Id()
	repositoryId, permissionType, name, err := parseLinkedRepositoryPermissionId(id)
	if err != nil {
		return nil, err
	}

	linkedRepositoryPermission, err := client.LinkedRepositoryPermission.GetOne(repositoryId, permissionType, name)
	if err != nil {
		return nil, err
	}

	if err = marshalLinkedRepositoryPermission(data, linkedRepositoryPermission); err != nil {
		return nil, err
	}

	data.SetId(getLinkedRepositoryPermissionId(linkedRepositoryPermission))

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}

func getLinkedRepositoryPermissionId(r *bamboo.LinkedRepositoryPermission) string {
	// Combination of Repository, Type and Name uniquely define the permission
	return fmt.Sprintf("%s|%s|%s", r.RepositoryId, r.Type, r.Name)
}

func parseLinkedRepositoryPermissionId(id string) (string, string, string, error) {

	s := strings.Split(id, "|")

	if len(s) < 3 {
		return "", "", "", errors.New("invalid import id format")
	}

	if !(s[1] == "USER" || s[1] == "GROUP") {
		return "", "", "", errors.New("invalid import id format")
	}

	repositoryId := s[0]
	permissionType := s[1]
	name := strings.Join(s[2:], "")

	return repositoryId, permissionType, name, nil
}
