package enry

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type attrType int

const (
	vendor attrType = iota
	documentation
	generated
	language
)

const attrTypeName = "vendordocumentationgeneratedlanguage"

var attrTypeIndex = [...]uint8{0, 6, 19, 28, 36}

func (i attrType) String() string {
	if i < 0 || i >= attrType(len(attrTypeIndex)-1) {
		return fmt.Sprintf("attrType(%d)", i)
	}

	return attrTypeName[attrTypeIndex[i]:attrTypeIndex[i+1]]
}

type boolAttribute struct {
	kind       attrType
	matchers   []string
	attributes map[string]bool
}

type regExpAttribute struct {
	matchers   []string
	attributes map[*regexp.Regexp]string
}

// GitAttributes is a struct that contains two maps:
// boolAttributes contains all the attributes that works like a boolean condition,
// regExpAttributes contains all the attributes that match a regExp to choose if an attribute is applied or not
type GitAttributes struct {
	boolAttributes   map[attrType]boolAttribute
	regExpAttributes map[attrType]regExpAttribute
}

type overrideError struct {
	attribute attrType
	path      string
}

func (e *overrideError) Error() string {
	return fmt.Sprintf("gitattributes: You are overriding a %v attribute of one of your previous lines %s\n", e.attribute, e.path)
}

// Returns whether or not path is a vendor path.
func (gitAttrs *GitAttributes) IsVendor(path string) bool {
	if val, ok := gitAttrs.boolAttributes[vendor].attributes[path]; ok {
		return val
	}

	return IsVendor(path)
}

// Returns whether or not path is a documentation path.
func (gitAttrs *GitAttributes) IsDocumentation(path string) bool {
	if val, ok := gitAttrs.boolAttributes[documentation].attributes[path]; ok {
		return val
	}

	return IsDocumentation(path)
}

// Returns whether or not path is a generated path.
func (gitAttrs *GitAttributes) IsGenerated(path string) bool {
	if val, ok := gitAttrs.boolAttributes[generated].attributes[path]; ok {
		return val
	}

	return false
}

// GetLanguage get the language of a file matching the language attributes given.
// Returns either OthetLanguage or the language if the regExp matches
func (gitAttrs *GitAttributes) GetLanguage(filename string) string {
	for regExp, language := range gitAttrs.regExpAttributes[language].attributes {
		if regExp.MatchString(filename) {
			return language
		}
	}

	return OtherLanguage
}

// NewGitAttributes initialize a Gitattributes object
func NewGitAttributes() *GitAttributes {
	gitAttrs := GitAttributes{
		boolAttributes: map[attrType]boolAttribute{
			vendor:        boolAttribute{matchers: []string{"linguist-vendored", "linguist-vendored=false"}, attributes: map[string]bool{}},
			documentation: boolAttribute{matchers: []string{"linguist-documentation", "linguist-documentation=false"}, attributes: map[string]bool{}},
			generated:     boolAttribute{matchers: []string{"linguist-generated", "linguist-generated=false"}, attributes: map[string]bool{}},
		},
		regExpAttributes: map[attrType]regExpAttribute{
			language: regExpAttribute{matchers: []string{"linguist-language="}, attributes: map[*regexp.Regexp]string{}},
		},
	}

	return &gitAttrs
}

// LoadGitattributes reads and parses the file .gitattributes which overrides the standard strategies.
// Returns slice of errors that have may ocurred in the load.
func (gitAttrs *GitAttributes) LoadGitAttributes(path string, reader io.Reader) []error {
	rawAttributes, errArr := loadRawGitAttributes(reader)
	if len(rawAttributes) == 0 {
		return []error{}
	}

	return append(gitAttrs.parseAttributes(path, rawAttributes), errArr...)
}

func loadRawGitAttributes(reader io.Reader) (map[string][]string, []error) {
	rawAttributes := map[string][]string{}
	var errArr []error
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		errArr = append(errArr, err)
		return nil, errArr
	}

	if len(data) > 0 {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			err := loadLine(line, rawAttributes)
			if err != nil {
				errArr = append(errArr, err)
			}
		}
	}

	return rawAttributes, errArr
}

func loadLine(line string, gitattributes map[string][]string) error {
	tokens := strings.Fields(line)
	if len(tokens) == 2 {
		gitattributes[tokens[0]] = append(gitattributes[tokens[0]], tokens[1])
		return nil
	} else if len(tokens) != 0 {
		err := errors.New("gitattributes: Each line only can have a pair of elements  E.g. path/to/file attribute")
		return err
	}

	return nil
}

func (gitAttrs *GitAttributes) parseAttributes(path string, attributes map[string][]string) []error {
	errArray := []error{}
	for key, values := range attributes {
		for _, val := range values {
			err := gitAttrs.parseAttribute(path+key, val)
			if err != nil {
				errArray = append(errArray, err)
			}
		}
	}

	return errArray
}

func (gitAttrs *GitAttributes) matches(kind attrType, str string) bool {
	if bollAttrs, ok := gitAttrs.boolAttributes[kind]; ok && strings.Contains(str, bollAttrs.matchers[0]) {
		return true
	} else if regExpAttrs, ok := gitAttrs.regExpAttributes[kind]; ok && strings.Contains(str, regExpAttrs.matchers[0]) {
		return true
	}

	return false
}

func (gitAttrs *GitAttributes) parseAttribute(key string, attribute string) error {
	var err error
	for kind := vendor; kind <= language; kind++ {
		if gitAttrs.matches(kind, attribute) {
			if kind < language {
				err = gitAttrs.processBoolAttr(kind, key, attribute)
			} else {
				err = gitAttrs.processRegExpAttr(kind, key, attribute)
			}
		}
	}

	return err
}

func (gitAttrs *GitAttributes) processBoolAttr(kind attrType, key string, attribute string) error {
	var err error
	if _, ok := gitAttrs.boolAttributes[kind].attributes[key]; ok {
		err = &overrideError{attribute: kind, path: key}
	}

	switch {
	case attribute == gitAttrs.boolAttributes[kind].matchers[0]:
		gitAttrs.boolAttributes[kind].attributes[key] = true
	case attribute == gitAttrs.boolAttributes[kind].matchers[1]:
		gitAttrs.boolAttributes[kind].attributes[key] = false
	default:
		err = errors.New(fmt.Sprintf("gitattributes: The matcher %s doesn't exists\n", attribute))
	}

	return err
}

func (gitAttrs *GitAttributes) processRegExpAttr(kind attrType, regExpString string, attribute string) error {
	tokens := strings.SplitN(attribute, "=", 2)
	regExp, err := regexp.Compile(regExpString)
	if err != nil {
		return err
	}

	lang, _ := GetLanguageByAlias(tokens[1])
	if lang != OtherLanguage {
		gitAttrs.regExpAttributes[kind].attributes[regExp] = lang
	} else {
		gitAttrs.regExpAttributes[kind].attributes[regExp] = tokens[1]
	}

	return nil
}
