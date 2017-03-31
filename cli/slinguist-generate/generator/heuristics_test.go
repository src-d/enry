package generator

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	heuristicsTestFile   = "test_files/heuristics.test.rb"
	contentGold          = "test_files/content.gold"
	contentTestTmplPath  = "test_files/content.test.go.tmpl"
	contentTestTmplName  = "content.test.go.tmpl"
	commitHeuristicsTest = "fe8b44ab8a225b1ffa75b983b916ea22fee5b6f7"
)

func TestHeuristics(t *testing.T) {
	gold, err := ioutil.ReadFile(contentGold)
	assert.NoError(t, err)

	input, err := ioutil.ReadFile(heuristicsTestFile)
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
			name:     "TestHeuristics",
			input:    input,
			tmplPath: contentTestTmplPath,
			tmplName: contentTestTmplName,
			commit:   commitHeuristicsTest,
			wantOut:  gold,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Heuristics(tt.input, tt.tmplPath, tt.tmplName, tt.commit)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.wantOut, out, fmt.Sprintf("Heuristics() = %v, want %v", string(out), string(tt.wantOut)))
		})
	}
}
