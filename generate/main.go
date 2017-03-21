package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	languagesURL      = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/languages.yml"
	langFile          = "languages.go"
	languagesTmplPath = "generate/languages.go.tmpl"
	languagesTmpl     = "languages.go.tmpl"

	heuristicsURL   = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/heuristics.rb"
	contentFile     = "content.go"
	contentTmplPath = "generate/content.go.tmpl"
	contentTmpl     = "content.go.tmpl"
)

// Generator is the type of functions that generate the files from templates.
type Generator func(io.Writer, []byte, string, string) error

func main() {
	genFile(languagesURL, langFile, languagesTmplPath, languagesTmpl, generateLanguages)
	// genFile(heuristicsURL, "/tmp/content.go", contentTmplPath, contentTmpl, generateHeuristics)
}

func genFile(fileURL, outFile, tmplPath, tmpl string, generate Generator) {
	f, err := os.Create(outFile)
	if err != nil {
		log.Println(err)
		return
	}

	buf, err := getRemoteFile(fileURL)
	if err != nil {
		log.Println(err)
		return
	}

	if err := generate(f, buf, tmplPath, tmpl); err != nil {
		log.Println(err)
	}
}

func getRemoteFile(fileURL string) ([]byte, error) {
	res, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
