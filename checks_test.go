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

	if want, got := line.source, problem.LineSource; got != want {
		t.Fatalf("checkSpaces problem line is %v, want %v", got, input)
	}

	if want, got := line.number, problem.Line; got != want {
		t.Fatalf("checkSpaces problem line number is %v, want %v", got, want)
	}

	if want, got := "leading space", problem.Message; want != got {
		t.Fatalf("checkSpaces problem message is %v, want %v", got, want)
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
		"suffix",

		".bad.suffix",
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

	if want, got := line.source, problem.LineSource; got != want {
		t.Fatalf("checkSpaces problem line is %v, want %v", got, input)
	}

	if want, got := line.number, problem.Line; got != want {
		t.Fatalf("checkSpaces problem line number is %v, want %v", got, want)
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
		"suffix",

		".bad.suffix",
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

	if want, got := line.source, problem.LineSource; got != want {
		t.Fatalf("checkRuleLowercase problem line is %v, want %v", got, input)
	}

	if want, got := line.number, problem.Line; got != want {
		t.Fatalf("checkRuleLowercase problem line number is %v, want %v", got, want)
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
		"// A comment", // ignore comments even if mixed case
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

func TestLinter_CheckRuleEmptyLabels(t *testing.T) {
	linter := &Linter{}
	input := "foo..bar"
	expectedLine := 2
	line := &line{source: input, number: expectedLine}

	problem, _ := linter.checkRuleEmptyLabels(line)
	if problem == nil {
		t.Fatalf("checkRuleEmptyLabels('%v') should NOT pass", input)
	}

	if want, got := line.source, problem.LineSource; got != want {
		t.Fatalf("checkRuleEmptyLabels problem line is %v, want %v", got, input)
	}

	if want, got := line.number, problem.Line; got != want {
		t.Fatalf("checkRuleEmptyLabels problem line number is %v, want %v", got, want)
	}

	if problemLevel := problem.Level; problemLevel != LEVEL_ERROR {
		t.Fatalf("checkRuleEmptyLabels problem level is %v, want %v", problemLevel, LEVEL_ERROR)
	}
}

func TestLinter_CheckRuleEmptyLabel_Cases(t *testing.T) {
	linter := &Linter{}
	var inputs []string

	inputs = []string{
		"",
		"foo",
		"foo.bar",
		"foo.bar.baz",
		"// .. this is a comment",

		".bad.suffix", // ignore bad leading dots
		"bad.suffix.", // ignore bad trailing dots
		"foo. .bar",   // consider spaces as non-empty
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkRuleEmptyLabels(line)
		if err != nil {
			t.Errorf("checkRuleEmptyLabels('%v') returned error: %v", input, err)
		}
		if problem != nil {
			t.Errorf("checkRuleEmptyLabels('%v') should pass", input)
		}
	}

	inputs = []string{
		"foo..bar",
		"foo.bar..",
		"..foo.bar",
		"foo...bar",
		"foo.. .bar",
		"foo.bar..baz",
		"foo..bar.baz",
	}
	for _, input := range inputs {
		line := &line{source: input, number: 2}
		problem, err := linter.checkRuleEmptyLabels(line)
		if err != nil {
			t.Errorf("checkRuleEmptyLabels('%v') returned error: %v", input, err)
		}
		if problem == nil {
			t.Errorf("checkRuleEmptyLabels('%v') should NOT pass", input)
		}
	}
}
