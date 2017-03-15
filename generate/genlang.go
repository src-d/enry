package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

const (
	languagesYAML     = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
	extField          = "extensions"
	languagesTmplPath = "generate/languages.go.tmpl"
	tmplName          = "languages.go.tmpl"
	langFile          = "languages.go"
)

var (
	ErrExtensionsNotFound = errors.New("yaml.MapSlice doesn't contain extensions")
)

func generateLanguages() {
	// get languages.yml
	res, err := http.Get(languagesYAML)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal yaml
	var out yaml.MapSlice
	if err := yaml.Unmarshal(buf, &out); err != nil {
		log.Fatal(err)
	}

	// build the extension->languages map
	languagesByExtension := make(map[string][]string)

	for _, lang := range out {
		extensions, err := findExtensions(lang.Value.(yaml.MapSlice), extField)
		if err != nil && err != ErrExtensionsNotFound {
			log.Println(err)
		}

		fillMap(languagesByExtension, lang.Key.(string), extensions)
	}

	// creage or truncate languages.go file
	f, err := os.Create(langFile)
	if err != nil {
		log.Fatal(err)
	}

	// generate languages.go from languages.go.tmpl
	fmap := template.FuncMap{
		"formatStringSlice": formatStringSlice,
	}

	t := template.Must(template.New(tmplName).Funcs(fmap).ParseFiles(languagesTmplPath))
	if err := t.Execute(f, languagesByExtension); err != nil {
		log.Fatal(err)
	}
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
	s := strings.Join(slice, `", "`)
	s = `"` + s + `"`
	return s
}
