package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

const (
	YAMLFile   = "test_files/languages_test.yml"
	langGold   = "test_files/languages.gold"
	commitTest = "76816a9eb2a9fc0ba440067e1d4dd8fbc62349e5"
)

type GenerateLanguages struct {
	suite.Suite
	YAMLParsed yaml.MapSlice
	Commit     string
}

func (l *GenerateLanguages) SetupSuite() {
	f, err := os.Open(YAMLFile)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(buf, &l.YAMLParsed); err != nil {
		log.Fatal(err)
	}

	l.Commit = commitTest
}

func (l *GenerateLanguages) Test_Languages() {
	f, err := os.Open(langGold)
	if err != nil {
		assert.Fail(l.T(), err.Error())
	}

	gold, err := ioutil.ReadAll(f)
	if err != nil {
		assert.Fail(l.T(), err.Error())
	}

	f, err = os.Open(YAMLFile)
	if err != nil {
		assert.Fail(l.T(), err.Error())
	}

	input, err := ioutil.ReadAll(f)
	if err != nil {
		assert.Fail(l.T(), err.Error())
	}

	tests := []struct {
		name              string
		input             []byte
		languagesTmplPath string
		languagesTmpl     string
		commit            string
		wantOut           []byte
	}{
		{
			name:              "Test_generateLanguages",
			input:             input,
			languagesTmplPath: "test_files/languages.test.tmpl",
			languagesTmpl:     "languages.test.tmpl",
			commit:            l.Commit,
			wantOut:           gold,
		},
	}

	for _, tt := range tests {
		l.T().Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			Languages(out, tt.input, tt.languagesTmplPath, tt.languagesTmpl, tt.commit)
			assert.EqualValues(t, tt.wantOut, out.Bytes(), fmt.Sprintf("generateLanguages() = %v, want %v", out, tt.wantOut))
		})
	}
}

func (l *GenerateLanguages) Test_buildExtensionLanguageMap() {
	type args struct {
		yamlSlice yaml.MapSlice
	}

	tests := []struct {
		name    string
		args    args
		want    map[string][]string
		wantErr error
	}{
		{
			name:    "Test_buildExtensionLanguageMap",
			args:    args{yamlSlice: l.YAMLParsed},
			want:    map[string][]string{".abap": []string{"ABAP"}, ".abnf": []string{"ABNF"}, ".bsl": []string{"1C Enterprise"}, ".os": []string{"1C Enterprise"}},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		l.T().Run(tt.name, func(t *testing.T) {
			got, err := buildExtensionLanguageMap(tt.args.yamlSlice)
			assert.EqualValues(t, tt.wantErr, err, fmt.Sprintf("buildExtensionLanguageMap() error = %v, wantErr %v", err, tt.wantErr))
			assert.EqualValues(t, tt.want, got, fmt.Sprintf("buildExtensionLanguageMap() = %v, want %v", got, tt.want))
		})
	}
}

func (l *GenerateLanguages) Test_findExtensions() {
	type args struct {
		items yaml.MapSlice
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test_findExtensions",
			args: args{items: l.YAMLParsed}, want: []string{".bsl", ".os", ".abap", ".abnf"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		l.T().Run(tt.name, func(t *testing.T) {
			extensions := make([]string, 0, len(tt.want))
			for _, lang := range tt.args.items {
				got, err := findExtensions(lang.Value.(yaml.MapSlice))
				if err != nil {
					assert.True(t, tt.wantErr, fmt.Sprintf("findExtensions() error = %v, wantErr %v", err, tt.wantErr))
				}

				extensions = append(extensions, got...)
			}

			assert.EqualValues(t, tt.want, extensions, fmt.Sprintf("findExtensions() = %v, want %v", extensions, tt.want))
		})
	}
}

func (l *GenerateLanguages) Test_toStringSlice() {
	type args struct {
		slice []interface{}
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test_toStringSlice",
			args: args{slice: []interface{}{"AGS Script", "AsciiDoc", "Public Key"}},
			want: []string{"AGS Script", "AsciiDoc", "Public Key"},
		},
	}

	for _, tt := range tests {
		l.T().Run(tt.name, func(t *testing.T) {
			got := toStringSlice(tt.args.slice)
			assert.EqualValues(t, tt.want, got, fmt.Sprintf("toStringSlice() = %v, want %v", got, tt.want))
		})
	}
}

func (l *GenerateLanguages) Test_fillMap() {
	type args struct {
		mapExt     map[string][]string
		lang       string
		extensions []string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "Test_fillMap", args: args{mapExt: make(map[string][]string), lang: "testlang", extensions: []string{".ex1", ".ex2", "ex3"}}},
	}

	for _, tt := range tests {
		l.T().Run(tt.name, func(t *testing.T) {
			fillMap(tt.args.mapExt, tt.args.lang, tt.args.extensions)
			assert.Len(t, tt.args.mapExt, len(tt.args.extensions), "the map doesn't contains the exact number of extensions")
			for _, v := range tt.args.extensions {
				langs, ok := tt.args.mapExt[v]
				assert.True(t, ok, fmt.Sprintf("the extension %v is not in the map", v))
				assert.Contains(t, langs, tt.args.lang, "languages don't match")
			}
		})
	}
}

func TestGenerateLanguages(t *testing.T) {
	suite.Run(t, new(GenerateLanguages))
}
