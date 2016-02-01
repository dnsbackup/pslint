package pslint

import (
	"testing"
)

func TestLinter_CheckSpaces_Leading(t *testing.T) {
	linter := &Linter{}
	input := " rule"
	expectedLine := 2
	line := &line{source: input, number: expectedLine}

	problem, _ := linter.checkSpaces(line)
	if problem == nil {
		t.Fatalf("checkSpaces('%v') should NOT pass", input)
	}

	if problemLine := problem.Line; problemLine != expectedLine {
		t.Fatalf("checkSpaces problem line is %v, want %v", problemLine, expectedLine)
	}

	if problemSource := problem.LineSource; problemSource != input {
		t.Fatalf("checkSpaces problem source is %v, want %v", problemSource, input)
	}

	expectedMessage := "leading space"
	if problemMessage := problem.Message; problemMessage != expectedMessage {
		t.Fatalf("checkSpaces problem message is %v, want %v", problemMessage, expectedMessage)
	}

	if problemLevel := problem.Level; problemLevel != LEVEL_WARN {
		t.Fatalf("checkSpaces problem level is %v, want %v", problemLevel, LEVEL_WARN)
	}
}

func TestLinter_CheckSpaces_LeadingCases(t *testing.T) {
	linter := &Linter{}
	var inputs []string

	inputs = []string{
		"",
		"// a comment",
		".bad.suffix",
		"suffix",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkSpaces(line)
		if err != nil {
			t.Errorf("checkSpaces('%v') returned error: %v", input, err)
		}
		if problem != nil {
			t.Errorf("checkSpaces('%v') should pass", input)
		}
	}

	inputs = []string{
		" a leading space",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkSpaces(line)
		if err != nil {
			t.Errorf("checkSpaces('%v') returned error: %v", input, err)
		}
		if problem == nil {
			t.Errorf("checkSpaces('%v') should NOT pass", input)
		}
	}
}

func TestLinter_CheckSpaces_Trailing(t *testing.T) {
	linter := &Linter{}
	input := "rule "
	expectedLine := 2
	line := &line{source: input, number: expectedLine}

	problem, _ := linter.checkSpaces(line)
	if problem == nil {
		t.Fatalf("checkSpaces('%v') should NOT pass", input)
	}

	if problemLine := problem.Line; problemLine != expectedLine {
		t.Fatalf("checkSpaces problem line is %v, want %v", problemLine, expectedLine)
	}

	if problemSource := problem.LineSource; problemSource != input {
		t.Fatalf("checkSpaces problem source is %v, want %v", problemSource, input)
	}

	expectedMessage := "trailing space"
	if problemMessage := problem.Message; problemMessage != expectedMessage {
		t.Fatalf("checkSpaces problem message is %v, want %v", problemMessage, expectedMessage)
	}

	if problemLevel := problem.Level; problemLevel != LEVEL_WARN {
		t.Fatalf("checkSpaces problem level is %v, want %v", problemLevel, LEVEL_WARN)
	}
}

func TestLinter_CheckSpaces_TrailingCases(t *testing.T) {
	linter := &Linter{}
	var inputs []string

	inputs = []string{
		"",
		"// a comment",
		".bad.suffix",
		"suffix",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkSpaces(line)
		if err != nil {
			t.Errorf("checkSpaces('%v') returned error: %v", input, err)
		}
		if problem != nil {
			t.Errorf("checkSpaces('%v') should pass", input)
		}
	}

	inputs = []string{
		"a trailing space ",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkSpaces(line)
		if err != nil {
			t.Errorf("checkSpaces('%v') returned error: %v", input, err)
		}
		if problem == nil {
			t.Errorf("checkSpaces('%v') should NOT pass", input)
		}
	}
}

func TestLinter_CheckRuleLowercase(t *testing.T) {
	linter := &Linter{}
	input := "mixedCase"
	expectedLine := 2
	line := &line{source: input, number: expectedLine}

	problem, _ := linter.checkRuleLowercase(line)
	if problem == nil {
		t.Fatalf("checkRuleLowercase('%v') should NOT pass", input)
	}

	if problemLine := problem.Line; problemLine != expectedLine {
		t.Fatalf("checkRuleLowercase problem line is %v, want %v", problemLine, expectedLine)
	}

	if problemSource := problem.LineSource; problemSource != input {
		t.Fatalf("checkRuleLowercase problem source is %v, want %v", problemSource, input)
	}

	expectedMessage := "non-lowercase suffix"
	if problemMessage := problem.Message; problemMessage != expectedMessage {
		t.Fatalf("checkRuleLowercase problem message is %v, want %v", problemMessage, expectedMessage)
	}

	if problemLevel := problem.Level; problemLevel != LEVEL_ERROR {
		t.Fatalf("checkRuleLowercase problem level is %v, want %v", problemLevel, LEVEL_ERROR)
	}
}

func TestLinter_CheckRuleLowercase_Cases(t *testing.T) {
	linter := &Linter{}
	var inputs []string

	inputs = []string{
		"",
		"foo",
		"// a comment",
		".bad.suffix",

		"// A comment", // ignores comments even if mixed case
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkRuleLowercase(line)
		if err != nil {
			t.Errorf("checkRuleLowercase('%v') returned error: %v", input, err)
		}
		if problem != nil {
			t.Errorf("checkRuleLowercase('%v') should pass", input)
		}
	}

	inputs = []string{
		"mixedCase",
		"mixed.Case",
		"mixed.caSe",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkRuleLowercase(line)
		if err != nil {
			t.Errorf("checkRuleLowercase('%v') returned error: %v", input, err)
		}
		if problem == nil {
			t.Errorf("checkRuleLowercase('%v') should NOT pass", input)
		}
	}
}
