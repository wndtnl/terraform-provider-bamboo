package bamboo

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bamboo "github.com/wndtnl/go-bamboo/pkg"
	"net/http"
)

func expandStringSlice(input []interface{}) *[]string {
	result := make([]string, 0)
	for _, item := range input {
		if item != nil {
			result = append(result, item.(string))
		} else {
			result = append(result, "")
		}
	}
	return &result
}

func flattenStringSlice(input *[]string) []interface{} {
	result := make([]interface{}, 0)
	if input != nil {
		for _, item := range *input {
			result = append(result, item)
		}
	}
	return result
}

// See: https://github.com/hashicorp/terraform-plugin-sdk/issues/534
func validateV1(fn schema.SchemaValidateFunc) schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		warnings, errors := fn(v, fmt.Sprintf("%#v", path))
		for _, w := range warnings {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Warning,
				Summary:       w,
				Detail:        w,
				AttributePath: path,
			})
		}
		for _, err := range errors {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Warning,
				Summary:       err.Error(),
				Detail:        err.Error(),
				AttributePath: path,
			})
		}
		return diags
	}
}

func ignoreNotFoundDiag(err error, data *schema.ResourceData) diag.Diagnostics {
	r := ignoreNotFound(err, data)
	if r != nil {
		return diag.FromErr(r)
	}
	return nil
}

func ignoreNotFound(err error, data *schema.ResourceData) error {
	if aErr, ok := err.(bamboo.Error); ok {
		if aErr.Status() == http.StatusNotFound {
			data.SetId("")
			return nil
		}
	}
	return err
}
