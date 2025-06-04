package structstr

import (
	"strings"
	"unicode"
)

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(snake string) string {
	var sb strings.Builder
	toUpper := false
	for idx, r := range snake {
		if idx == 0 || toUpper {
			r = unicode.ToUpper(r)
			toUpper = false // reset
		}

		if r == '_' {
			toUpper = true
			continue // skip underscore
		}

		sb.WriteRune(r)
	}

	return sb.String()
}
