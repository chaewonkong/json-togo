package structstr

import (
	"fmt"
	"strings"
)

// Generate creates a Go struct definition from a JSON map.
func Generate(jsonMap map[string]any, pkgName, structName string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// iterate map and generate struct fields
	for key, val := range jsonMap {
		fieldName := ToPascalCase(key)
		fieldType := InferTypeString(val)
		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
	}
	sb.WriteString("}")

	return sb.String()
}
