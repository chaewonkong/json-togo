package structstr

import (
	"encoding/json"
	"strings"
)

// InferTypeString infers the Go type string from a given value.
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
