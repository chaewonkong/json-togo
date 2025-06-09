package cmd

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"

	"github.com/chaewonkong/json-togo/structstr"
	"github.com/spf13/cobra"
)

// NewCommand creates a new cobra command for the json-togo CLI tool.
func New() *cobra.Command {
	var pkgName, structName, outFile string

	var rootCmd = &cobra.Command{
		Use:   "json-togo",
		Short: "Convert JSON to Go struct",
		Long:  `A simple CLI tool to convert JSON data to Go struct definitions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(pkgName, structName, outFile)
		},
	}

	rootCmd.Flags().StringVarP(&pkgName, "package", "p", "main", "Package name for the generated struct")
	rootCmd.Flags().StringVarP(&structName, "struct", "s", "Data", "Name of the struct to generate")
	rootCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file to write the struct definition")

	return rootCmd
}

func run(pkgName, structName, outputFile string) error {
	var jsonMap map[string]any

	// read from stdin
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber() // UseNumber to preserve number types
	if err := decoder.Decode(&jsonMap); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	s := structstr.Generate(jsonMap, pkgName, structName)

	// format the generated struct
	formatted, err := format.Source([]byte(s))
	if err != nil {
		return fmt.Errorf("failed to format Go source: %w", err)
	}

	// write to file if specified
	if outputFile == "" {
		outputFile = fmt.Sprintf("%s.go", pkgName)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()
	if _, err := file.Write(formatted); err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}

	// print results
	fmt.Printf("Struct written to: %s\n", outputFile)
	fmt.Printf("Generated struct:\n\n%s\n", formatted)

	return nil
}
