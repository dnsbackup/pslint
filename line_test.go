package pslint

import (
	"testing"
)

func TestLineIsEmpty(t *testing.T) {
	cases := map[string]bool{
		"\n":      false,
		"\t":      false,
		"":        true,
		" ":       false,
		"// foo":  false,
		" // foo": false,
		"foo":     false,
		" foo":    false,
		".foo":    false,
	}

	for input, expected := range cases {
		line := &Line{source: input}
		if actual := line.isEmpty(); actual != expected {
			t.Errorf("Expected isEmpty('%v') => %v, got %v", input, expected, actual)
		}
	}
}

func TestLineIsBlank(t *testing.T) {
	cases := map[string]bool{
		"\n":      true,
		"\t":      true,
		"":        true,
		" ":       true,
		"// foo":  false,
		" // foo": false,
		"foo":     false,
		" foo":    false,
		".foo":    false,
	}

	for input, expected := range cases {
		line := &Line{source: input}
		if actual := line.isBlank(); actual != expected {
			t.Errorf("Expected isBlank('%v') => %v, got %v", input, expected, actual)
		}
	}
}

func TestLineIsComment(t *testing.T) {
	cases := map[string]bool{
		"\n":      false,
		"\t":      false,
		"":        false,
		" ":       false,
		"// foo":  true,
		" // foo": true,
		"foo":     false,
		" foo":    false,
		".foo":    false,
	}

	for input, expected := range cases {
		line := &Line{source: input}
		if actual := line.isComment(); actual != expected {
			t.Errorf("Expected isComment('%v') => %v, got %v", input, expected, actual)
		}
	}
}

func TestLineIsRule(t *testing.T) {
	cases := map[string]bool{
		"\n":      false,
		"\t":      false,
		"":        false,
		" ":       false,
		"// foo":  false,
		" // foo": false,
		"foo":     true,
		" foo":    true,
		".foo":    true,
	}

	for input, expected := range cases {
		line := &Line{source: input}
		if actual := line.isRule(); actual != expected {
			t.Errorf("Expected isRule('%v') => %v, got %v", input, expected, actual)
		}
	}
}
