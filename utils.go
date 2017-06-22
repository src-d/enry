package enry

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

	gitattributes         = map[string]bool{}
	languageGitattributes = map[*regexp.Regexp]string{}
)

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
	if val, ok := gitattributes[path]; ok {
		return val
	}

	return vendorMatchers.Match(path)
}

// IsDocumentation returns whether or not path is a documentation path.
func IsDocumentation(path string) bool {
	if val, ok := gitattributes[path]; ok {
		return val
	}

	return documentationMatchers.Match(path)
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

// LoadGitattributes reads and parse the file .gitattributes wich overrides the standards strategies
func LoadGitattributes() {
	rawAttributes, err := loadGitattributes()
	if err == nil && len(rawAttributes) > 0 {
		parseAttributes(rawAttributes)
	}
}

func loadGitattributes() (map[string]string, error) {
	gitattributes := map[string]string{}
	data, err := ioutil.ReadFile(".gitattributes")
	if err != nil {
		if err != os.ErrNotExist {
			log.Println(".gitattributes: " + err.Error())
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

func loadLine(line string, gitattributes map[string]string) {
	tokens := strings.Fields(line)
	if len(tokens) == 2 {
		gitattributes[tokens[0]] = tokens[1]
	} else {
		log.Println(".gitattributes: Each line only can have a pair of elements\nE.g. /path/to/file attribute")
	}
}

func parseAttributes(attributes map[string]string) {
	for key, val := range attributes {
		if isInGitattributes(key) {
			log.Printf("You are overriding one of your previous lines %s", key)
		}
		switch {
		case val == "enry-vendored" || val == "enry-documentation":
			gitattributes[key] = true
		case val == "enry-vendored=false" || val == "enry-documentation=false":
			gitattributes[key] = false
		case strings.Contains(val, "enry-language="):
			processLanguageAttr(key, val)
		default:
			log.Printf("The matcher %s doesn't exists\n", val)
		}
	}
}

func isInGitattributes(key string) bool {
	if _, ok := gitattributes[key]; ok {
		return ok
	}

	regExp, err := regexp.Compile(key)
	if err == nil {
		if _, ok := languageGitattributes[regExp]; ok {
			return ok
		}
	}

	return false
}

func processLanguageAttr(regExpString string, attribute string) {
	tokens := strings.Split(attribute, "=")
	if len(tokens) == 2 {
		regExp, err := regexp.Compile(regExpString)
		if err == nil {
			languageGitattributes[regExp] = tokens[1]
		}
	}
}
