package bamboo

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceUserImportState,
		},
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				StateFunc: func(value interface{}) string {
					// Prevent exposure of the password by storing a hash
					return hashString(value.(string))
				},
			},
			"jabber_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func unmarshalUser(s *schema.ResourceData) *bamboo.User {

	u := new(bamboo.User)

	u.Username = s.Get("username").(string)
	u.FullName = s.Get("full_name").(string)
	u.Email = s.Get("email").(string)
	u.Password = s.Get("password").(string)
	u.JabberAddress = s.Get("jabber_address").(string)
	u.Enabled = s.Get("active").(bool)

	return u
}

func marshalUser(s *schema.ResourceData, u *bamboo.User) error {

	if err := s.Set("username", u.Username); err != nil {
		return err
	}

	if err := s.Set("full_name", u.FullName); err != nil {
		return err
	}

	if err := s.Set("email", u.Email); err != nil {
		return err
	}

	// Do not set password, as it cannot be read

	if err := s.Set("jabber_address", u.JabberAddress); err != nil {
		return err
	}

	if err := s.Set("active", u.Enabled); err != nil {
		return err
	}

	return nil
}

func hashString(str string) string {
	hash := sha256.Sum256([]byte(str))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func resourceUserCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	user := unmarshalUser(data)

	newUser, err := client.User.Create(user)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newUser.Username)

	return resourceUserRead(ctx, data, meta)
}

func resourceUserRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	username := data.Id()

	user, err := client.User.GetOne(username)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalUser(data, user); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceUserUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	username := data.Id()
	user := unmarshalUser(data)

	err := client.User.Update(username, user)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceUserRead(ctx, data, meta)
}

func resourceUserDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	username := data.Id()

	err := client.User.Delete(username)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceUserImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	username := data.Id()

	client := meta.(*bamboo.Client)

	user, err := client.User.GetOne(username)
	if err != nil {
		return nil, err
	}

	if err = marshalUser(data, user); err != nil {
		return nil, err
	}

	data.SetId(username)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
