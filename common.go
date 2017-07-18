package enry

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/src-d/enry.v1/data"
)

// OtherLanguage is used as a zero value when a function can not return a specific language.
const OtherLanguage = ""

// Strategy type fix the signature for the functions that can be used as a strategy.
type Strategy func(filename string, content []byte, candidates []string) (languages []string)

// DefaultStrategies is the strategies' sequence GetLanguage uses to detect languages.
var DefaultStrategies = []Strategy{
	GetLanguagesByModeline,
	GetLanguagesByFilename,
	GetLanguagesByShebang,
	GetLanguagesByExtension,
	GetLanguagesByContent,
	GetLanguagesByClassifier,
}

var DefaultClassifier Classifier = &classifier{
	languagesLogProbabilities: data.LanguagesLogProbabilities,
	tokensLogProbabilities:    data.TokensLogProbabilities,
	tokensTotal:               data.TokensTotal,
}

// GetLanguage applies a sequence of strategies based on the given filename and content
// to find out the most probably language to return.
func GetLanguage(filename string, content []byte) (language string) {
	languages := GetLanguages(filename, content)
	return firstLanguage(languages)
}

func firstLanguage(languages []string) string {
	if len(languages) == 0 {
		return OtherLanguage
	}

	return languages[0]
}

// GetLanguageByModeline returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByModeline(content []byte) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByModeline, "", content, nil)
}

// GetLanguageByEmacsModeline returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByEmacsModeline(content []byte) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByEmacsModeline, "", content, nil)
}

// GetLanguageByVimModeline returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByVimModeline(content []byte) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByVimModeline, "", content, nil)
}

// GetLanguageByFilename returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByFilename(filename string) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByFilename, filename, nil, nil)
}

// GetLanguageByShebang returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByShebang(content []byte) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByShebang, "", content, nil)
}

// GetLanguageByExtension returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByExtension(filename string) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByExtension, filename, nil, nil)
}

// GetLanguageByContent returns detected language. If there are more than one possibles languages
// it returns the first language by alphabetically order and safe to false.
func GetLanguageByContent(content []byte) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByContent, "", content, nil)
}

// GetLanguageByClassifier returns the most probably language detected for the given content. It uses
// DefaultClassifier, if no candidates are provided it returns OtherLanguage.
func GetLanguageByClassifier(content []byte, candidates []string) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByClassifier, "", content, candidates)
}

// GetLanguageByGitattributes returns the language assigned to a file for a given regular expresion in .gitattributes.
// This strategy needs to be initialized calling LoadGitattributes
func GetLanguageByGitattributes(filename string) (language string, safe bool) {
	return getLanguageByStrategy(GetLanguagesByGitAttributes, filename, nil, nil)
}

func getLanguageByStrategy(strategy Strategy, filename string, content []byte, candidates []string) (string, bool) {
	languages := strategy(filename, content, candidates)
	return getFirstLanguageAndSafe(languages)
}

func getFirstLanguageAndSafe(languages []string) (language string, safe bool) {
	language = firstLanguage(languages)
	safe = len(languages) == 1
	return
}

// GetLanguageBySpecificClassifier returns the most probably language for the given content using
// classifier to detect language.
func GetLanguageBySpecificClassifier(content []byte, candidates []string, classifier Classifier) (language string, safe bool) {
	languages := GetLanguagesBySpecificClassifier(content, candidates, classifier)
	return getFirstLanguageAndSafe(languages)
}

// GetLanguages applies a sequence of strategies based on the given filename and content
// to find out the most probably languages to return.
func GetLanguages(filename string, content []byte) []string {
	if IsBinary(content) {
		return nil
	}

	var languages []string
	candidates := []string{}
	for _, strategy := range DefaultStrategies {
		languages = strategy(filename, content, candidates)
		if len(languages) == 1 {
			return languages
		}

		if len(languages) > 0 {
			candidates = append(candidates, languages...)
		}
	}

	return languages
}

// GetLanguagesByModeline returns a slice of possible languages for the given content, filename will be ignored.
// It complies with the signature to be a Strategy type.
func GetLanguagesByModeline(filename string, content []byte, candidates []string) []string {
	headFoot := getHeaderAndFooter(content)
	var languages []string
	for _, getLang := range modelinesFunc {
		languages = getLang("", headFoot, candidates)
		if len(languages) > 0 {
			break
		}
	}

	return languages
}

var modelinesFunc = []Strategy{
	GetLanguagesByEmacsModeline,
	GetLanguagesByVimModeline,
}

