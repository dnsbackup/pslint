// Package pslint is a linter for Public Suffix list.
package pslint

import (
	"bufio"
	"io"
	"os"
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
	Line       int    // line in the source file
	LineSource string // the source line
	Message    string // a short explanation of the problem
	Level      Level  // a short string that represents the level (info, warning, error)
}

func (l *Linter) newProblem(line *line, message string, level Level) *Problem {
	problem := &Problem{
		Line:       line.number,
		LineSource: line.source,
		Message:    message,
		Level:      level,
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

	var sline *line
	var problems []Problem

	index := 0
	checks := l.ListChecks()

FileLoop:
	for scanner.Scan() {
		index = index + 1
		sline = &line{number: index, source: scanner.Text()}

	LineLoop:
		for _, check := range checks {
			if p, _ := check(sline); p != nil {
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
type CheckFunc func(line *line) (*Problem, error)
