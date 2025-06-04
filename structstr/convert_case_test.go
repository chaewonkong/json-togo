package structstr_test

import (
	"testing"

	. "github.com/chaewonkong/json-togo/structstr"
)

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "HelloWorld"},
		{"snake_case_example", "SnakeCaseExample"},
		{"alreadyPascalCase", "AlreadyPascalCase"},
		{"", ""},
		{"snake_andCamelCase", "SnakeAndCamelCase"},
	}

	for _, test := range tests {
		result := ToPascalCase(test.input)
		if result != test.expected {
			t.Errorf("ToPascalCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