func getHeaderAndFooter(content []byte) []byte {
	const searchScope = 5
	if bytes.Count(content, []byte("\n")) < 2*searchScope {
		return content
	}

	header := headScope(content, searchScope)
	footer := footScope(content, searchScope)
	headerAndFooter := make([]byte, 0, len(content[:header])+len(content[footer:]))
	headerAndFooter = append(headerAndFooter, content[:header]...)
	headerAndFooter = append(headerAndFooter, content[footer:]...)
	return headerAndFooter
}

func headScope(content []byte, scope int) (index int) {
	for i := 0; i < scope; i++ {
		eol := bytes.IndexAny(content, "\n")
		content = content[eol+1:]
		index += eol
	}

	return index + scope - 1
}

func footScope(content []byte, scope int) (index int) {
	for i := 0; i < scope; i++ {
		index = bytes.LastIndexAny(content, "\n")
		content = content[:index]
	}

	return index + 1
}

var (
	reEmacsModeline = regexp.MustCompile(`.*-\*-\s*(.+?)\s*-\*-.*(?m:$)`)
	reEmacsLang     = regexp.MustCompile(`.*(?i:mode)\s*:\s*([^\s;]+)\s*;*.*`)
	reVimModeline   = regexp.MustCompile(`(?:(?m:\s|^)vi(?:m[<=>]?\d+|m)?|[\t\x20]*ex)\s*[:]\s*(.*)(?m:$)`)
	reVimLang       = regexp.MustCompile(`(?i:filetype|ft|syntax)\s*=(\w+)(?:\s|:|$)`)
)

// GetLanguagesByEmacsModeline returns a slice of possible languages for the given content, filename and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByEmacsModeline(filename string, content []byte, candidates []string) []string {
	matched := reEmacsModeline.FindAllSubmatch(content, -1)
	if matched == nil {
		return nil
	}

	// only take the last matched line, discard previous lines
	lastLineMatched := matched[len(matched)-1][1]
	matchedAlias := reEmacsLang.FindSubmatch(lastLineMatched)
	var alias string
	if matchedAlias != nil {
		alias = string(matchedAlias[1])
	} else {
		alias = string(lastLineMatched)
	}

	language, ok := GetLanguageByAlias(alias)
	if !ok {
		return nil
	}

	return []string{language}
}

// GetLanguagesByVimModeline returns a slice of possible languages for the given content, filename and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByVimModeline(filename string, content []byte, candidates []string) []string {
	matched := reVimModeline.FindAllSubmatch(content, -1)
	if matched == nil {
		return nil
	}

	// only take the last matched line, discard previous lines
	lastLineMatched := matched[len(matched)-1][1]
	matchedAlias := reVimLang.FindAllSubmatch(lastLineMatched, -1)
	if matchedAlias == nil {
		return nil
	}

	alias := string(matchedAlias[0][1])
	if len(matchedAlias) > 1 {
		// cases:
		// matchedAlias = [["syntax=ruby " "ruby"] ["ft=python " "python"] ["filetype=perl " "perl"]] returns OtherLanguage;
		// matchedAlias = [["syntax=python " "python"] ["ft=python " "python"] ["filetype=python " "python"]] returns "Python";
		for _, match := range matchedAlias {
			otherAlias := string(match[1])
			if otherAlias != alias {
				return nil
			}
		}
	}

	language, ok := GetLanguageByAlias(alias)
	if !ok {
		return nil
	}

	return []string{language}
}

// GetLanguagesByFilename returns a slice of possible languages for the given filename, content and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByFilename(filename string, content []byte, candidates []string) []string {
	return data.LanguagesByFilename[filepath.Base(filename)]
}

// GetLanguagesByShebang returns a slice of possible languages for the given content, filename and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByShebang(filename string, content []byte, candidates []string) (languages []string) {
	interpreter := getInterpreter(content)
	return data.LanguagesByInterpreter[interpreter]
}

var (
	shebangExecHack = regexp.MustCompile(`exec (\w+).+\$0.+\$@`)
	pythonVersion   = regexp.MustCompile(`python\d\.\d+`)
)

func getInterpreter(data []byte) (interpreter string) {
	line := getFirstLine(data)
	if !hasShebang(line) {
		return ""
	}

	// skip shebang
	line = bytes.TrimSpace(line[2:])

	splitted := bytes.Fields(line)
	if bytes.Contains(splitted[0], []byte("env")) {
		if len(splitted) > 1 {
			interpreter = string(splitted[1])
		}
	} else {

		splittedPath := bytes.Split(splitted[0], []byte{'/'})
		interpreter = string(splittedPath[len(splittedPath)-1])
	}

	if interpreter == "sh" {
		interpreter = lookForMultilineExec(data)
	}

	if pythonVersion.MatchString(interpreter) {
		interpreter = interpreter[:strings.Index(interpreter, `.`)]
	}

	return
}

