package data

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testContentHeuristics = map[string]*Heuristics{
	".md": &Heuristics{ // final pattern for parsed YAML rule
		&OrRule{
			&Languages{[]string{"Markdown"}},
			regexp.MustCompile(`(^[-A-Za-z0-9=#!\*\[|>])|<\/ | \A\z`),
		},
		&OrRule{
			&Languages{[]string{"GCC Machine Description"}},
			regexp.MustCompile(`^(;;|\(define_)`),
		},
		&AlwaysRule{
			&Languages{[]string{"Markdown"}},
		},
	},
	".ms": &Heuristics{
		// Order defines precedence: And, Or, Not, Named, Always
		&AndRule{
			&Languages{[]string{"Unix Assembly"}},
			[]Matcher{
				&NotRule{
					nil,
					[]*regexp.Regexp{regexp.MustCompile(`/\*`)},
				},
				&OrRule{
					nil,
					regexp.MustCompile(`^\s*\.(?:include\s|globa?l\s|[A-Za-z][_A-Za-z0-9]*:)`),
				},
			},
		},
		&OrRule{
			&Languages{[]string{"Roff"}},
			regexp.MustCompile(`^[.''][A-Za-z]{2}(\s|$)`),
		},
		&AlwaysRule{
			&Languages{[]string{"MAXScript"}},
		},
	},
}

func TestContentHeuristics_MatchingAlways(t *testing.T) {
	lang := testContentHeuristics[".md"].matchString("")
	assert.Equal(t, []string{"Markdown"}, lang)

	lang = testContentHeuristics[".ms"].matchString("")
	assert.Equal(t, []string{"MAXScript"}, lang)
}

func TestContentHeuristics_MatchingAnd(t *testing.T) {
	lang := testContentHeuristics[".md"].matchString(";;")
	assert.Equal(t, []string{"GCC Machine Description"}, lang)
}

func TestContentHeuristics_MatchingOr(t *testing.T) {
	lang := testContentHeuristics[".ms"].matchString("	.include \"math.s\"")
	assert.Equal(t, []string{"Unix Assembly"}, lang)
}
