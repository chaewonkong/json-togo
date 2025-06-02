package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chaewonkong/json-togo/utils"
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
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// iterate map and generate struct fields
	for key, val := range jsonMap {
		fieldName := utils.ToPascalCase(key)
		fieldType := utils.InferTypeString(val)
		sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, fieldType, key))
	}
	sb.WriteString("}")

	// print and return
	fmt.Println(sb.String())

	return nil
}
