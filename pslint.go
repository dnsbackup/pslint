// Package pslint is a linter for Public Suffix list.
package pslint

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	LEVEL_WARN  = Level("warning")
	LEVEL_ERROR = Level("error")
)

// Level represents a problem level.
type Level string

// A Linter lints a Public Suffix list source.
type Linter struct {
	// When FailFast is true, the linter will stop running the tests for the entire list
	// on the first failed test.
	FailFast bool

	// When FailFirst is true, the linter will stop running the tests for the current line
	// on the first failed test.
	FailFirst bool
}

// NewLinter creates a new Linter with the recommended settings.
func NewLinter() *Linter {
	l := &Linter{FailFast: false, FailFirst: true}
	return l
}

// Problem represents a problem in a Public Suffix list source.
type Problem struct {
	Line    *Line  // line in the source file
	Message string // a short explanation of the problem
	Level   Level  // a short string that represents the level (info, warning, error)
}

func (l *Linter) newProblem(line *Line, message string, level Level) *Problem {
	problem := &Problem{
		Line:    line,
		Message: message,
		Level:   level,
	}
	return problem
}

// LintString lints the content of the string passed as argument one line at time.
func (l *Linter) LintString(src string) ([]Problem, error) {
	file := strings.NewReader(src)
	return l.lint(file)
}

// LintString reads the content from the file and lints one line at time.
func (l *Linter) LintFile(path string) ([]Problem, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return l.lint(file)
}

func (l *Linter) lint(r io.Reader) ([]Problem, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var line *Line
	var problems []Problem

	index := 0
	checks := l.ListChecks()

FileLoop:
	for scanner.Scan() {
		index = index + 1
		line = &Line{number: index, source: scanner.Text()}

	LineLoop:
		for _, check := range checks {
			if p, _ := check(line); p != nil {
				problems = append(problems, *p)
				if l.FailFast {
					break FileLoop
				}
				if l.FailFirst {
					break LineLoop
				}
			}
		}
	}
	return problems, nil
}

// CheckFunc represents a single check
type CheckFunc func(line *Line) (*Problem, error)

// ListCheck creates and returns a list of checks to be run.
func (l *Linter) ListChecks() []CheckFunc {
	return []CheckFunc{
		l.checkSpaces,
		l.checkRuleLowercase,
		l.checkRuleEmptyLabels,
	}
}

// Spaces: checks the Line does not have irrelevant spaces.
//
// - The line should not have a leading space
// - The line should not have a trailing space
func (l *Linter) checkSpaces(line *Line) (*Problem, error) {
	if match := regexp.MustCompile(`^\s`).MatchString(line.source); match {
		problem := l.newProblem(line, "leading space", LEVEL_WARN)
		return problem, nil
	}

	if match := regexp.MustCompile(`\s$`).MatchString(line.source); match {
		problem := l.newProblem(line, "trailing space", LEVEL_WARN)
		return problem, nil
	}

	return nil, nil
}

// Lowercase: checks the Rule is entirely lower-case.
func (l *Linter) checkRuleLowercase(line *Line) (*Problem, error) {
	if !line.isRule() {
		return nil, nil
	}

	match := regexp.MustCompile(`[A-Z]`).MatchString(line.source)
	if match {
		problem := l.newProblem(line, "non-lowercase suffix", LEVEL_ERROR)
		return problem, nil
	}

	return nil, nil
}

// EmptyLabels: checks the Rule contains empty labels.
//
// An empty label is caused by two `..` (dots) with no content.
// The token `. .` will not be detected by this check, as there is already a more general check
// that checks the presence of spaces in a Rule.
func (l *Linter) checkRuleEmptyLabels(line *Line) (*Problem, error) {
	if !line.isRule() {
		return nil, nil
	}

	match := regexp.MustCompile(`\.{2,}`).MatchString(line.source)
	if match {
		problem := l.newProblem(line, "empty label", LEVEL_ERROR)
		return problem, nil
	}

	return nil, nil
}
