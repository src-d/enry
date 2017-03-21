package main

import (
	"errors"
	"io"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	// ErrExtensionsNotFound is the error returned if a yaml.MapSlice doesn't contain a key named extField.
	ErrExtensionsNotFound = errors.New("yaml.MapSlice doesn't contain extensions")
)

// generateLanguages read from buf and builds languages.go file from languagesTmplPath.
func generateLanguages(out io.Writer, buf []byte, languagesTmplPath, languagesTmpl string) error {
	// unmarshal yaml
	var yamlSlice yaml.MapSlice
	if err := yaml.Unmarshal(buf, &yamlSlice); err != nil {
		return err
	}

	// build the extension->languages map
	languagesByExtension, err := getExtMap(yamlSlice)
	if err != nil {
		return err
	}

	// generate languages.go from languages.go.tmpl
	fmap := template.FuncMap{
		"formatStringSlice": formatStringSlice,
	}

	t := template.Must(template.New(languagesTmpl).Funcs(fmap).ParseFiles(languagesTmplPath))
	if err := t.Execute(out, languagesByExtension); err != nil {
		return err
	}

	return nil
}

// get ExtMap takes in a yaml.MapSlice and builds the map of relations extensions->languages.
func getExtMap(yamlSlice yaml.MapSlice) (map[string][]string, error) {
	extLang := make(map[string][]string)
	for _, lang := range yamlSlice {
		extensions, err := findExtensions(lang.Value.(yaml.MapSlice))
		if err != nil && err != ErrExtensionsNotFound {
			return nil, err
		}

		fillMap(extLang, lang.Key.(string), extensions)
	}

	return extLang, nil
}

// findExtensions takes in a yaml.MapSlice and search for the field extField. It returns a slice of extensions
// found or ErrExtensionsNotFound if none extensions were found.
func findExtensions(items yaml.MapSlice) ([]interface{}, error) {
	const extField = "extensions"
	for _, item := range items {
		if item.Key == extField {
			return item.Value.([]interface{}), nil
		}
	}

	return nil, ErrExtensionsNotFound
}

// fillMap takes in a mapExt to insert the key lang, and give it the value of a []string from extensions.
func fillMap(mapExt map[string][]string, lang string, extensions []interface{}) {
	for _, extension := range extensions {
		ex := extension.(string)
		if _, ok := mapExt[ex]; !ok {
			mapExt[ex] = make([]string, 0, 10)
		}

		mapExt[ex] = append(mapExt[ex], lang)
	}
}

// formatStringSlice takes in slice and build the strings with the properly format for the template.
func formatStringSlice(slice []string) string {
	s := strings.Join(slice, `","`)
	s = `"` + s + `"`
	return s
}
