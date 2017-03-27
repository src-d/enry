package generator

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	disambLine  = regexp.MustCompile(`^(\s*)disambiguate`)
	definedRegs = make(map[string]string)
)

type languageHeuristics struct {
	Language       string   `json:"language,omitempty"`
	Heuristics     []string `json:"heuristics,omitempty"`
	LogicRelations []string `json:"logic_relations,omitempty"`
}

type disambiguator struct {
	Extension string
	Languages []*languageHeuristics
}

// Heuristics read from buf and builds content.go file from contentTmplPath.
func Heuristics(out io.Writer, heuristics []byte, contentTmplPath, contentTmpl, commit string) error {
	disambiguators, err := getDisambiguators(heuristics)
	if err != nil {
		return err
	}

	// debug
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if err := enc.Encode(disambiguators); err != nil {
		return err
	}
	//

	return nil
}

// A disambiguate block looks like:
// disambiguate ".mod", ".extension" do |data|
// 	if data.include?('<!ENTITY ') && data.include?('patata')
// 		Language["XML"]
// 	elsif /^\s*MODULE [\w\.]+;/i.match(data) || /^\s*END [\w\.]+;/i.match(data) || data.empty?
// 		Language["Modula-2"]
//	elsif (/^\s*import (scala|java)\./.match(data) || /^\s*val\s+\w+\s*=/.match(data) || /^\s*class\b/.match(data))
//              Language["Scala"]
//      elsif (data.include?("gap> "))
//		Language["GAP"]
// 	else
// 		[Language["Linux Kernel Module"], Language["AMPL"]]
// 	end
// end
func getDisambiguators(heuristics []byte) ([]*disambiguator, error) {
	buf := bufio.NewScanner(bytes.NewReader(heuristics))
	disambiguators := make([]*disambiguator, 0, 50)
	for buf.Scan() {
		line := buf.Text()
		if disambLine.MatchString(line) {
			d, err := parseDisambiguators(line, buf)
			if err != nil {
				return nil, err
			}

			disambiguators = append(disambiguators, d...)
		}

		lookForRegexpVariables(line)
	}

	if err := buf.Err(); err != nil {
		return nil, err
	}

	return disambiguators, nil
}

func lookForRegexpVariables(line string) {
	if strings.Contains(line, "ObjectiveCRegex = ") {
		line = strings.TrimSpace(line)
		reg := strings.TrimPrefix(line, "ObjectiveCRegex = ")
		definedRegs["ObjectiveCRegex"] = reg
	}

	if strings.Contains(line, "fortran_rx = ") {
		line = strings.TrimSpace(line)
		reg := strings.TrimPrefix(line, "fortran_rx = ")
		definedRegs["fortran_rx"] = reg
	}
}

func parseDisambiguators(line string, buf *bufio.Scanner) ([]*disambiguator, error) {
	disambList := make([]*disambiguator, 0, 2)
	splitted := strings.Fields(line)

	for _, v := range splitted {
		if strings.HasPrefix(v, `"`) {
			extension := strings.Trim(v, `",`)
			d := &disambiguator{Extension: extension}
			disambList = append(disambList, d)
		}
	}

	lh, err := getLanguagesHeuristics(buf)
	if err != nil {
		return nil, err
	}

	for _, v := range disambList {
		v.Languages = lh
	}

	return disambList, nil
}

func getLanguagesHeuristics(buf *bufio.Scanner) ([]*languageHeuristics, error) {
	langs := make([][]string, 0, 2)
	regs := make([][]string, 0, 1)
	logicRels := make([][]string, 0, 1)

	lastWasMatch := false
	for buf.Scan() {
		line := buf.Text()
		if strings.Contains(line, "end") {
			break
		}

		if hasRegExp(line) {
			line := cleanRegExpLine(line)

			lr := getLogicRelations(line)
			logicRels = append(logicRels, lr)

			r := getRegExp(line)
			if lastWasMatch {
				i := len(regs) - 1
				regs[i] = append(regs[i], r...)
			} else {
				regs = append(regs, r)
			}

			lastWasMatch = true
		}

		if strings.Contains(line, "Language") {
			l := getLanguages(line)
			langs = append(langs, l)
			lastWasMatch = false
		}

	}

	if err := buf.Err(); err != nil {
		return nil, err
	}

	langsHeuristics := buildLanguagesHeuristics(langs, regs, logicRels)
	return langsHeuristics, nil
}

func hasRegExp(line string) bool {
	return strings.Contains(line, ".match") || strings.Contains(line, ".include?") || strings.Contains(line, ".empty?")
}

func cleanRegExpLine(line string) string {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "if ")
	line = strings.TrimPrefix(line, "elsif ")
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, `(`)
	if strings.Contains(line, "))") {
		line = strings.TrimSuffix(line, `)`)
	}

	return line
}

func getLogicRelations(line string) []string {
	rels := make([]string, 0)
	splitted := strings.Split(line, "||")
	for i, v := range splitted {
		if strings.Contains(v, "&&") {
			rels = append(rels, "&&")
		}

		if i < len(splitted)-1 {
			rels = append(rels, "||")
		}
	}

	if len(rels) == 0 {
		rels = nil
	}

	return rels
}

func getRegExp(line string) []string {
	splitted := splitByLogicOps(line)
	regs := make([]string, 0, len(splitted))
	for _, v := range splitted {
		v = strings.TrimSpace(v)

		if strings.Contains(v, ".match") {
			r := v[:strings.Index(v, ".match")]
			r = replaceRegexpVariables(r)
			regs = append(regs, r)
		}

		if strings.Contains(v, ".include?") {
			r := includeToRegExp(v)
			regs = append(regs, r)
		}

		if strings.Contains(v, ".empty?") {
			r := `/^$/`
			regs = append(regs, r)
		}
	}

	return regs
}

func splitByLogicOps(line string) []string {
	splitted := make([]string, 0, 1)
	splitOr := strings.Split(line, "||")
	for _, v := range splitOr {
		splitAnd := strings.Split(v, "&&")
		splitted = append(splitted, splitAnd...)
	}

	return splitted
}

func replaceRegexpVariables(reg string) string {
	repl := reg
	if v, ok := definedRegs[reg]; ok {
		repl = v
	}

	return repl
}

func includeToRegExp(include string) string {
	content := include[strings.Index(include, `(`)+1 : strings.Index(include, `)`)]
	content = strings.Trim(content, `"'`)
	return regexp.QuoteMeta(content)
}

func getLanguages(line string) []string {
	languages := make([]string, 0)
	splitted := strings.Split(line, `,`)
	for _, lang := range splitted {
		lang = trimLanguage(lang)
		languages = append(languages, lang)
	}

	return languages
}

func trimLanguage(enclosedLang string) string {
	lang := strings.TrimSpace(enclosedLang)
	lang = lang[strings.Index(lang, `"`)+1:]
	lang = lang[:strings.Index(lang, `"`)]
	return lang
}

func buildLanguagesHeuristics(langs, regs, logicRels [][]string) []*languageHeuristics {
	langsHeuristics := make([]*languageHeuristics, 0, len(langs))
	for i, langSlice := range langs {
		var heuristics []string
		if i < len(regs) {
			heuristics = regs[i]
		}

		var rels []string
		if i < len(logicRels) {
			rels = logicRels[i]
		}

		for _, lang := range langSlice {
			lh := &languageHeuristics{
				Language:       lang,
				Heuristics:     heuristics,
				LogicRelations: rels,
			}

			langsHeuristics = append(langsHeuristics, lh)
		}
	}

	return langsHeuristics
}
