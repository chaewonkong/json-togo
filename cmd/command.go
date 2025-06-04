package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chaewonkong/json-togo/structstr"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var pkgName, structName, outFile string

	var rootCmd = &cobra.Command{
		Use:   "json-togo",
		Short: "Convert JSON to Go struct",
		Long:  `A simple CLI tool to convert JSON data to Go struct definitions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			runOpt := RunOptions{
				PackageName: pkgName,
				StructName:  structName,
				OutputFile:  outFile,
			}
			return Run(runOpt)
		},
	}

	rootCmd.Flags().StringVarP(&pkgName, "package", "p", "main", "Package name for the generated struct")
	rootCmd.Flags().StringVarP(&structName, "struct", "s", "Data", "Name of the struct to generate")
	rootCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file to write the struct definition")

	return rootCmd
}

type RunOptions struct {
	PackageName string
	StructName  string
	OutputFile  string
}

func Run(o RunOptions) error {
	var jsonMap map[string]any

	// read from stdin
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber() // UseNumber to preserve number types
	if err := decoder.Decode(&jsonMap); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	s := structstr.Generate(jsonMap, o.PackageName, o.StructName)

	// print and return
	fmt.Println(s)

	// write to file if specified
	if o.OutputFile == "" {
		o.OutputFile = fmt.Sprintf("%s.go", o.OutputFile)
	}

	file, err := os.Create(o.OutputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()
	if _, err := file.WriteString(s); err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}
	fmt.Printf("Struct written to %s\n", o.OutputFile)

	return nil
}
