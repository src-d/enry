package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// func main() {
// 	generateHeuristics()
// }

// LangHeur represents the relation between a language and the necessary regexp to apply heuristics.
type LangHeur map[string][]string

// Disamb contains the information to apply heuristics in base to an extension.
type Disamb struct {
	Extension string
	Languages LangHeur
}

func generateHeuristics() {
	const heuristicsURL = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/heuristics.rb"

	// get heuristics.rb
	res, err := http.Get(heuristicsURL)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	heuRB := bufio.NewScanner(bytes.NewReader(buf))
	disambiguators := getDisamb(heuRB)

	// debug
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(disambiguators); err != nil {
		log.Fatal(err)
	}
}

// getDisamb takes in a buf to parse and builds a slice of *Disamb to return.
func getDisamb(buf *bufio.Scanner) []*Disamb {
	disambLine := regexp.MustCompile(`^(\s*)disambiguate`)
	disambiguators := make([]*Disamb, 0, 50)
	for buf.Scan() {
		line := buf.Text()
		if disambLine.MatchString(line) {
			disambiguators = append(disambiguators, parseDisamb(line, buf)...)
		}

	}

	return disambiguators
}

// parseDisamb takes in the line of the heuristics file where a disambiguate block
// starts and the Scanner to read. It returns a *[]Disamb, with the diasambiguators
// found for this block.
func parseDisamb(line string, buf *bufio.Scanner) []*Disamb {
	disambList := make([]*Disamb, 0, 2)
	splitted := strings.Fields(line)
	// line looks like:	disambiguate ".lsp", ".lisp" do |data|
	// get the extensions associated to this disambiguate block
	for _, v := range splitted {
		if strings.HasPrefix(v, `"`) {
			extension := strings.Trim(v, `",`)
			d := &Disamb{Extension: extension}
			disambList = append(disambList, d)
		}
	}

	for _, v := range disambList {
		v.Languages = getLangHeur(buf)
	}

	return disambList
}

// getLangHeur builds the LangHeur associated to a disambiguate block.
func getLangHeur(buf *bufio.Scanner) LangHeur {
	langs := make([]string, 0, 2)
	regs := make([][]string, 0, 1)

	// a disambiguate block looks like:
	// disambiguate ".mod" do |data|
	// 	if data.include?('<!ENTITY ')
	// 		Language["XML"]
	// 	elsif /^\s*MODULE [\w\.]+;/i.match(data) || /^\s*END [\w\.]+;/i.match(data)
	// 		Language["Modula-2"]
	// 	else
	// 		[Language["Linux Kernel Module"], Language["AMPL"]]
	// 	end
	// end
	for buf.Scan() {
		l := buf.Text()
		if strings.Contains(l, "end") {
			break
		}

		if strings.Contains(l, ".match") {
			l = strings.TrimSpace(l)
			l = strings.TrimPrefix(l, "if ")
			l = strings.TrimPrefix(l, "elsif ")
			r := strings.Split(l, "||")
			for i, v := range r {
				r[i] = strings.TrimSpace(v)
				if !strings.Contains(v, ".match") {
					continue
				}

				r[i] = v[:strings.Index(v, ".match")]
			}

			regs = append(regs, r)
		}

		if strings.Contains(l, "Language") {
			langs = append(langs, strings.TrimSpace(l))
		}
	}

	languages := make(LangHeur)
	for i, v := range langs {
		v = strings.Trim(v, `[]`)
		languages[v] = nil
		if i < len(regs) {
			languages[v] = regs[i]
		}
	}

	return languages
}
