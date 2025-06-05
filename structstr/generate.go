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

	return generate(&sb, jsonMap, 4)
}

// TODO: test code
func generate(sb *strings.Builder, jsonMap map[string]any, tab int) string {
	if tab < 4 {
		tab = 4 // minimum indentation
	}

	indent := strings.Repeat(" ", tab)
	for key, val := range jsonMap {
		fieldName := ToPascalCase(key)
		fieldType := InferTypeString(val)
		if fieldType == "map[string]any" {
			sb.WriteString(fmt.Sprintf("%s%s struct {\n", indent, fieldName))
			generate(sb, val.(map[string]any), tab+4)
			continue
		}
		sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s\"`\n", indent, fieldName, fieldType, key))
	}
	sb.WriteString(fmt.Sprintf("%s}\n", strings.Repeat(" ", tab-4)))

	return sb.String()
}
