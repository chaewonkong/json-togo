package structstr_test

import (
	"strings"
	"testing"

	. "github.com/chaewonkong/json-togo/structstr"
)

func TestGenerate(t *testing.T) {
	var jsonMap map[string]any
	pkgName := "example"
	structName := "ExampleStruct"
	s := Generate(jsonMap, pkgName, structName)

	if s == "" {
		t.Error("Generate returned an empty string")
	}

	if strings.Contains(s, "package "+pkgName) == false {
		t.Errorf("Generate did not include package name %s", pkgName)
	}

	if strings.Contains(s, "type "+structName+" struct {") == false {
		t.Errorf("Generate did not include struct name %s", structName)
	}

	if s[len(s)-1] != '}' {
		t.Error("Generate did not end with '}'")
	}
}
