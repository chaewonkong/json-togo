package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/chaewonkong/json-togo/utils"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var pkgName, structName, outFile string

	var rootCmd = &cobra.Command{
		Use:   "json-togo",
		Short: "Convert JSON to Go struct",
		Long:  `A simple CLI tool to convert JSON data to Go struct definitions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var jsonMap map[string]any

			// read from stdin
			decoder := json.NewDecoder(os.Stdin)
			decoder.UseNumber() // UseNumber to preserve number types
			if err := decoder.Decode(&jsonMap); err != nil {
				return fmt.Errorf("failed to decode JSON: %w", err)
			}

			// generate struct
			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
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

			// write to file if specified
			if outFile == "" {
				outFile = fmt.Sprintf("%s.go", pkgName)
			}
			if outFile != "" {
				file, err := os.Create(outFile)
				if err != nil {
					return fmt.Errorf("failed to create output file: %w", err)
				}
				defer file.Close()
				if _, err := file.WriteString(sb.String()); err != nil {
					return fmt.Errorf("failed to write to output file: %w", err)
				}
				fmt.Printf("Struct written to %s\n", outFile)
			}

			return nil
		},
	}

	rootCmd.Flags().StringVarP(&pkgName, "package", "p", "main", "Package name for the generated struct")
	rootCmd.Flags().StringVarP(&structName, "struct", "s", "Data", "Name of the struct to generate")
	rootCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file to write the struct definition")

	return rootCmd
}
