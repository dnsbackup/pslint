// Package pslint is a linter for Public Suffix list.
package pslint

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

// A Linter lints a Public Suffix list source.
type Linter struct {
}

// Problem represents a problem in a Public Suffix list source.
type Problem struct {
	Line    *Line  // line in the source file
	Message string // a short explanation of the problem
	Level   Level  // a short string that represents the level (info, warning, error)
}

type Level string

const (
	LEVEL_WARN  = Level("warning")
	LEVEL_ERROR = Level("error")
)

func (l *Linter) newProblem(line *Line, message string, level Level) *Problem {
	problem := &Problem{
		Line:    line,
		Message: message,
		Level:   level,
	}
	return problem
}

func (l *Linter) lintString(src string) ([]Problem, error) {
	file := strings.NewReader(src)
	return l.lint(file)
}

func (l *Linter) lintFile(path string) ([]Problem, error) {
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

	for scanner.Scan() {
		index = index + 1
		line = &Line{number: index, source: scanner.Text()}

		if p, _ := l.checkSpaces(line); p != nil {
			problems = append(problems, *p)
		}
		if p, _ := l.checkRuleLowercase(line); p != nil {
			problems = append(problems, *p)
		}
		if p, _ := l.checkRuleEmptyLabels(line); p != nil {
			problems = append(problems, *p)
		}
	}
	return problems, nil
}

//func (l *Linter) Lint(lines []string) {
//
//}

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
