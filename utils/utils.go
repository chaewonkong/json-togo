package utils

import (
	"encoding/json"
	"strings"
	"unicode"
)

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

func InferTypeString(val any) string {
	switch v := val.(type) {
	case string:
		return "string"
	case json.Number:
		if strings.Contains(v.String(), ".") {
			return "float64"
		}
		return "int"
	case bool:
		return "bool"
	case []any:
		if len(v) > 0 {
			return "[]" + InferTypeString(v[0])
		}

		return "[]any"
	case map[string]any:
		return "map[string]any" // TODO: generate sub-struct
	default:
		return "any" // Default type for unknown types
	}
}
