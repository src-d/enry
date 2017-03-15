package main

import (
	"log"
	"os"
)

func main() {
	// creage or truncate languages.go file
	f, err := os.Create(langFile)
	if err != nil {
		log.Fatal(err)
	}

	generateLanguages(f, languagesTmplPath, tmplName)
}
