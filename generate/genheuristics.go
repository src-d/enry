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

const (
	heuristicsURL = "https://raw.githubusercontent.com/github/linguist/master/lib/linguist/heuristics.rb"
)

// LangHeur represents the relation between a language and the necessary regexp to apply heuristics.
type LangHeur map[string][]string

// Disamb contains the information to apply heuristics in base to an extension.
type Disamb struct {
	Extension string
	Languages LangHeur
}

// ParseDisamb takes in the line of the heuristics file where a disambiguate block
// starts and the Scanner to read. It returns a *[]Disamb, with the diasambiguators found.
func ParseDisamb(line string, buf *bufio.Scanner) []*Disamb {
	disambList := make([]*Disamb, 0, 2)
	splitted := strings.Fields(line)
	for _, v := range splitted {
		if strings.HasPrefix(v, `"`) {
			extension := strings.Trim(v, `",`)
			d := &Disamb{Extension: extension}
			disambList = append(disambList, d)
		}
	}

	langs := make([]string, 0, 2)
	regs := make([][]string, 0, 1)
	for buf.Scan() {
		l := buf.Text()
		if strings.Contains(line, "end") {
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

	for _, v := range disambList {
		v.Languages = languages
	}

	return disambList
}

var (
	disambiguators = make([]*Disamb, 0, 50)
	disambLine     = regexp.MustCompile(`^(\s*)disambiguate`)
)

func main() {
	// langsFound["Language X"] = []string{"regexp1", "regexp2", "regexp3"}
	// langsFound["Language Y"] = []string{"regexp4", "regexp4", "regexp5"}

	// fmt.Printf("%#v\n", langsFound)

	// for k, v := range langsFound {
	// 	fmt.Println(k, v)
	// }

	generateHeuristics()
}

func generateHeuristics() {
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

	for heuRB.Scan() {
		line := heuRB.Text()
		if disambLine.MatchString(line) {
			disambiguators = append(disambiguators, ParseDisamb(line, heuRB)...)
		}

	}

	// for _, d := range disambiguators {
	// 	fmt.Printf("%#v\n", d)
	// }

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(disambiguators); err != nil {
		log.Fatal(err)
	}
}
