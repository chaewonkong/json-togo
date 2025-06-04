package structstr_test

import (
	"encoding/json"
	"testing"

	. "github.com/chaewonkong/json-togo/structstr"
)

func TestInferTypeString(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{"string", "string"},
		{input: json.Number("123"), expected: "int"},         // json.Number can be int or float
		{input: json.Number("123.456"), expected: "float64"}, // json.Number can be int or float
		{true, "bool"},
		{[]any{"item1", 2, 3.14}, "[]string"}, // Assuming first item is string
		{[]any{json.Number("1"), json.Number("2"), json.Number("3")}, "[]int"},
		{map[string]any{"key": "value"}, "map[string]any"},
		{nil, "any"}, // nil should return any
	}

	for _, test := range tests {
		result := InferTypeString(test.input)
		if result != test.expected {
			t.Errorf("InferTypeString(%v) = %q; want %q", test.input, result, test.expected)
		}
	}
}
