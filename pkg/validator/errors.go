package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) string {
	result := ""
	for field, err := range errors.Errors {
		result += field + ": " + strings.Join(err, ", ") + "\r\n"
	}
	return result
}
