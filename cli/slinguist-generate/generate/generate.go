package generate

import "io"

// GeneratorFunc is the function's type that generate the files from templates.
type GeneratorFunc func(out io.Writer, dataToParse []byte, templatePath string, template string, commit string) error
