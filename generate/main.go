package main

import (
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

func main() {
	genFileLang(languagesURL, langFile, languagesTmplPath, languagesTmpl)
	genFileContent(heuristicsURL, contentFile, contentTmplPath, contentTmpl)
}

func genFileLang(languagesURL, langFile, languagesTmplPath, languagesTmpl string) {
	f, err := os.Create(langFile)
	if err != nil {
		log.Println(err)
		return
	}

	buf, err := getRemoteFile(languagesURL)
	if err != nil {
		log.Println(err)
		return
	}

	if err := generateLanguages(f, buf, languagesTmplPath, languagesTmpl); err != nil {
		log.Println(err)
	}
}

func genFileContent(heuristicsURL, contentFile, contentTmplPath, contentTmpl string) {
	// f, err := os.Create(heuristicsFile)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// buf, err := getRemoteFile(heuristicsURL)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// if err := generateHeuristics(f, buf, heuristicsTmplPath, heuristicsTmpl); err != nil {
	// 	log.Println(err)
	// }
}

func getRemoteFile(url string) ([]byte, error) {
	res, err := http.Get(languagesURL)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
