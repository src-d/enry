package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/enry.v1"
)

var (
	Version = "undefined"
	GitHash = "undefined"
)

func main() {
	args, flags := splitArgs()
	if len(args) == 1 {
		args = append(args, ".")
	}

	jsonFlag, breakdownFlag := parseFlags(flags)
	root, err := filepath.Abs(args[1])
	if err != nil {
		usage()
		return
	}

	errors := false
	out := make(map[string][]string, 0)
	err = filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			errors = true
			usage()
			return filepath.SkipDir
		}

		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			errors = true
			log.Println(err)
			return nil
		}

		if relativePath == "." {
			return nil
		}

		if f.IsDir() {
			relativePath = relativePath + "/"
		}

		if enry.IsVendor(relativePath) || enry.IsDotFile(relativePath) ||
			enry.IsDocumentation(relativePath) || enry.IsConfiguration(relativePath) {
			if f.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if f.IsDir() {
			return nil
		}

		language, ok := enry.GetLanguageByExtension(path)
		if !ok {
			if language, ok = enry.GetLanguageByFilename(path); !ok {
				content, err := ioutil.ReadFile(path)
				if err != nil {
					errors = true
					log.Println(err)
					return nil
				}

				language = enry.GetLanguage(filepath.Base(path), content)
				if language == enry.OtherLanguage {
					return nil
				}
			}
		}

		out[language] = append(out[language], relativePath)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var buff bytes.Buffer
	switch {
	case jsonFlag && !breakdownFlag:
		printJson(out, &buff)
	case jsonFlag && breakdownFlag:
		printBreakDown(out, &buff)
	case breakdownFlag:
		printPercents(out, &buff)
		buff.WriteByte('\n')
		printBreakDown(out, &buff)
	default:
		printPercents(out, &buff)
	}

	fmt.Print(buff.String())
}

func usage() {
	fmt.Fprintf(
		os.Stderr,
		`  %[1]s %[2]s commit: %[3]s
  enry, A simple (and faster) implementation of github/linguist
  usage: %[1]s <path>
         %[1]s [-json] [-breakdown] <path>
         %[1]s [-json] [-breakdown]
`,
		os.Args[0], Version, GitHash,
	)
}

func printBreakDown(out map[string][]string, buff *bytes.Buffer) {
	for name, language := range out {
		writeStringLn(name, buff)
		for _, file := range language {
			writeStringLn(file, buff)
		}

		writeStringLn("", buff)
	}
}

func printJson(out map[string][]string, buff *bytes.Buffer) {
	data, _ := json.Marshal(out)
	buff.Write(data)
	buff.WriteByte('\n')
}

func printPercents(out map[string][]string, buff *bytes.Buffer) {
	fileCount := make(map[string]int, len(out))
	total := 0
	for name, language := range out {
		fileCount[name] = len(language)
		total += len(language)
	}

	for name, count := range fileCount {
		percent := float32(count) / float32(total) * 100
		buff.WriteString(fmt.Sprintf("%.2f%%	%s\n", percent, name))
	}
}

func writeStringLn(s string, buff *bytes.Buffer) {
	buff.WriteString(s)
	buff.WriteByte('\n')
}

func splitArgs() (args []string, flagArgs []string) {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			flagArgs = append(flagArgs, arg)
		} else {
			args = append(args, arg)
		}
	}

	return
}

func parseFlags(flags []string) (jsonFlag bool, breakdownFlag bool) {
	for _, flag := range flags {
		switch {
		case "--json" == flag:
			jsonFlag = true
		case "--breakdown" == flag:
			breakdownFlag = true
		}
	}

	return
}
