package data

import "regexp"

type (
	Heuristics []Matcher

	Matcher interface {
		Match(data []byte) bool
	}

	Languages struct {
		langs []string
	}

	Rule interface {
		GetLanguages() []string
	}
)

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

type OrRule struct {
	*Languages
	Pattern *regexp.Regexp
}

func (r *OrRule) Match(data []byte) bool {
	return r.Pattern.Match(data)
}

type AndRule struct {
	*Languages
	Patterns []Matcher
}

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

type NotRule struct {
	*Languages
	Patterns []*regexp.Regexp
}

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

type AlwaysRule struct {
	*Languages
}

func (r *AlwaysRule) Match(data []byte) bool {
	return true
}
