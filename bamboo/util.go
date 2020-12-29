package bamboo

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ExpandStringSlice(input []interface{}) *[]string {
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

func FlattenStringSlice(input *[]string) []interface{} {
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