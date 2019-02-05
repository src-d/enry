// Package rule contains rule-based heuristic implementations.
// It is used in the generated code in content.go for disambiguation of languages
// with colliding extensions based on regexps from Linguist data.
package rule

import (
	"regexp"
)

// Heuristic consist of a number of rules where each, if matches,
// identifes content as belonging to a programming language(s).
type Heuristic interface {
	Matcher
	Languages() []string
}

// Matcher checks if the data matches (number of) pattern.
// Every rule below implements this interface: a rule is matcher that identifies
// given programming language(s) in case of the match.
type Matcher interface {
	Match(data []byte) bool
}

// languages struct incapsulate data common to every Matcher: all languages
// that it identifies.
type languages struct {
	langs []string
}

// Languages returns all languages, identified by this Matcher.
func (l *languages) Languages() []string {
	return l.langs
}

// MatchingLanguages is a helper to create new languages.
func MatchingLanguages(langs ...string) *languages {
	return &languages{langs}
}

// Implements a Heuristic.
type or struct {
	*languages
	Pattern *regexp.Regexp
}

// Or rule matches, if a single matching pattern exists.
// It defines only one pattern as it relies on compile-time optimization that
// represtes union with | in a single regexp.
func Or(l *languages, r *regexp.Regexp) *or {
	return &or{l, r}
}

// Match implements rule.Matcher.
func (r *or) Match(data []byte) bool {
	return r.Pattern.Match(data)
}

// Implements a Heuristic.
type and struct {
	*languages
	Patterns []Matcher
}

// And rule matches, if each of the patterns does match.
func And(l *languages, m ...Matcher) *and {
	return &and{l, m}
}

// Match implements data.Matcher.
func (r *and) Match(data []byte) bool {
	allMatch := true
	for _, p := range r.Patterns {
		if !p.Match(data) {
			allMatch = false
			break
		}
	}
	return allMatch
}

// Implements a Heuristic.
type not struct {
	*languages
	Patterns []*regexp.Regexp
}

// Not rule matches if none of the patterns match.
func Not(l *languages, r ...*regexp.Regexp) *not {
	return &not{l, r}
}

// Match implements data.Matcher.
func (r *not) Match(data []byte) bool {
	allDontMatch := true
	for _, p := range r.Patterns {
		if p.Match(data) {
			allDontMatch = false
			break
		}
	}
	return allDontMatch
}

// Implements a Heuristic.
type always struct {
	*languages
}

// Always rule always matches. Often is used as a default fallback.
func Always(l *languages) *always {
	return &always{l}
}

// Match implements Matcher.
func (r *always) Match(data []byte) bool {
	return true
}
