package generator

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ymlTestFile           = "test_files/languages.test.yml"
	langGold              = "test_files/languages.gold"
	languagesTestTmplPath = "test_files/languages.test.tmpl"
	languagesTestTmplName = "languages.test.tmpl"
	commitLangTest        = "fe8b44ab8a225b1ffa75b983b916ea22fee5b6f7"
)

func TestLanguages(t *testing.T) {
	gold, err := ioutil.ReadFile(langGold)
	assert.NoError(t, err)

	input, err := ioutil.ReadFile(ymlTestFile)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		input    []byte
		tmplPath string
		tmplName string
		commit   string
		wantOut  []byte
	}{
		{
			name:     "TestLanguages",
			input:    input,
			tmplPath: languagesTestTmplPath,
			tmplName: languagesTestTmplName,
			commit:   commitLangTest,
			wantOut:  gold,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Languages(tt.input, tt.tmplPath, tt.tmplName, tt.commit)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.wantOut, out, fmt.Sprintf("Languages() = %v, want %v", string(out), string(tt.wantOut)))
		})
	}
}
