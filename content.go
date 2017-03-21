package slinguist

import (
	"io"
	"path/filepath"
	"strings"

	"gopkg.in/toqueteos/substring.v1"
)

func GetLanguageByContent(filename string, content []byte) (lang string, safe bool) {
	ext := strings.ToLower(filepath.Ext(filename))
	if fnMatcher, ok := matchers[ext]; ok {
		println(fnMatcher)
		// lang, safe = fnMatcher(content)
		return
	}

	return GetLanguageByExtension(filename)
}

type languageMatcher func(io.RuneReader) []string

var matchers = map[string]languageMatcher{
// ".bf": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`(fprintf|function|return)`).Match(i) {
// 		return "HyPhy", true
// 	}

// 	return "Brainfuck", false
// },
// ".b": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`(include|modules)`).Match(i) {
// 		return "Limbo", true
// 	}

// 	return "Brainfuck", false
// },
// ".bb": func(i io.RuneReader) []string {
// 	if blitzBasicMatcher.Match(i) {
// 		return "BlitzBasic", true
// 	} else if substring.BytesRegexp(`^\s*(# |include|require)\b`).Match(i) {
// 		return "BitBake", true
// 	}

// 	return OtherLanguage, false
// },
// ".cl": func(i io.RuneReader) []string {
// 	if commonLispMatcher.Match(i) {
// 		return "Common Lisp", true
// 	} else if coolMatcher.Match(i) {
// 		return "Cool", true
// 	} else if openCLMatcher.Match(i) {
// 		return "OpenCL", true
// 	}

// 	return OtherLanguage, false
// },
// ".cls": func(i io.RuneReader) []string {
// 	if apexMatcher.Match(i) {
// 		return "Apex", true
// 	} else if openEdgeABLMatcher.Match(i) {
// 		return "OpenEdge ABL", true
// 	} else if texMatcher.Match(i) {
// 		return "TeX", true
// 	} else if visualBasicMatcher.Match(i) {
// 		return "Visual Basic", true

// 	}

// 	return OtherLanguage, false
// },
// ".cs": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`![\w\s]+methodsFor: `).Match(i) {
// 		return "Smalltalk", true
// 	}

// 	return "C#", true
// },
// ".ch": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`(?i)^\s*#\s*(if|ifdef|ifndef|define|command|xcommand|translate|xtranslate|include|pragma|undef)\b`).Match(i) {
// 		return "xBase", true
// 	}

// 	return "Charity", true
// },
// ".d": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^module `).Match(i) {
// 		return "D", true
// 	} else if substring.BytesRegexp(`((dtrace:::)?BEGIN|provider |#pragma (D (option|attributes)|ident)\s)`).Match(i) {
// 		return "DTrace", true
// 	} else if substring.BytesRegexp(`(\/.*:( .* \\)$| : \\$|^ : |: \\$)`).Match(i) {
// 		return "Makefile", true
// 	}

// 	return OtherLanguage, true
// },
// ".ecl": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^[^#]+:-`).Match(i) {
// 		return "ECLiPSe", true
// 	} else if substring.BytesHas(`:=`).Match(i) {
// 		return "ECL", true
// 	}

// 	return OtherLanguage, true
// },
// ".es": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\s*(?:%%|main\s*\(.*?\)\s*->)`).Match(i) {
// 		return "Erlang", true
// 	} else if substring.BytesRegexp(`(?:\/\/|("|')use strict\\1|export\s+default\s|\/\*.*?\*\/)`).Match(i) {
// 		return "JavaScript", true
// 	}

// 	return OtherLanguage, true
// },
// ".f": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\n: `).Match(i) {
// 		return "Forth", true
// 	} else if substring.BytesRegexp(`(?i)^([c*][^abd-z]|      (subroutine|program|end)\s|\s*!)`).Match(i) {
// 		return "FORTRAN", true
// 	}

// 	return OtherLanguage, false
// },
// ".fr": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^(: |also |new-device|previous )`).Match(i) {
// 		return "Forth", true
// 	} else if substring.BytesRegexp(`\s*(import|module|package|data|type)`).Match(i) {
// 		return "Frege", true
// 	}

// 	return "Text", false
// },
// ".j": func(i io.RuneReader) []string {
// 	if objectiveCMatcher.Match(i) {
// 		return "Objective-J", true
// 	}

// 	return "Jasmin", false
// },
// ".inc": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^<\?(?:php)?`).Match(i) {
// 		return "PHP", true
// 	}

// 	return OtherLanguage, true
// },
// ".m": func(i io.RuneReader) []string {
// 	if objectiveCMatcher.Match(i) {
// 		return "Objective-C", true
// 	} else if substring.BytesHas(`:- module`).Match(i) {
// 		return "Mercury", true
// 	} else if substring.BytesRegexp(`\n: `).Match(i) {
// 		return "MUF", true
// 	} else if substring.BytesRegexp(`\n\s*;`).Match(i) {
// 		return "M", true
// 	} else if substring.BytesRegexp(`\n\s*\(\*`).Match(i) {
// 		return "Mathematica", true
// 	} else if substring.BytesRegexp(`\n\s*%`).Match(i) {
// 		return "Matlab", true
// 	} else if substring.BytesRegexp(`\w+\s*:\s*module\s*{`).Match(i) {
// 		return "Limbo", true
// 	}

// 	return OtherLanguage, false
// },
// ".ms": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`[.'][a-z][a-z](\s|$)`).Match(i) {
// 		return "Groff", true
// 	}

// 	return "MAXScript", true
// },
// ".md": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\n[-a-z0-9=#!\*\[|]`).Match(i) {
// 		return "Markdown", true
// 	} else if substring.BytesRegexp(`\n(;;|\(define_)`).Match(i) {
// 		return "GCC Machine Description", true
// 	}

// 	return OtherLanguage, false
// },
// ".moo": func(i io.RuneReader) []string {
// 	if substring.BytesHas(`:- module`).Match(i) {
// 		return "Mercury", true
// 	}

// 	return "Moocode", false
// },
// ".e": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`feature\s--`).Match(i) {
// 		return "Eiffel", true
// 	}

// 	return "E", false
// },
// ".fs": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\n(: |new-device)`).Match(i) {
// 		return "Forth", true
// 	} else if substring.BytesRegexp(`\s*(#light|import|let|module|namespace|open|type)`).Match(i) {
// 		return "F#", true
// 	} else if substring.BytesRegexp(`(#version|precision|uniform|varying|vec[234])`).Match(i) {
// 		return "GLSL", true
// 	} else if substring.BytesRegexp(`#include|#pragma\s+(rs|version)|__attribute__`).Match(i) {
// 		return "Filterscript", true
// 	}

// 	return OtherLanguage, false
// },
// ".gs": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`uses java\.`).Match(i) {
// 		return "Gosu", true
// 	}

// 	return "JavaScript", false
// },
// ".h": func(i io.RuneReader) []string {
// 	if objectiveCMatcher.Match(i) {
// 		return "Objective-C", true
// 	} else if cPlusPlusMatcher.Match(i) {
// 		return "C++", true
// 	}

// 	return "C", true
// },
// ".hh": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^<\?(?:hh)?`).Match(i) {
// 		return "Hack", true
// 	} else if cPlusPlusMatcher.Match(i) {
// 		return "C++", true
// 	}

// 	return OtherLanguage, false
// },
// ".l": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\(def(un|macro)\s`).Match(i) {
// 		return "Common Lisp", true
// 	} else if substring.BytesRegexp(`(%[%{}]xs|<.*>)`).Match(i) {
// 		return "Lex", true
// 	} else if substring.BytesRegexp(`\.[a-z][a-z](\s|$)`).Match(i) {
// 		return "Groff", true
// 	} else if substring.BytesRegexp(`(de|class|rel|code|data|must)`).Match(i) {
// 		return "PicoLisp", true
// 	}

// 	return OtherLanguage, false
// },
// ".ls": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\s*package\s*[\w\.\/\*\s]*\s*{`).Match(i) {
// 		return "LoomScript", true
// 	}

// 	return "LiveScript", false
// },
// ".n": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`^[.']`).Match(i) {
// 		return "Groff", true
// 	} else if substring.BytesRegexp(`(module|namespace|using)`).Match(i) {
// 		return "Nemerle", true
// 	}

// 	return OtherLanguage, false
// },
// ".ncl": func(i io.RuneReader) []string {
// 	if substring.BytesHas("THE_TITLE").Match(i) {
// 		return "Text", true
// 	}

// 	return "NCL", true
// },
// ".mod": func(i io.RuneReader) []string {
// 	if substring.BytesHas("<!ENTITY ").Match(i) {
// 		return "XML", true
// 	} else if substring.BytesRegexp(`MODULE\s\w+\s*;`).Match(i) || substring.BytesRegexp(`(?i)\s*END \w+;$`).Match(i) {
// 		return "Modula-2", true
// 	}

// 	return "Linux Kernel Module", true
// },
// ".lisp": func(i io.RuneReader) []string {
// 	if commonLispMatcher.Match(i) {
// 		return "Common Lisp", true
// 	} else if substring.BytesRegexp(`\s*\(define `).Match(i) {
// 		return "NewLisp", true
// 	}

// 	return OtherLanguage, false
// },
// ".pm": func(i io.RuneReader) []string {
// 	if perlMatcher.Match(i) {
// 		return "Perl", true
// 	} else if perl6Matcher.Match(i) {
// 		return "Perl6", true
// 	}

// 	return "Perl", false
// },
// ".pp": func(i io.RuneReader) []string {
// 	if pascalMatcher.Match(i) {
// 		return "Pascal", true
// 	}
// 	return "Puppet", false
// },
// ".t": func(i io.RuneReader) []string {
// 	if perlMatcher.Match(i) {
// 		return "Perl", true
// 	} else if perl6Matcher.Match(i) {
// 		return "Perl6", true
// 	} else if substring.BytesRegexp(`^\s*%|^\s*var\s+\w+\s*:\s*\w+`).Match(i) {
// 		return "Turing", true
// 	} else if substring.BytesRegexp(`^\s*use\s+v6\s*;`).Match(i) {
// 		return "Perl6", true
// 	} else if substring.BytesRegexp(`terra\s`).Match(i) {
// 		return "Terra", true
// 	}

// 	return "Perl", false
// },
// ".ts": func(i io.RuneReader) []string {
// 	if substring.BytesHas("</TS>").Match(i) {
// 		return "XML", true
// 	}

// 	return "TypeScript", true
// },
// ".tsx": func(i io.RuneReader) []string {
// 	if substring.BytesHas("</tileset>").Match(i) {
// 		return "XML", true
// 	}

// 	return "TypeScript", true
// },
// ".tst": func(i io.RuneReader) []string {
// 	if substring.BytesHas("gap> ").Match(i) {
// 		return "GAP", true
// 	}

// 	return "Scilab", true
// },
// ".r": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`(?i)\bRebol\b`).Match(i) {
// 		return "Rebol", true
// 	} else if substring.BytesHas("<-").Match(i) {
// 		return "R", true
// 	}

// 	return OtherLanguage, false
// },
// ".rs": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\n(use |fn |mod |pub |macro_rules|impl|#!?\[)`).Match(i) {
// 		return "Rust", true
// 	} else if substring.BytesRegexp(`#include|#pragma\s+(rs|version)|__attribute__`).Match(i) {
// 		return "RenderScript", true
// 	}

// 	return OtherLanguage, false
// },
// ".rpy": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`(import|from|class|def)\s`).Match(i) {
// 		return "Python", true
// 	}

// 	return "Ren'Py", false
// },
// ".v": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`\nendmodule`).Match(i) {
// 		return "Verilog", true
// 	} else if substring.BytesRegexp(`(Require|Import)`).Match(i) {
// 		return "Coq", true
// 	}

// 	return OtherLanguage, false
// },
// ".pl": func(i io.RuneReader) []string {
// 	if prologMatcher.Match(i) {
// 		return "Prolog", true
// 	} else if perl6Matcher.Match(i) {
// 		return "Perl6", true
// 	}

// 	return "Perl", false
// },
// ".pro": func(i io.RuneReader) []string {
// 	if prologMatcher.Match(i) {
// 		return "Prolog", true
// 	}

// 	return OtherLanguage, false
// },
// ".pod": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp(`=\w+\n`).Match(i) {
// 		return "Pod", true
// 	}

// 	return "Perl", false
// },
// ".toc": func(i io.RuneReader) []string {
// 	if substring.BytesRegexp("## |@no-lib-strip@").Match(i) {
// 		return "World of Warcraft Addon Data", true
// 	} else if substring.BytesRegexp("(contentsline|defcounter|beamer|boolfalse)").Match(i) {
// 		return "TeX", true
// 	}

// 	return OtherLanguage, false
// },
// ".sls": func(i io.RuneReader) []string {
// 	if schemeMatcher.Match(i) {
// 		return "Scheme", true
// 	}
// 	return "SaltStack", false
// },
// ".sql": func(i io.RuneReader) []string {
// 	if pgSQLMatcher.Match(i) {
// 		return "PLpgSQL", true
// 	} else if db2SQLMatcher.Match(i) {
// 		return "SQLPL", true
// 	} else if oracleSQLMatcher.Match(i) {
// 		return "PLSQL", true
// 	}

// 	return "SQL", false
// },
}

