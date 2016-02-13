// Package pslint is a linter for Public Suffix list
package pslint

import (
	"strings"
)

// line represents a single line being linted.
type Line struct {
	number int    // the line number
	source string // the source line
}

// isEmpty returns true if the line source contains only whitespaces.
func (l *Line) isEmpty() bool {
	return l.source == ""
}

// isBlank returns true if the line source is blank.
func (l *Line) isBlank() bool {
	return strings.TrimSpace(l.source) == ""
}

// isComment returns true if the line source is a comment.
func (l *Line) isComment() bool {
	text := strings.TrimSpace(l.source)
	return strings.HasPrefix(text, "//")
}

// isRule returns true if the line source is a rule.
func (l *Line) isRule() bool {
	return !l.isComment() && !l.isBlank()
}
