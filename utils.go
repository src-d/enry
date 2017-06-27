package enry

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/src-d/enry.v1/data"
)

var (
	auxiliaryLanguages = map[string]bool{
		"Other": true, "XML": true, "YAML": true, "TOML": true, "INI": true,
		"JSON": true, "TeX": true, "Public Key": true, "AsciiDoc": true,
		"AGS Script": true, "VimL": true, "Diff": true, "CMake": true, "fish": true,
		"Awk": true, "Graphviz (DOT)": true, "Markdown": true, "desktop": true,
		"XSLT": true, "SQL": true, "RMarkdown": true, "IRC log": true,
		"reStructuredText": true, "Twig": true, "CSS": true, "Batchfile": true,
		"Text": true, "HTML+ERB": true, "HTML": true, "Gettext Catalog": true,
		"Smarty": true, "Raw token data": true,
	}

	configurationLanguages = map[string]bool{
		"XML": true, "JSON": true, "TOML": true, "YAML": true, "INI": true, "SQL": true,
	}

	vendorGitattributes        = map[string]bool{}
	documentationGitattributes = map[string]bool{}
	languageGitattributes      = map[*regexp.Regexp]string{}
)

type OverrideError struct {
	attribute string
	path      string
}

func (e *OverrideError) Error() string {
	return fmt.Sprintf(".gitattributes: You are overriding a %s attribute of one of your previous lines %s\n", e.attribute, e.path)
}

// IsAuxiliaryLanguage returns whether or not lang is an auxiliary language.
func IsAuxiliaryLanguage(lang string) bool {
	_, ok := auxiliaryLanguages[lang]
	return ok
}

// IsConfiguration returns whether or not path is using a configuration language.
func IsConfiguration(path string) bool {
	language, _ := GetLanguageByExtension(path)
	_, is := configurationLanguages[language]
	return is
}

// IsDotFile returns whether or not path has dot as a prefix.
func IsDotFile(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

// IsVendor returns whether or not path is a vendor path.
func IsVendor(path string) bool {
	if val, ok := vendorGitattributes[path]; ok {
		return val
	}

	return data.VendorMatchers.Match(path)
}

// IsDocumentation returns whether or not path is a documentation path.
func IsDocumentation(path string) bool {
	if val, ok := documentationGitattributes[path]; ok {
		return val
	}

	return data.DocumentationMatchers.Match(path)
}

const sniffLen = 8000

// IsBinary detects if data is a binary value based on:
// http://git.kernel.org/cgit/git/git.git/tree/xdiff-interface.c?id=HEAD#n198
func IsBinary(data []byte) bool {
	if len(data) > sniffLen {
		data = data[:sniffLen]
	}

	if bytes.IndexByte(data, byte(0)) == -1 {
		return false
	}

	return true
}

// LoadGitattributes reads and parses the file .gitattributes which overrides the standard strategies
func LoadGitattributes() {
	rawAttributes, err := loadRawGitattributes(".gitattributes")
	if err == nil && len(rawAttributes) > 0 {
		parseAttributes(rawAttributes)
	}
}

func loadRawGitattributes(name string) (map[string][]string, error) {
	gitattributes := map[string][]string{}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		if err != os.ErrNotExist {
			log.Println(name + ": " + err.Error())
		}

		return nil, err
	}

	if len(data) > 0 {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			loadLine(line, gitattributes)
		}
	}

	return gitattributes, nil
}

func loadLine(line string, gitattributes map[string][]string) error {
	tokens := strings.Fields(line)
	if len(tokens) == 2 {

		gitattributes[tokens[0]] = append(gitattributes[tokens[0]], tokens[1])
		return nil
	} else if len(tokens) != 0 {
		err := errors.New(".gitattributes: Each line only can have a pair of elements  E.g. path/to/file attribute")
		log.Println(err.Error())
		return err
	}
	return nil
}

func parseAttributes(attributes map[string][]string) []error {
	errArray := []error{}
	for key, values := range attributes {
		for _, val := range values {
			err := parseAttribute(key, val)
			if err != nil {
				errArray = append(errArray, err)
			}
		}
	}

	return errArray
}

func parseAttribute(key string, attribute string) error {
	var err error
	switch {
	case strings.Contains(attribute, "linguist-vendored"):
		err = processVendorAttr(key, attribute)
	case strings.Contains(attribute, "linguist-documentation"):
		err = processDocumentationAttr(key, attribute)
	case strings.Contains(attribute, "linguist-language="):
		err = processLanguageAttr(key, attribute)
	default:
		err = errors.New(fmt.Sprintf("gitattributes: The matcher %s doesn't exists\n", attribute))
		log.Printf(err.Error())
	}
	return err
}

func processVendorAttr(key string, attribute string) error {
	var err error
	if _, ok := vendorGitattributes[key]; ok {
		err = &OverrideError{attribute: "vendor", path: key}
	}

	switch {
	case attribute == "linguist-vendored":
		vendorGitattributes[key] = true
	case attribute == "linguist-vendored=false":
		vendorGitattributes[key] = false
	}

	return err
}

func processDocumentationAttr(key string, attribute string) error {
	var err error
	if _, ok := documentationGitattributes[key]; ok {
		err = &OverrideError{attribute: "documentation", path: key}
	}

	switch {
	case attribute == "linguist-documentation":
		documentationGitattributes[key] = true
	case attribute == "linguist-documentation=false":
		documentationGitattributes[key] = false
	}

	return err
}

func processLanguageAttr(regExpString string, attribute string) error {
	tokens := strings.SplitN(attribute, "=", 2)
	regExp, err := regexp.Compile(regExpString)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	lang, _ := GetLanguageByAlias(tokens[1])
	if lang != OtherLanguage {
		languageGitattributes[regExp] = lang
	} else {
		languageGitattributes[regExp] = tokens[1]
	}

	return nil
}
