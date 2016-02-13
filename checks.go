package pslint

import (
	"regexp"
)

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
