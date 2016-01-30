// Package pslint is a linter for Public Suffix list
package pslint

import "regexp"

// Problem represents a problem in a Public Suffix list source.
type Problem struct {
	Line       int    // line in the source file
	LineSource string // the source line
	Message    string // a short explanation of the problem
	Level      Level  // a short string that represents the level (info, warning, error)
}

type Level string

// line represents a single line being linted.
type line struct {
	number int    // the line number
	source string // the source line
}

type Checker interface {
	CheckLine(l *line) (*Problem, error)
}

type baseChecker struct {
	name string
}

type regexpChecker struct {
	include  baseChecker

	regex    *regexp.Regexp
	message  string
	level    Level
	negative bool
}

func (c *regexpChecker) CheckLine(l *line) (problem *Problem, err error) {
	matched, err := c.regexp.MatchString(l.source)
	if err != nil {
		return nil, err
	}

	if matched {
		problem = &Problem{
			LineSource: l.number,
			LineSource: l.source,
			Message:    c.message,
			Level:      c.level,
		}
	}

	return problem, nil
}

const (
	LEVEL_WARN  = Level("warning")
	LEVEL_ERROR = Level("error")

	LINT_LOWER_CASE = regexpChecker{
		name:    "LowerCase",
		regexp:  regexp.MustCompile("[A-Z]"),
		message: "Fooo",
		level:   LEVEL_ERROR,
	}
)

var checks = map[string]Checker{
	"LowerCase": LINT_LOWER_CASE,
}

func lint(l) {
	l := &line{number:1, source:" foo.bar"}
	for _, check := range checks {
		p, err := check.CheckLine(l)
	}
}
