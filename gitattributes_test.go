package enry

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/assert"
)

func (s *EnryTestSuite) TestLoadGitAttributes() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	data := []byte("path linguist-vendored\n path/foo linguist-vendored=false\n path/vendor linguist-vendored=false \n path/foo linguist-documentation\n path/generated linguist-generated\n" +
		"path/bar linguist-vendored=fail\n path/foo linguist-documentation=false\n path/bar not-a-matcher\n path/a linguist-documentation linguist-vendored")
	tmpGitAttributes.Write(data)
	tmpGitAttributes.Close()
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("test/", reader)
	if len(errArr) != 3 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}

	tests := []struct {
		name     string
		expected int
	}{
		{name: "TestLoadGitAttributes_1", expected: 3},
		{name: "TestLoadGitAttributes_2", expected: 1},
		{name: "TestLoadGitAttributes_3", expected: 1},
		{name: "TestLoadGitAttributes_4", expected: 0},
	}

	for i, test := range tests {
		if attrType(i) < language {
			assert.Equal(s.T(), len(gitAttrs.boolAttributes[attrType(i)].attributes), test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, len(gitAttrs.boolAttributes[attrType(i)].attributes), test.expected))
		} else {
			assert.Equal(s.T(), len(gitAttrs.regExpAttributes[attrType(i)].attributes), test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, len(gitAttrs.regExpAttributes[attrType(i)].attributes), test.expected))
		}
	}

	err = os.RemoveAll(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
}

func (s *EnryTestSuite) TestLoadGitAttributesEmpty() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("test/", reader)
	if len(errArr) != 0 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}
}

func (s *EnryTestSuite) TestIsVendorGitAttributes() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	data := []byte("path linguist-vendored\n path/foo linguist-vendored=false\n path/vendor linguist-vendored=false")
	tmpGitAttributes.Write(data)
	tmpGitAttributes.Close()
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("", reader)
	if len(errArr) != 0 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{name: "TestIsVendorGitAttributes_1", path: "path", expected: true},
		{name: "TestIsVendorGitAttributes_2", path: "path/foo", expected: false},
		{name: "TestIsVendorGitAttributes_3", path: "path/vendor", expected: false},
		{name: "TestIsVendorGitAttributes_4", path: "vendor/", expected: true},
		{name: "TestIsVendorGitAttributes_5", path: "dir/", expected: false},
	}
	for _, test := range tests {
		is := gitAttrs.IsVendor(test.path)
		assert.Equal(s.T(), is, test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, is, test.expected))
	}

	err = os.RemoveAll(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
}

func (s *EnryTestSuite) TestIsDocumentationGitAttributes() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	data := []byte("path linguist-documentation\n path/foo linguist-documentation=false\n path/documentation linguist-vendored=false")
	tmpGitAttributes.Write(data)
	tmpGitAttributes.Close()
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("", reader)
	if len(errArr) != 0 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{name: "TestIsDocumentationGitAttributes_1", path: "path", expected: true},
		{name: "TestIsDocumentationGitAttributes_2", path: "path/foo", expected: false},
		{name: "TestIsDocumentationGitAttributes_3", path: "path/documentation", expected: false},
		{name: "TestIsDocumentationGitAttributes_4", path: "README", expected: true},
		{name: "TestIsDocumentationGitAttributes_5", path: "dir/", expected: false},
	}
	for _, test := range tests {
		is := gitAttrs.IsDocumentation(test.path)
		assert.Equal(s.T(), is, test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, is, test.expected))
	}

	err = os.RemoveAll(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
}

func (s *EnryTestSuite) TestIsGeneratedGitAttributes() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	data := []byte("path linguist-generated\n path/foo linguist-generated=false\n path/generated linguist-generated=false")
	tmpGitAttributes.Write(data)
	tmpGitAttributes.Close()
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("", reader)
	if len(errArr) != 0 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{name: "TestIsGeneratedGitAttributes_1", path: "path", expected: true},
		{name: "TestIsGeneratedGitAttributes_2", path: "path/foo", expected: false},
		{name: "TestIsGeneratedGitAttributes_3", path: "path/generated", expected: false},
		{name: "TestIsGeneratedGitAttributes_4", path: "path2", expected: false},
	}
	for _, test := range tests {
		is := gitAttrs.IsGenerated(test.path)
		assert.Equal(s.T(), is, test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, is, test.expected))
	}

	err = os.RemoveAll(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
}

func (s *EnryTestSuite) TestGetLanguageGitAttributes() {
	gitAttrs := NewGitAttributes()
	tmpGitAttributes, err := ioutil.TempFile("/tmp", "gitattributes")
	assert.NoError(s.T(), err)
	data := []byte(".*\\.go linguist-language=GO\n path/not-java/.*\\.java linguist-language=notJava\n")
	tmpGitAttributes.Write(data)
	tmpGitAttributes.Close()
	reader, err := os.Open(tmpGitAttributes.Name())
	assert.NoError(s.T(), err)
	errArr := gitAttrs.LoadGitAttributes("", reader)
	if len(errArr) != 0 {
		fmt.Println(errArr)
		s.Fail("The error length it's not the expected")
	}

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{name: "TestGetLanguageGitAttributes_1", path: "path/files/a.go", expected: "Go"},
		{name: "TestGetLanguageGitAttributes_2", path: "path/files/subdir/b.go", expected: "Go"},
		{name: "TestGetLanguageGitAttributes_3", path: "path/not-java/c.java", expected: "notJava"},
		{name: "TestGetLanguageGitAttributes_4", path: "path/d.py", expected: ""},
	}

	for _, test := range tests {
		is := gitAttrs.GetLanguage(test.path)
		assert.Equal(s.T(), is, test.expected, fmt.Sprintf("%v: is = %v, expected: %v", test.name, is, test.expected))
	}
}
