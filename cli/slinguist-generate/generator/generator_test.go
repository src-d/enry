package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	formatedLangGold    = "test_files/formated_languages.gold"
	formatedContentGold = "test_files/formated_content.gold"
)

func TestFromFile(t *testing.T) {
	goldLang, err := ioutil.ReadFile(formatedLangGold)
	assert.NoError(t, err)

	goldContent, err := ioutil.ReadFile(formatedContentGold)
	assert.NoError(t, err)

	outPathLang, err := ioutil.TempFile("/tmp", "generator-test-")
	assert.NoError(t, err)
	defer os.Remove(outPathLang.Name())

	outPathContent, err := ioutil.TempFile("/tmp", "generator-test-")
	assert.NoError(t, err)
	defer os.Remove(outPathContent.Name())

	tests := []struct {
		name        string
		fileToParse string
		outPath     string
		tmplPath    string
		tmplName    string
		commit      string
		generate    Func
		wantOut     []byte
	}{
		{
			name:        "TestFromFile_Language",
			fileToParse: ymlTestFile,
			outPath:     outPathLang.Name(),
			tmplPath:    languagesTestTmplPath,
			tmplName:    languagesTestTmplName,
			commit:      commitLangTest,
			generate:    Languages,
			wantOut:     goldLang,
		},
		{
			name:        "TestFromFile_Heuristics",
			fileToParse: heuristicsTestFile,
			outPath:     outPathContent.Name(),
			tmplPath:    contentTestTmplPath,
			tmplName:    contentTestTmplName,
			commit:      commitHeuristicsTest,
			generate:    Heuristics,
			wantOut:     goldContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FromFile(tt.fileToParse, tt.outPath, tt.tmplPath, tt.tmplName, tt.commit, tt.generate)
			assert.NoError(t, err)
			out, err := ioutil.ReadFile(tt.outPath)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.wantOut, out, fmt.Sprintf("FromFile() = %v, want %v", string(out), string(tt.wantOut)))
		})
	}
}