func init() {
	matchers[".for"] = matchers[".f"]
	matchers[".lsp"] = matchers[".lisp"]
}

var (
	blitzBasicMatcher = substring.BytesOr(
		substring.BytesHas(`End Function`),
		substring.BytesRegexp(`\\s*;`),
	)
	cPlusPlusMatcher = substring.BytesOr(
		substring.BytesRegexp(`\s*template\s*<`),
		substring.BytesRegexp(`\s*#\s*include <(cstdint|string|vector|map|list|array|bitset|queue|stack|forward_list|unordered_map|unordered_set|(i|o|io)stream)>`),
		substring.BytesRegexp(`\n[ \t]*try`),
		substring.BytesRegexp(`\n[ \t]*(class|(using[ \t]+)?namespace)\s+\w+`),
		substring.BytesRegexp(`\n[ \t]*(private|public|protected):\n`),
		substring.BytesRegexp(`std::\w+`),
		substring.BytesRegexp(`[ \t]*catch\s*`),
	)
	commonLispMatcher = substring.BytesRegexp("(?i)(defpackage|defun|in-package)")
	coolMatcher       = substring.BytesRegexp("(?i)class")
	openCLMatcher     = substring.BytesOr(
		substring.BytesHas("\n}"),
		substring.BytesHas("}\n"),
		substring.BytesHas(`/*`),
		substring.BytesHas(`//`),
	)
	apexMatcher = substring.BytesOr(
		substring.BytesHas("{\n"),
		substring.BytesHas("}\n"),
	)
	texMatcher = substring.BytesOr(
		substring.BytesHas(`%`),
		substring.BytesHas(`\`),
	)
	openEdgeABLMatcher = substring.BytesRegexp(`(?i)(class|define|interface|method|using)\b`)
	visualBasicMatcher = substring.BytesOr(
		substring.BytesHas("'*"),
		substring.BytesRegexp(`(?i)(attribute|option|sub|private|protected|public|friend)\b`),
	)
	mathematicaMatcher = substring.BytesHas(`^\s*\(\*`)
	matlabMatcher      = substring.BytesRegexp(`\b(function\s*[\[a-zA-Z]+|pcolor|classdef|figure|end|elseif)\b`)
	objectiveCMatcher  = substring.BytesRegexp(
		`@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">]`)
	pascalMatcher = substring.BytesRegexp(`(?ims)^\s*(PROGRAM|UNIT|USES|FUNCTION)[\s\n]+.*?;`)
	prologMatcher = substring.BytesRegexp(`^[^#]+:-`)
	perlMatcher   = substring.BytesRegexp(`use strict|use\s+v?5\.`)
	perl6Matcher  = substring.BytesRegexp(`(use v6|(my )?class|module)`)
	pgSQLMatcher  = substring.BytesOr(
		substring.BytesRegexp(`(?i)\\i\b|AS \$\$|LANGUAGE '?plpgsql'?`),
		substring.BytesRegexp(`(?i)SECURITY (DEFINER|INVOKER)`),
		substring.BytesRegexp(`BEGIN( WORK| TRANSACTION)?;`),
	)
	db2SQLMatcher = substring.BytesOr(
		substring.BytesRegexp(`(?i)(alter module)|(language sql)|(begin( NOT)+ atomic)`),
		substring.BytesRegexp(`(?i)signal SQLSTATE '[0-9]+'`),
	)
	oracleSQLMatcher = substring.BytesOr(
		substring.BytesRegexp(`(?i)\$\$PLSQL_|XMLTYPE|sysdate|systimestamp|\.nextval|connect by|AUTHID (DEFINER|CURRENT_USER)`),
		substring.BytesRegexp(`(?i)constructor\W+function`),
	)
	schemeMatcher = substring.BytesRegexp(`(?m)\A(^\s*;;.*$)*\s*\(`)
)
