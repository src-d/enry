package generate

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
)

type languageHeuristics struct {
	Language   string
	Heuristics []string
}

type disambiguator struct {
	Extension string
	Languages []*languageHeuristics
}

// Heuristics read from buf and builds content.go file from contentTmplPath.
func Heuristics(out io.Writer, buf []byte, contentTmplPath, contentTmpl, commit string) error {
	heuristics := bufio.NewScanner(bytes.NewReader(buf))
	disambiguators := getDisambiguators(heuristics)

	// debug
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	if err := enc.Encode(disambiguators); err != nil {
		return err
	}

	return nil
}

// A disambiguate block looks like:
// disambiguate ".mod", ".extension" do |data|
// 	if data.include?('<!ENTITY ') && data.include?('patata')
// 		Language["XML"]
// 	elsif /^\s*MODULE [\w\.]+;/i.match(data) || /^\s*END [\w\.]+;/i.match(data) || data.empty?
// 		Language["Modula-2"]
// 	else
// 		[Language["Linux Kernel Module"], Language["AMPL"]]
// 	end
// end
func getDisambiguators(buf *bufio.Scanner) []*disambiguator {
	disambLine := regexp.MustCompile(`^(\s*)disambiguate`)
	disambiguators := make([]*disambiguator, 0, 50)
	for buf.Scan() {
		line := buf.Text()
		if disambLine.MatchString(line) {
			disambiguators = append(disambiguators, parseDisambiguators(line, buf)...)
		}

	}

	return disambiguators
}

func parseDisambiguators(line string, buf *bufio.Scanner) []*disambiguator {
	disambList := make([]*disambiguator, 0, 2)
	splitted := strings.Fields(line)

	// get the extensions associated to this disambiguate block
	for _, v := range splitted {
		if strings.HasPrefix(v, `"`) {
			extension := strings.Trim(v, `",`)
			d := &disambiguator{Extension: extension}
			disambList = append(disambList, d)
		}
	}

	lh := getLanguagesHeuristics(buf)
	for _, v := range disambList {
		v.Languages = lh
	}

	return disambList
}

func getLanguagesHeuristics(buf *bufio.Scanner) []*languageHeuristics {
	langs := make([][]string, 0, 2)
	regs := make([][]string, 0, 1)

	lastWasMatch := false
	for buf.Scan() {
		line := buf.Text()
		if strings.Contains(line, "end") {
			break
		}

		if hasRegExp(line) {
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
		}
	}

	langsHeuristics := buildLanguagesHeuristics(langs, regs)
	return langsHeuristics
}

func hasRegExp(line string) bool {
	return strings.Contains(line, ".match") || strings.Contains(line, ".include?") || strings.Contains(line, ".empty?")
}

func getRegExp(line string) []string {
	splitted := cleanRegExpLine(line)
	regs := make([]string, 0, len(splitted))
	for _, v := range splitted {
		v = strings.TrimSpace(v)

		if strings.Contains(v, ".match") {
			r := v[:strings.Index(v, ".match")]
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

func cleanRegExpLine(line string) []string {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "if ")
	line = strings.TrimPrefix(line, "elsif ")
	splitted := strings.Split(line, "||")

	return splitted
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

func buildLanguagesHeuristics(langs, regs [][]string) []*languageHeuristics {
	langsHeuristics := make([]*languageHeuristics, 0, len(langs))
	for i, langSlice := range langs {
		var heuristics []string
		if i < len(regs) {
			heuristics = regs[i]
		}

		for _, lang := range langSlice {
			lh := &languageHeuristics{
				Language:   lang,
				Heuristics: heuristics,
			}

			langsHeuristics = append(langsHeuristics, lh)
		}
	}

	return langsHeuristics
}
