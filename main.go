package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// read json

	var jsonMap map[string]any

	// read from stdin
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber() // UseNumber to preserve number types
	if err := decoder.Decode(&jsonMap); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	// TODO: get struc name from args or input
	// generate struct
	structName := "Data"
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n}", structName))

	// iterate map and generate struct fields
	for key, val := range jsonMap {
		fieldName := ToPascalCase(key)    // TODO make PascalCase
		fieldType := InterTypeString(val) // TODO: determine type based on value
		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
	}
	sb.WriteString("}")

	// print and return
	fmt.Println(sb.String())

	return nil
}

type MyStruct struct {
	YourStruct struct {
		Field1 string `json:"field_1"`
	}
}

func ToPascalCase(snake string) string {

	return ""
}

func InterTypeString(val any) string {
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
			return "[]" + InterTypeString(v[0])
		}

		return "[]any"
	case map[string]any:
		return "map[string]any" // TODO: generate sub-struct
	default:
		return "any" // Default type for unknown types
	}
}

/*
# 규칙

1. snake_case는 PascalCase로 변환한다.
2. JSON 태그는 snake_case를 유지한다.
3.
*/
