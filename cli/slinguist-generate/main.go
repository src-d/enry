package main

import (
	"io/ioutil"
	"log"

	"srcd.works/simple-linguist.v1/cli/slinguist-generate/generator"
)

const (
	languagesYAML     = ".linguist/lib/linguist/languages.yml"
	langFile          = "languages.go"
	languagesTmplPath = "cli/slinguist-generate/assets/languages.go.tmpl"
	languagesTmpl     = "languages.go.tmpl"

	heuristicsRuby  = ".linguist/lib/linguist/heuristics.rb"
	contentFile     = "content.go"
	contentTmplPath = "cli/slinguist-generate/assets/content.go.tmpl"
	contentTmpl     = "content.go.tmpl"

	commitPath = ".git/refs/heads/master"
)

func main() {
	commit, err := getCommit(commitPath)
	if err != nil {
		log.Printf("couldn't find commit: %v", err)
	}

	if err := generator.FromFile(languagesYAML, langFile, languagesTmplPath, languagesTmpl, commit, generator.Languages); err != nil {
		log.Println(err)
	}

	if err := generator.FromFile(heuristicsRuby, contentFile, contentTmplPath, contentTmpl, commit, generator.Heuristics); err != nil {
		log.Println(err)
	}
}

func getCommit(path string) (string, error) {
	commit, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(commit), nil
}
