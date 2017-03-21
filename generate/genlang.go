package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	languagesURL = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
	extField     = "extensions"
)

var (
	// ErrExtensionsNotFound is the error returned if a yaml.MapSlice doesn't contain a key named extField.
	ErrExtensionsNotFound = errors.New("yaml.MapSlice doesn't contain extensions")
)

func generateLanguages(out io.Writer, languagesTmplPath, tmplName string) {
	// get languages.yml
	res, err := http.Get(languagesURL)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal yaml
	var yamlSlice yaml.MapSlice
	if err := yaml.Unmarshal(buf, &yamlSlice); err != nil {
		log.Fatal(err)
	}

	// build the extension->languages map
	languagesByExtension := getExtMap(yamlSlice)

	// generate languages.go from languages.go.tmpl
	fmap := template.FuncMap{
		"formatStringSlice": formatStringSlice,
	}

	t := template.Must(template.New(tmplName).Funcs(fmap).ParseFiles(languagesTmplPath))
	if err := t.Execute(out, languagesByExtension); err != nil {
		log.Fatal(err)
	}
}

func getExtMap(yamlSlice yaml.MapSlice) map[string][]string {
	extLang := make(map[string][]string)
	for _, lang := range yamlSlice {
		extensions, err := findExtensions(lang.Value.(yaml.MapSlice), extField)
		if err != nil && err != ErrExtensionsNotFound {
			log.Println(err)
		}

		fillMap(extLang, lang.Key.(string), extensions)
	}

	return extLang
}

func findExtensions(items yaml.MapSlice, key string) ([]interface{}, error) {
	for _, item := range items {
		if item.Key == key {
			return item.Value.([]interface{}), nil
		}
	}

	return nil, ErrExtensionsNotFound
}

func fillMap(mapExt map[string][]string, lang string, extensions []interface{}) {
	for _, extension := range extensions {
		ex := extension.(string)
		if _, ok := mapExt[ex]; !ok {
			mapExt[ex] = make([]string, 0, 10)
		}

		mapExt[ex] = append(mapExt[ex], lang)
	}
}

func formatStringSlice(slice []string) string {
	s := strings.Join(slice, `","`)
	s = `"` + s + `"`
	return s
}
