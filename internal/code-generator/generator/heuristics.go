package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// GenHeuristics generates language identification heuristics in Go.
// It is of generator.File type.
func GenHeuristics(fileToParse, _, outPath, tmplPath, tmplName, commit string) error {
	heuristicsYaml, err := parseYaml(fileToParse)
	if err != nil {
		return err
	}

	langPatterns, err := loadHeuristics(heuristicsYaml)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = executeTemplate(buf, tmplName, tmplPath, commit, nil, langPatterns)
	if err != nil {
		return err
	}

	return formatedWrite(outPath, buf.Bytes())
}

func loadHeuristics(yaml *Heuristics) (map[string][]*LanguagePattern, error) {
	var patterns = make(map[string][]*LanguagePattern)
	for _, disambiguation := range yaml.Disambiguations {
		var rules []*LanguagePattern
		for _, rule := range disambiguation.Rules {
			langPattern := loadRule(yaml.NamedPatterns, rule)
			rules = append(rules, langPattern)
		}

		for _, ext := range disambiguation.Extensions {
			//  	 ["ext1", "ext2"]->[]*Rules =>
			//   		"ext1"->{lang: "language", patterns: ["pattern1", "pattern2"]}
			//          "ext2"->{ ... }
			if _, ok := patterns[ext]; ok {
				return nil, fmt.Errorf("cannt add extension '%s', it already exists for %q", ext, patterns[ext])
			}
			patterns[ext] = rules
		}

	}
	return patterns, nil
}

func loadRule(namedPatterns map[string]StringArray, rule *Rule) *LanguagePattern {
	// fmt.Printf("loading rule: \n\t%q\n", rule)
	var result *LanguagePattern
	if len(rule.And) != 0 { // - AndPattern
		var subPatterns []*LanguagePattern
		for _, r := range rule.And {
			subPattern := loadRule(namedPatterns, r)
			subPatterns = append(subPatterns, subPattern)
		}
		result = &LanguagePattern{"And", rule.Languages, "", subPatterns}
	} else if len(rule.Pattern) != 0 { // OrPattern
		// combines multiple patterns into a single one
		conjunction := strings.Join(rule.Pattern, " | ")
		result = &LanguagePattern{"Or", rule.Languages, conjunction, nil}
	} else if rule.NegativePattern != "" { // NotPattern
		result = &LanguagePattern{"Not", rule.Languages, rule.NegativePattern, nil}
	} else if rule.NamedPattern != "" { // Named OrPattern
		conjunction := strings.Join(namedPatterns[rule.NamedPattern], " | ")
		result = &LanguagePattern{"Or", rule.Languages, conjunction, nil}
	} else { // AlwaysPattern
		result = &LanguagePattern{"Always", rule.Languages, "", nil}
	}
	// fmt.Printf("\tgot: \n\t%q\n", result)
	return result
}

// LanguagePattern is a representation of parsed Rule.
// Strings are used as this will be consumed by a template.
type LanguagePattern struct {
	Op      string
	Langs   []string
	Pattern string
	Rules   []*LanguagePattern
}

type Heuristics struct {
	Disambiguations []*Disambiguation
	NamedPatterns   map[string]StringArray `yaml:"named_patterns"`
}

type Disambiguation struct {
	Extensions []string `yaml:"extensions,flow"`
	Rules      []*Rule  `yaml:"rules"`
}

type Rule struct {
	Patterns  `yaml:",inline"`
	Languages StringArray `yaml:"language"`
	And       []*Rule
}

type Patterns struct {
	Pattern         StringArray `yaml:"pattern,omitempty"`
	NamedPattern    string      `yaml:"named_pattern,omitempty"`
	NegativePattern string      `yaml:"negative_pattern,omitempty"`
}

// StringArray is workaround for parsing named_pattern,
// wich is sometimes arry and sometimes not.
// See https://github.com/go-yaml/yaml/issues/100
type StringArray []string

// UnmarshalYAML allowes to parse element always as a []string
func (sa *StringArray) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	if err := unmarshal(&multi); err != nil {
		var single string
		if err := unmarshal(&single); err != nil {
			return err
		}
		*sa = []string{single}
	} else {
		*sa = multi
	}
	return nil
}

func parseYaml(file string) (*Heuristics, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	h := &Heuristics{}
	if err := yaml.Unmarshal(data, &h); err != nil {
		return nil, err
	}

	return h, nil
}
