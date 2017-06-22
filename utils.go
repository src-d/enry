package enry

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
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

	gitAttributes      = map[string]bool{}
	languageAttributes = map[string]string{}
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
<<<<<<< HEAD
	return data.VendorMatchers.Match(path)
=======
	val, ok := gitAttributes[path]
	if ok {
		return val
	}

	return vendorMatchers.Match(path)
>>>>>>> 132a9bb... Added support for vendor and documentation attributes
}

// IsDocumentation returns whether or not path is a documentation path.
func IsDocumentation(path string) bool {
<<<<<<< HEAD
	return data.DocumentationMatchers.Match(path)
=======
	val, ok := gitAttributes[path]
	if ok {
		return val
	}

	return documentationMatchers.Match(path)
>>>>>>> 132a9bb... Added support for vendor and documentation attributes
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

func loadGitattributes() (map[string]string, error) {
	gitAttributes := map[string]string{}
	data, err := ioutil.ReadFile(".gitattributes")
	if err != nil {
		return nil, err
	}

	if data != nil {
		tokens := strings.Fields(string(data))
		for i := 0; i < len(tokens); i = i + 2 {
			gitAttributes[tokens[i]] = tokens[i+1]
		}
	}

	return gitAttributes, nil
}

func parseAttributes(attributes map[string]string) {
	for key, val := range attributes {
		switch {
		case val == "enry-vendored" || val == "enry-documentation":
			gitAttributes[key] = true
		case val == "enry-vendored=false" || val == "enry-documentation=false":
			gitAttributes[key] = false
		case strings.Contains(val, "enry-language="):
			tokens := strings.Split(val, "=")
			if len(tokens) == 2 {
				languageAttributes[key] = tokens[1]
			}
		}
	}
}

func init() {
	rawAttributes, err := loadGitattributes()
	if err == nil && len(rawAttributes) > 0 {
		parseAttributes(rawAttributes)
	}
}
