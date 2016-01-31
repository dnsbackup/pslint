// Package pslint is a linter for Public Suffix list
package pslint

import (
	"strings"
)

// line represents a single line being linted.
type line struct {
	number int    // the line number
	source string // the source line
}

// isEmpty returns true if the line source contains only whitespaces.
func (l *line) isEmpty() bool {
	return l.source == ""
}

// isBlank returns true if the line source is blank.
func (l *line) isBlank() bool {
	return strings.TrimSpace(l.source) == ""
}

// isComment returns true if the line source is a comment.
func (l *line) isComment() bool {
	text := strings.TrimSpace(l.source)
	return strings.HasPrefix(text, "//")
}

// isRule returns true if the line source is a rule.
func (l *line) isRule() bool {
	return !l.isComment() && !l.isBlank()
}
