package bamboo

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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