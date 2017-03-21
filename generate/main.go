package main

import (
	"log"
	"os"
)

const (
	languagesTmplPath = "generate/languages.go.tmpl"
	tmplName          = "languages.go.tmpl"
	langFile          = "languages.go"
)

func main() {
	// creage or truncate languages.go file
	f, err := os.Create(langFile)
	if err != nil {
		log.Fatal(err)
	}

	generateLanguages(f, languagesTmplPath, tmplName)
}
