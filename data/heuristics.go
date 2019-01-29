package data

// Implmentation of a rule-based content heuristics matching engine.
// Every Rule defines a patterns that content must match in order to be identifed as
// belonging to a language(s).
// It is used to generate a content.go code for disambiguation of languages with
// colliding extensions based on regexps from Linguist.

import "regexp"

type (
	// Heuristics consists of a number of sequntially applied Matchers.
	Heuristics []Matcher

	// Matcher checks if a given data matches (number of) patterns.
	Matcher interface {
		Match(data []byte) bool
	}

	// Languages incapsulates data common to every Rule: number of languages
	// it identifies.
	Languages struct {
		langs []string
	}

	// Rule interface provides access to a languages that this rule identifies.
	Rule interface {
		GetLanguages() []string
	}
)

// Match returns languages identified by the matching rules of the heuristic.
func (h *Heuristics) Match(data []byte) []string {
	var matchedLangs []string
	for _, matcher := range *h {
		if matcher.Match(data) {
			for _, langOrAlias := range matcher.(Rule).GetLanguages() {
				lang, ok := LanguageByAlias(langOrAlias)
				if !ok { // should never happen
					// language name/alias in heuristics.yml is not consistent with languages.yml
					// but we do not surface any error on the API
					continue
				}
				matchedLangs = append(matchedLangs, lang)
			}
			break
		}
	}
	return matchedLangs
}

// matchString is a convenience used only in tests.
func (h *Heuristics) matchString(data string) []string {
	return h.Match([]byte(data))
}

// GetLanguages returns languages, defined by this data.Rule.
func (l *Languages) GetLanguages() []string {
	return l.langs
}

// OrRule matches if a single matching pattern exists.
// It defines only one pattern as it relis on compile-time optimization that
// represtes union with | in a single regexp pattern.
type OrRule struct {
	*Languages
	Pattern *regexp.Regexp
}

// Match implements data.Matcher.
func (r *OrRule) Match(data []byte) bool {
	return r.Pattern.Match(data)
}

// AndRule matches if all of the patterns match.
type AndRule struct {
	*Languages
	Patterns []Matcher
}

// Match implements data.Matcher.
func (r *AndRule) Match(data []byte) bool {
	allMatch := true
	for _, p := range r.Patterns {
		if !p.Match(data) {
			allMatch = false
			break
		}
	}
	return allMatch
}

// NotRule matches if none of the patterns match.
type NotRule struct {
	*Languages
	Patterns []*regexp.Regexp
}

// Match implements data.Matcher.
func (r *NotRule) Match(data []byte) bool {
	allDontMatch := true
	for _, p := range r.Patterns {
		if p.Match(data) {
			allDontMatch = false
			break
		}
	}
	return allDontMatch
}

// AlwaysRule always matches.
// Used as default fallback.
type AlwaysRule struct {
	*Languages
}

// Match implements data.Matcher.
func (r *AlwaysRule) Match(data []byte) bool {
	return true
}
