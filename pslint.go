// Package pslint is a linter for Public Suffix list.
package pslint

import (
	"regexp"
)

// A Linter lints a Public Suffix list source.
type Linter struct {
}

// Problem represents a problem in a Public Suffix list source.
type Problem struct {
	Line       int    // line in the source file
	LineSource string // the source line
	Message    string // a short explanation of the problem
	Level      Level  // a short string that represents the level (info, warning, error)
}

type Level string

const (
	LEVEL_WARN  = Level("warning")
	LEVEL_ERROR = Level("error")
)

//func (l *Linter) Lint(lines []string) {
//
//}

// Spaces: checks the Line does not have irrelevant spaces.
//
// - The line should not have a leading space
// - The line should not have a trailing space
func (l *Linter) checkSpaces(line *line) (*Problem, error) {
	if match := regexp.MustCompile(`^\s`).MatchString(line.source); match {
		problem := &Problem{
			Line:       line.number,
			LineSource: line.source,
			Message:    "leading space",
			Level:      LEVEL_WARN,
		}
		return problem, nil
	}

	if match := regexp.MustCompile(`\s$`).MatchString(line.source); match {
		problem := &Problem{
			Line:       line.number,
			LineSource: line.source,
			Message:    "trailing space",
			Level:      LEVEL_WARN,
		}
		return problem, nil
	}

	return nil, nil
}

// Lowercase: checks the Rule is entirely lower-case.
func (l *Linter) checkRuleLowercase(line *line) (*Problem, error) {
	if !line.isRule() {
		return nil, nil
	}

	match := regexp.MustCompile(`[A-Z]`).MatchString(line.source)
	if match {
		problem := &Problem{
			Line:       line.number,
			LineSource: line.source,
			Message:    "non-lowercase suffix",
			Level:      LEVEL_ERROR,
		}
		return problem, nil
	}

	return nil, nil
}

// EmptyLabels: checks the Rule contains empty labels.
//
// An empty label is caused by two `..` (dots) with no content.
// The token `. .` will not be detected by this check, as there is already a more general check
// that checks the presence of spaces in a Rule.
func (l *Linter) checkRuleEmptyLabels(line *line) (*Problem, error) {
	if !line.isRule() {
		return nil, nil
	}

	match := regexp.MustCompile(`\.{2,}`).MatchString(line.source)
	if match {
		problem := &Problem{
			Line:       line.number,
			LineSource: line.source,
			Message:    "empty label",
			Level:      LEVEL_ERROR,
		}
		return problem, nil
	}

	return nil, nil
}
