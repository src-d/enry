package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/src-d/simple-linguist/cli/slinguist-generate/generator"
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
		log.Println("couldn't find commit")
	}

	// generateFile(languagesYAML, langFile, languagesTmplPath, languagesTmpl, commit, generate.Languages)
	generateFile(heuristicsRuby, "/tmp/content.go", contentTmplPath, contentTmpl, commit, generator.Heuristics)
}

func getCommit(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func generateFile(fileToParsePath, outPath, tmplPath, tmpl, commit string, generate generator.Func) {
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Println(err)
		return
	}

	buf, err := loadFile(fileToParsePath)
	if err != nil {
		log.Println(err)
		return
	}

	if err := generate(outFile, buf, tmplPath, tmpl, commit); err != nil {
		log.Println(err)
	}
}

func loadFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
