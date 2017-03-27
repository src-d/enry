package generator

import "io"

// Func is the function's type that generate the files from templates.
type Func func(out io.Writer, dataToParse []byte, templatePath string, template string, commit string) error