func getFirstLine(data []byte) []byte {
	buf := bufio.NewScanner(bytes.NewReader(data))
	buf.Scan()
	line := buf.Bytes()
	if err := buf.Err(); err != nil {
		return nil
	}

	return line
}

func hasShebang(line []byte) bool {
	const shebang = `#!`
	prefix := []byte(shebang)
	return bytes.HasPrefix(line, prefix)
}

func lookForMultilineExec(data []byte) string {
	const magicNumOfLines = 5
	interpreter := "sh"

	buf := bufio.NewScanner(bytes.NewReader(data))
	for i := 0; i < magicNumOfLines && buf.Scan(); i++ {
		line := buf.Bytes()
		if shebangExecHack.Match(line) {
			interpreter = shebangExecHack.FindStringSubmatch(string(line))[1]
			break
		}
	}

	if err := buf.Err(); err != nil {
		return interpreter
	}

	return interpreter
}

// GetLanguagesByExtension returns a slice of possible languages for the given filename, content and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByExtension(filename string, content []byte, candidates []string) []string {
	if !strings.Contains(filename, ".") {
		return nil
	}

	filename = strings.ToLower(filename)
	dots := getDotIndexes(filename)
	for _, dot := range dots {
		ext := filename[dot:]
		languages, ok := data.LanguagesByExtension[ext]
		if ok {
			return languages
		}
	}

	return nil
}

func getDotIndexes(filename string) []int {
	dots := make([]int, 0, 2)
	for i, letter := range filename {
		if letter == rune('.') {
			dots = append(dots, i)
		}
	}

	return dots
}

// GetLanguagesByContent returns a slice of possible languages for the given content, filename and candidates
// will be ignored. It complies with the signature to be a Strategy type.
func GetLanguagesByContent(filename string, content []byte, candidates []string) []string {
	ext := strings.ToLower(filepath.Ext(filename))
	fnMatcher, ok := data.ContentMatchers[ext]
	if !ok {
		return nil
	}

	return fnMatcher(content)
}

// GetLanguagesByClassifier uses DefaultClassifier as a Classifier and returns a sorted slice of possible languages ordered by
// decreasing language's probability. If there are not candidates it returns nil. It complies with the signature to be a Strategy type.
func GetLanguagesByClassifier(filename string, content []byte, candidates []string) (languages []string) {
	if len(candidates) == 0 {
		return nil
	}

	return GetLanguagesBySpecificClassifier(content, candidates, DefaultClassifier)
}

// GetLanguagesBySpecificClassifier returns a slice of possible languages. It takes in a Classifier to be used.
func GetLanguagesBySpecificClassifier(content []byte, candidates []string, classifier Classifier) (languages []string) {
	mapCandidates := make(map[string]float64)
	for _, candidate := range candidates {
		mapCandidates[candidate]++
	}

	return classifier.Classify(content, mapCandidates)
}

// GetLanguagesByGitAttributes returns either a string slice with the language
// if the filename matches with a regExp in .gitattributes or returns a empty slice
// in case no regExp matches the filename. It complies with the signature to be a Strategy type.
func GetLanguagesByGitAttributes(filename string, content []byte, candidates []string) []string {
	gitAttributes := NewGitAttributes()
	reader, err := os.Open(".gitattributes")
	if err != nil {
		return nil
	}

	gitAttributes.LoadGitAttributes("", reader)
	lang := gitAttributes.GetLanguage(filename)
	if lang != OtherLanguage {
		return []string{}
	}

	return []string{lang}
}

// GetLanguageExtensions returns the different extensions being used by the language.
func GetLanguageExtensions(language string) []string {
	return data.ExtensionsByLanguage[language]
}

// Type represent language's type. Either data, programming, markup, prose, or unknown.
type Type int

// Type's values.
const (
	Unknown Type = iota
	Data
	Programming
	Markup
	Prose
)

// GetLanguageType returns the type of the given language.
func GetLanguageType(language string) (langType Type) {
	intType, ok := data.LanguagesType[language]
	langType = Type(intType)
	if !ok {
		langType = Unknown
	}
	return langType
}

// GetLanguageByAlias returns either the language related to the given alias and ok set to true
// or Otherlanguage and ok set to false if the alias is not recognized.
func GetLanguageByAlias(alias string) (lang string, ok bool) {
	a := strings.Split(alias, `,`)[0]
	a = strings.ToLower(a)
	lang, ok = data.LanguagesByAlias[a]
	if !ok {
		lang = OtherLanguage
	}

	return
}
