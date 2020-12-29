package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"strings"
)

func resourceSharedCredential() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSharedCredentialCreate,
		ReadContext:   resourceSharedCredentialRead,
		UpdateContext: resourceSharedCredentialUpdate,
		DeleteContext: resourceSharedCredentialDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceSharedCredentialImportState,
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
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
					"PASSWORD",
					"SSH",
					"AWS",
				}, false)),
			},
			"username": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"ssh_key", "ssh_passphrase", "access_key", "secret_key"},
			},
			"password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"ssh_key", "ssh_passphrase", "access_key", "secret_key"},
			},
			"ssh_key": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"username", "password", "access_key", "secret_key"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if strings.TrimSpace(old) == strings.TrimSpace(new) {
						return true
					}
					return false
				},
			},
			"ssh_passphrase": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"username", "password", "access_key", "secret_key"},
			},
			"access_key": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"username", "password", "ssh_key", "ssh_passphrase"},
			},
			"secret_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Sensitive:        true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"username", "password", "ssh_key", "ssh_passphrase"},
			},
		},
	}
}

func unmarshalSharedCredential(s *schema.ResourceData) *bamboo.SharedCredential {

	c := new(bamboo.SharedCredential)

	c.Id = s.Get("id").(string)
	c.Name = s.Get("name").(string)
	c.Type = s.Get("type").(string)

	switch c.Type {
	case "PASSWORD":
		c.Username = s.Get("username").(string)
		c.Password = s.Get("password").(string)
	case "SSH":
		c.SshKey = s.Get("ssh_key").(string)
		c.SshPassphrase = s.Get("ssh_passphrase").(string)
	case "AWS":
		c.AccessKey = s.Get("access_key").(string)
		c.SecretKey = s.Get("secret_key").(string)
	}

	return c
}

func marshalSharedCredentials(s *schema.ResourceData, c *bamboo.SharedCredential) error {

	if err := s.Set("id", c.Id); err != nil {
		return err
	}

	if err := s.Set("name", c.Name); err != nil {
		return err
	}

	if err := s.Set("type", c.Type); err != nil {
		return err
	}

	switch c.Type {
	case "PASSWORD":
		if err := s.Set("username", c.Username); err != nil {
			return err
		}
		if err := s.Set("password", c.Password); err != nil {
			return err
		}
	case "SSH":
		if err := s.Set("ssh_key", c.SshKey); err != nil {
			return err
		}
		if err := s.Set("ssh_passphrase", c.SshPassphrase); err != nil {
			return err
		}
	case "AWS":
		if err := s.Set("access_key", c.AccessKey); err != nil {
			return err
		}
		if err := s.Set("secret_key", c.SecretKey); err != nil {
			return err
		}
	}

	return nil
}

func resourceSharedCredentialCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	sharedCredential := unmarshalSharedCredential(data)

	newSharedCredential, err := client.SharedCredential.Create(sharedCredential)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newSharedCredential.Id)

	return resourceSharedCredentialRead(ctx, data, meta)
}

func resourceSharedCredentialRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	sharedCredential, err := client.SharedCredential.GetOne(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = marshalSharedCredentials(data, sharedCredential); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSharedCredentialUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	sharedCredential := unmarshalSharedCredential(data)

	err := client.SharedCredential.Update(id, sharedCredential)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSharedCredentialRead(ctx, data, meta)
}

func resourceSharedCredentialDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.SharedCredential.Delete(id)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags
}

func resourceSharedCredentialImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	key := data.Id()

	client := meta.(*bamboo.Client)

	sharedCredential, err := client.SharedCredential.Search(key)
	if err != nil {
		return nil, err
	}

	if err = marshalSharedCredentials(data, sharedCredential); err != nil {
		return nil, err
	}

	data.SetId(sharedCredential.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
