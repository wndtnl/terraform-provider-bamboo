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

func resourceGlobalPermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGlobalPermissionCreate,
		ReadContext:   resourceGlobalPermissionRead,
		UpdateContext: resourceGlobalPermissionUpdate,
		DeleteContext: resourceGlobalPermissionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGlobalPermissionImportState,
		},
		Schema: map[string]*schema.Schema{
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
						"CREATE",
						"CREATEREPOSITORY",
						"ADMINISTRATION",
					}, false)),
				},
				Set: schema.HashString,
			},
		},
	}
}

func unmarshalGlobalPermission(s *schema.ResourceData) *bamboo.GlobalPermission {

	g := new(bamboo.GlobalPermission)

	g.Name = s.Get("name").(string)
	g.Type = s.Get("type").(string)
	rawPermissions := s.Get("permissions").(*schema.Set).List()
	g.Permissions = *expandStringSlice(rawPermissions)

	return g
}

func marshalGlobalPermission(s *schema.ResourceData, g *bamboo.GlobalPermission) error {

	if err := s.Set("name", g.Name); err != nil {
		return err
	}

	if err := s.Set("type", g.Type); err != nil {
		return err
	}

	permissionSet := schema.NewSet(schema.HashString, flattenStringSlice(&g.Permissions))
	if err := s.Set("permissions", permissionSet); err != nil {
		return err
	}

	return nil
}

func resourceGlobalPermissionCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	globalPermission := unmarshalGlobalPermission(data)

	newGlobalPermission, err := client.GlobalPermission.Upsert(globalPermission)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(getGlobalPermissionId(newGlobalPermission))

	return resourceGlobalPermissionRead(ctx, data, meta)
}

func resourceGlobalPermissionRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	permissionType := data.Get("type").(string)
	name := data.Get("name").(string)

	globalPermission, err := client.GlobalPermission.GetOne(permissionType, name)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalGlobalPermission(data, globalPermission); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGlobalPermissionUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	globalPermission := unmarshalGlobalPermission(data)

	_, err := client.GlobalPermission.Upsert(globalPermission)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGlobalPermissionRead(ctx, data, meta)
}

func resourceGlobalPermissionDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	permissionType := data.Get("type").(string)
	name := data.Get("name").(string)

	err := client.GlobalPermission.Delete(permissionType, name)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceGlobalPermissionImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	client := meta.(*bamboo.Client)

	id := data.Id()
	permissionType, name, err := parseGlobalPermissionId(id)
	if err != nil {
		return nil, err
	}

	globalPermission, err := client.GlobalPermission.GetOne(permissionType, name)
	if err != nil {
		return nil, err
	}

	if err = marshalGlobalPermission(data, globalPermission); err != nil {
		return nil, err
	}

	data.SetId(getGlobalPermissionId(globalPermission))

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}

func getGlobalPermissionId(g *bamboo.GlobalPermission) string {
	// Combination of Type and Name uniquely define the permission
	return fmt.Sprintf("%s|%s", g.Type, g.Name)
}

func parseGlobalPermissionId(id string) (string, string, error) {

	s := strings.Split(id, "|")

	if len(s) < 2 {
		return "", "", errors.New("invalid import id format")
	}

	if !(s[0] == "USER" || s[0] == "GROUP") {
		return "", "", errors.New("invalid import id format")
	}

	permissionType := s[0]
	name := strings.Join(s[1:], "")

	return permissionType, name, nil
}
