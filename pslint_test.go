package pslint

import (
	"testing"
)

func TestLinterCheckRuleLowercase(t *testing.T) {
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
		t.Fatalf("checkRuleLowercase problem message is %v, want %v", problemMessage, input)
	}

	if problemLevel := problem.Level; problemLevel != LEVEL_ERROR {
		t.Fatalf("checkRuleLowercase problem level is %v, want %v", problemLevel, LEVEL_ERROR)
	}
}

func TestLinterCheckRuleLowercaseValid(t *testing.T) {
	linter := &Linter{}
	inputs := []string{
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
}

func TestLinterCheckRuleLowercaseInValid(t *testing.T) {
	linter := &Linter{}
	inputs := []string{
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
