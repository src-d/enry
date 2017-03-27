package generator

import (
	"errors"
	"io"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	// ErrExtensionsNotFound is the error returned if a yaml.MapSlice doesn't contain a key named extField.
	ErrExtensionsNotFound = errors.New("extensions not found")
)

// Languages read from buf and builds languages.go file from languagesTmplPath.
func Languages(out io.Writer, buf []byte, languagesTmplPath, languagesTmpl, commit string) error {
	var yamlSlice yaml.MapSlice
	if err := yaml.Unmarshal(buf, &yamlSlice); err != nil {
		return err
	}

	languagesByExtension, err := buildExtensionLanguageMap(yamlSlice)
	if err != nil {
		return err
	}

	if err := executeTemplate(out, languagesByExtension, languagesTmplPath, languagesTmpl, commit); err != nil {
		return err
	}

	return nil
}

func buildExtensionLanguageMap(yamlSlice yaml.MapSlice) (map[string][]string, error) {
	extensionLangsMap := make(map[string][]string)
	for _, lang := range yamlSlice {
		extensions, err := findExtensions(lang.Value.(yaml.MapSlice))
		if err != nil && err != ErrExtensionsNotFound {
			return nil, err
		}

		fillMap(extensionLangsMap, lang.Key.(string), extensions)
	}

	return extensionLangsMap, nil
}

func findExtensions(items yaml.MapSlice) ([]string, error) {
	const extField = "extensions"
	for _, item := range items {
		if item.Key == extField {
			extensions := toStringSlice(item.Value.([]interface{}))
			return extensions, nil
		}
	}

	return nil, ErrExtensionsNotFound
}

func toStringSlice(slice []interface{}) []string {
	extensions := make([]string, 0, len(slice))
	for _, element := range slice {
		extension := element.(string)
		extensions = append(extensions, extension)
	}

	return extensions
}

func fillMap(extensionLangs map[string][]string, lang string, extensions []string) {
	for _, extension := range extensions {
		extensionLangs[extension] = append(extensionLangs[extension], lang)
	}
}

func executeTemplate(out io.Writer, languagesByExtension map[string][]string, languagesTmplPath, languagesTmpl, commit string) error {
	fmap := template.FuncMap{
		"getCommit":         func() string { return commit },
		"formatStringSlice": func(slice []string) string { return `"` + strings.Join(slice, `","`) + `"` },
	}

	t := template.Must(template.New(languagesTmpl).Funcs(fmap).ParseFiles(languagesTmplPath))
	if err := t.Execute(out, languagesByExtension); err != nil {
		return err
	}

	return nil
}
