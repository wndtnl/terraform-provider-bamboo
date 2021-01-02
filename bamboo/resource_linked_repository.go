package bamboo

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"strings"
)

func resourceLinkedRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLinkedRepositoryCreate,
		ReadContext:   resourceLinkedRepositoryRead,
		UpdateContext: resourceLinkedRepositoryUpdate,
		DeleteContext: resourceLinkedRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceLinkedRepositoryImportState,
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
				ForceNew: true,
				ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
					"GIT",
				}, false)),
			},
			"repository_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "master",
			},
			"auth_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
					"NONE",
					"PASSWORD",
					"SSH",
					"PASSWORD_SHARED",
					"SSH_SHARED",
				}, false)),
			},
			"username": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"ssh_key", "ssh_passphrase", "shared_credential_id"},
			},
			"password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"ssh_key", "ssh_passphrase", "shared_credential_id"},
			},
			"ssh_key": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateV1(validation.StringIsNotEmpty),
				ConflictsWith:    []string{"username", "password", "shared_credential_id"},
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
				ConflictsWith: []string{"username", "password", "shared_credential_id"},
			},
			"shared_credential_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"username", "password", "ssh_key", "ssh_passphrase"},
			},
			"shallow_clones": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// There is a bug in Bamboo, at least since 6.8.0 and still present
			// in the latest (7.2.1) versions which prevents changing the default value
			// of the remote agent cache property at creation time!
			"remote_agent_cache": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"submodules": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"verbose_logs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fetch_whole_repository": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"lfs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// There is a bug in Bamboo, at least since 6.8.0 and still present
			// in the latest (7.2.1) versions which prevents changing the default value
			// of the command timeout property at creation time!
			"command_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  180,
			},
			"quit_period": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"quiet_period_wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"quiet_period_max_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"filter_pattern": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validateV1(validation.StringInSlice([]string{
					"NONE",
					"INCLUDE_ONLY",
					"EXCLUDE_ALL",
				}, false)),
				Default: "NONE",
			},
			"filter_pattern_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"change_set_regex": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func unmarshalLinkedRepository(s *schema.ResourceData) *bamboo.LinkedRepository {

	r := new(bamboo.LinkedRepository)

	r.Id = s.Get("id").(string)
	r.Name = s.Get("name").(string)
	r.Type = s.Get("type").(string)

	r.RepositoryUrl = s.Get("repository_url").(string)
	r.Branch = s.Get("branch").(string)

	r.AuthType = s.Get("auth_type").(string)

	switch r.AuthType {
	case "PASSWORD":
		r.Username = s.Get("username").(string)
		r.Password = s.Get("password").(string)
	case "SSH":
		r.SshKey = s.Get("ssh_key").(string)
		r.SshPassphrase = s.Get("ssh_passphrase").(string)
	case "PASSWORD_SHARED", "SSH_SHARED":
		r.SharedCredentialId = s.Get("shared_credential_id").(string)
	}

	r.UseShallowClones = s.Get("shallow_clones").(bool)
	r.UseRemoteAgentCache = s.Get("remote_agent_cache").(bool)
	r.UseSubmodules = s.Get("submodules").(bool)
	r.VerboseLogs = s.Get("verbose_logs").(bool)
	r.FetchWholeRepository = s.Get("fetch_whole_repository").(bool)
	r.UseLFS = s.Get("lfs").(bool)

	r.CommandTimeout = s.Get("command_timeout").(int)

	r.QuietPeriodEnabled = s.Get("quit_period").(bool)
	r.QuietPeriodWaitTime = s.Get("quiet_period_wait_time").(int)
	r.QuietPeriodMaxRetries = s.Get("quiet_period_max_retries").(int)

	r.FilterPattern = s.Get("filter_pattern").(string)
	r.FilterPatternRegex = s.Get("filter_pattern_regex").(string)
	r.ChangeSetRegex = s.Get("change_set_regex").(string)

	return r
}

