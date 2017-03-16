package main

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
	YAMLFile = "test_files/languages_test.yml"
	langGold = "test_files/languages.gold"
)

type GenerateLanguages struct {
	suite.Suite
	YAMLParsed yaml.MapSlice
}

func (sgl *GenerateLanguages) SetupSuite() {
	f, err := os.Open(YAMLFile)
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(buf, &sgl.YAMLParsed); err != nil {
		log.Fatal(err)
	}
}

func (sgl *GenerateLanguages) Test_generateLanguages() {
	f, err := os.Open(langGold)
	if err != nil {
		assert.Fail(sgl.T(), err.Error())
	}

	gold, err := ioutil.ReadAll(f)
	if err != nil {
		assert.Fail(sgl.T(), err.Error())
	}

	tests := []struct {
		name    string
		wantOut []byte
	}{
		{name: "Test_generateLanguages", wantOut: gold},
	}

	for _, tt := range tests {
		sgl.T().Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			generateLanguages(out, tmplName, tmplName)
			assert.EqualValues(t, gold, out.Bytes(), fmt.Sprintf("generateLanguages() = %v, want %v", out, tt.wantOut))
		})
	}
}

func (sgl *GenerateLanguages) Test_findExtensions() {
	type args struct {
		items yaml.MapSlice
		key   string
	}

	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "Test_findExtensions",
			args: args{items: sgl.YAMLParsed, key: extField}, want: []interface{}{".bsl", ".os", ".abap", ".abnf"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		sgl.T().Run(tt.name, func(t *testing.T) {
			extensions := make([]interface{}, 0, len(tt.want))
			for _, lang := range tt.args.items {
				got, err := findExtensions(lang.Value.(yaml.MapSlice), tt.args.key)
				if err != nil {
					assert.True(t, tt.wantErr, fmt.Sprintf("findExtensions() error = %v, wantErr %v", err, tt.wantErr))
				}

				extensions = append(extensions, got...)
			}

			assert.EqualValues(t, tt.want, extensions, fmt.Sprintf("findExtensions() = %v, want %v", extensions, tt.want))
		})
	}
}

func (sgl *GenerateLanguages) Test_fillMap() {
	type args struct {
		mapExt     map[string][]string
		lang       string
		extensions []interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "Test_fillMap", args: args{mapExt: make(map[string][]string), lang: "testlang", extensions: []interface{}{".ex1", ".ex2", "ex3"}}},
	}

	for _, tt := range tests {
		sgl.T().Run(tt.name, func(t *testing.T) {
			fillMap(tt.args.mapExt, tt.args.lang, tt.args.extensions)
			assert.Len(t, tt.args.mapExt, len(tt.args.extensions), "the map doesn't contains the exact number of extensions")
			for _, v := range tt.args.extensions {
				langs, ok := tt.args.mapExt[v.(string)]
				assert.True(t, ok, fmt.Sprintf("the extension %v is not in the map", v))
				assert.Contains(t, langs, tt.args.lang, "languages don't match")
			}
		})
	}
}

func (sgl *GenerateLanguages) Test_formatStringSlice() {
	type args struct {
		slice []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test_formatStringSlice", args: args{slice: []string{"hello", "world", "!!!"}}, want: `"hello","world","!!!"`},
	}

	for _, tt := range tests {
		sgl.T().Run(tt.name, func(t *testing.T) {
			got := formatStringSlice(tt.args.slice)
			assert.Equal(t, tt.want, got, fmt.Sprintf("formatStringSlice() = %v, want %v", got, tt.want))
		})
	}
}

func TestGenerateLanguages(t *testing.T) {
	suite.Run(t, new(GenerateLanguages))
}