func marshalLinkedRepository(s *schema.ResourceData, r *bamboo.LinkedRepository) error {

	if err := s.Set("id", r.Id); err != nil {
		return err
	}

	if err := s.Set("name", r.Name); err != nil {
		return err
	}

	if err := s.Set("type", r.Type); err != nil {
		return err
	}

	if err := s.Set("repository_url", r.RepositoryUrl); err != nil {
		return err
	}

	if err := s.Set("branch", r.Branch); err != nil {
		return err
	}

	if err := s.Set("auth_type", r.AuthType); err != nil {
		return err
	}

	switch r.AuthType {
	case "PASSWORD":
		if err := s.Set("username", r.Username); err != nil {
			return err
		}
		if err := s.Set("password", r.Password); err != nil {
			return err
		}
	case "SSH":
		if err := s.Set("ssh_key", r.SshKey); err != nil {
			return err
		}
		if err := s.Set("ssh_passphrase", r.SshPassphrase); err != nil {
			return err
		}
	case "PASSWORD_SHARED", "SSH_SHARED":
		if err := s.Set("shared_credential_id", r.SharedCredentialId); err != nil {
			return err
		}
	}

	if err := s.Set("shallow_clones", r.UseShallowClones); err != nil {
		return err
	}

	if err := s.Set("remote_agent_cache", r.UseRemoteAgentCache); err != nil {
		return err
	}

	if err := s.Set("submodules", r.UseSubmodules); err != nil {
		return err
	}

	if err := s.Set("verbose_logs", r.VerboseLogs); err != nil {
		return err
	}

	if err := s.Set("fetch_whole_repository", r.FetchWholeRepository); err != nil {
		return err
	}

	if err := s.Set("lfs", r.UseLFS); err != nil {
		return err
	}

	if err := s.Set("command_timeout", r.CommandTimeout); err != nil {
		return err
	}

	if err := s.Set("quit_period", r.QuietPeriodEnabled); err != nil {
		return err
	}

	if err := s.Set("quiet_period_wait_time", r.QuietPeriodWaitTime); err != nil {
		return err
	}

	if err := s.Set("quiet_period_max_retries", r.QuietPeriodMaxRetries); err != nil {
		return err
	}

	if err := s.Set("filter_pattern", r.FilterPattern); err != nil {
		return err
	}

	if err := s.Set("filter_pattern_regex", r.FilterPatternRegex); err != nil {
		return err
	}

	if err := s.Set("change_set_regex", r.ChangeSetRegex); err != nil {
		return err
	}

	return nil
}

func resourceLinkedRepositoryCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	linkedRepository := unmarshalLinkedRepository(data)

	newLinkedRepository, err := client.LinkedRepository.Create(linkedRepository)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(newLinkedRepository.Id)

	return resourceLinkedRepositoryRead(ctx, data, meta)
}

func resourceLinkedRepositoryRead(_ context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()
	linkedRepository, err := client.LinkedRepository.GetOne(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	if err = marshalLinkedRepository(data, linkedRepository); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceLinkedRepositoryUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*bamboo.Client)

	id := data.Id()
	linkedRepository := unmarshalLinkedRepository(data)

	err := client.LinkedRepository.Update(id, linkedRepository)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceLinkedRepositoryRead(ctx, data, meta)
}

func resourceLinkedRepositoryDelete(_ context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*bamboo.Client)

	id := data.Id()

	err := client.LinkedRepository.Delete(id)
	if err != nil {
		return ignoreNotFoundDiag(err, data)
	}

	data.SetId("")

	return diags
}

func resourceLinkedRepositoryImportState(
	ctx context.Context, data *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := data.Id()

	client := meta.(*bamboo.Client)

	linkedRepository, err := client.LinkedRepository.Search(name)
	if err != nil {
		return nil, err
	}

	if err = marshalLinkedRepository(data, linkedRepository); err != nil {
		return nil, err
	}

	data.SetId(linkedRepository.Id)

	return schema.ImportStatePassthroughContext(ctx, data, meta)
}
