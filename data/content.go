// Code generated by gopkg.in/src-d/enry.v1/internal/code-generator DO NOT EDIT.
// Extracted from github/linguist commit: e761f9b013e5b61161481fcb898b59721ee40e3d

package data

import (
	"regexp"

	"gopkg.in/src-d/enry.v1/data/rule"
)

var ContentHeuristics = map[string]*Heuristics{
	".as": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("ActionScript"),
			regexp.MustCompile(`(?m)^\s*(package\s+[a-z0-9_\.]+|import\s+[a-zA-Z0-9_\.]+;|class\s+[A-Za-z0-9_]+\s+extends\s+[A-Za-z0-9_]+)`),
		),
		rule.Always(
			rule.MatchingLanguages("AngelScript"),
		),
	},
	".asc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Public Key"),
			regexp.MustCompile(`(?m)^(----[- ]BEGIN|ssh-(rsa|dss)) `),
		),
		rule.Or(
			rule.MatchingLanguages("AsciiDoc"),
			regexp.MustCompile(`(?m)^[=-]+(\s|\n)|{{[A-Za-z]`),
		),
		rule.Or(
			rule.MatchingLanguages("AGS Script"),
			regexp.MustCompile(`(?m)^(\/\/.+|((import|export)\s+)?(function|int|float|char)\s+((room|repeatedly|on|game)_)?([A-Za-z]+[A-Za-z_0-9]+)\s*[;\(])`),
		),
	},
	".bb": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("BlitzBasic"),
			regexp.MustCompile(`(?m)(<^\s*; |End Function)`),
		),
		rule.Or(
			rule.MatchingLanguages("BitBake"),
			regexp.MustCompile(`(?m)^\s*(# |include|require)\b`),
		),
	},
	".builds": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^(\s*)(?i:<Project|<Import|<Property|<?xml|xmlns)`),
		),
		rule.Always(
			rule.MatchingLanguages("Text"),
		),
	},
	".ch": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("xBase"),
			regexp.MustCompile(`(?m)^\s*#\s*(?i:if|ifdef|ifndef|define|command|xcommand|translate|xtranslate|include|pragma|undef)\b`),
		),
	},
	".cl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("Cool"),
			regexp.MustCompile(`(?m)^class`),
		),
		rule.Or(
			rule.MatchingLanguages("OpenCL"),
			regexp.MustCompile(`(?m)\/\* |\/\/ |^\}`),
		),
	},
	".cls": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("TeX"),
			regexp.MustCompile(`(?m)\\\w+{`),
		),
	},
	".cs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Smalltalk"),
			regexp.MustCompile(`(?m)![\w\s]+methodsFor: `),
		),
		rule.Or(
			rule.MatchingLanguages("C#"),
			regexp.MustCompile(`(?m)^(\s*namespace\s*[\w\.]+\s*{|\s*\/\/)`),
		),
	},
	".d": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("D"),
			regexp.MustCompile(`(?m)^module\s+[\w.]*\s*;|import\s+[\w\s,.:]*;|\w+\s+\w+\s*\(.*\)(?:\(.*\))?\s*{[^}]*}|unittest\s*(?:\(.*\))?\s*{[^}]*}`),
		),
		rule.Or(
			rule.MatchingLanguages("DTrace"),
			regexp.MustCompile(`(?m)^(\w+:\w*:\w*:\w*|BEGIN|END|provider\s+|(tick|profile)-\w+\s+{[^}]*}|#pragma\s+D\s+(option|attributes|depends_on)\s|#pragma\s+ident\s)`),
		),
		rule.Or(
			rule.MatchingLanguages("Makefile"),
			regexp.MustCompile(`(?m)([\/\\].*:\s+.*\s\\$|: \\$|^ : |^[\w\s\/\\.]+\w+\.\w+\s*:\s+[\w\s\/\\.]+\w+\.\w+)`),
		),
	},
	".ecl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("ECLiPSe"),
			regexp.MustCompile(`(?m)^[^#]+:-`),
		),
		rule.Or(
			rule.MatchingLanguages("ECL"),
			regexp.MustCompile(`(?m):=`),
		),
	},
	".es": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Erlang"),
			regexp.MustCompile(`(?m)^\s*(?:%%|main\s*\(.*?\)\s*->)`),
		),
	},
	".f": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("Filebench WML"),
			regexp.MustCompile(`(?m)flowop`),
		),
		rule.Or(
			rule.MatchingLanguages("Fortran"),
			regexp.MustCompile(`(?m)^(?i:[c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`),
		),
	},
	".for": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("Fortran"),
			regexp.MustCompile(`(?m)^(?i:[c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)`),
		),
	},
	".fr": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^(: |also |new-device|previous )`),
		),
		rule.Or(
			rule.MatchingLanguages("Frege"),
			regexp.MustCompile(`(?m)^\s*(import|module|package|data|type) `),
		),
		rule.Always(
			rule.MatchingLanguages("Text"),
		),
	},
	".fs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Forth"),
			regexp.MustCompile(`(?m)^(: |new-device)`),
		),
		rule.Or(
			rule.MatchingLanguages("F#"),
			regexp.MustCompile(`(?m)^\s*(#light|import|let|module|namespace|open|type)`),
		),
		rule.Or(
			rule.MatchingLanguages("GLSL"),
			regexp.MustCompile(`(?m)^\s*(#version|precision|uniform|varying|vec[234])`),
		),
		rule.Or(
			rule.MatchingLanguages("Filterscript"),
			regexp.MustCompile(`(?m)#include|#pragma\s+(rs|version)|__attribute__`),
		),
	},
	".gml": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)(?i:^\s*(\<\?xml|xmlns))`),
		),
		rule.Or(
			rule.MatchingLanguages("Graph Modeling Language"),
			regexp.MustCompile(`(?m)(?i:^\s*(graph|node)\s+\[$)`),
		),
		rule.Always(
			rule.MatchingLanguages("Game Maker Language"),
		),
	},
	".gs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Gosu"),
			regexp.MustCompile(`(?m)^uses java\.`),
		),
	},
	".h": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Objective-C"),
			regexp.MustCompile(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`),
		),
		rule.Or(
			rule.MatchingLanguages("C++"),
			regexp.MustCompile(`(?m)^\s*#\s*include <(cstdint|string|vector|map|list|array|bitset|queue|stack|forward_list|unordered_map|unordered_set|(i|o|io)stream)>|^\s*template\s*<|^[ \t]*try|^[ \t]*catch\s*\(|^[ \t]*(class|(using[ \t]+)?namespace)\s+\w+|^[ \t]*(private|public|protected):$|std::\w+`),
		),
	},
	".hh": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Hack"),
			regexp.MustCompile(`(?m)<\?hh`),
		),
	},
	".ice": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Slice"),
			regexp.MustCompile(`(?m)^\s*(#\s*(include|if[n]def|pragma)|module\s+[A-Za-z][_A-Za-z0-9]*)`),
		),
	},
	".inc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("PHP"),
			regexp.MustCompile(`(?m)^<\?(?:php)?`),
		),
		rule.Or(
			rule.MatchingLanguages("POV-Ray SDL"),
			regexp.MustCompile(`(?m)^\s*#(declare|local|macro|while)\s`),
		),
	},
	".l": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)\(def(un|macro)\s`),
		),
		rule.Or(
			rule.MatchingLanguages("Lex"),
			regexp.MustCompile(`(?m)^(%[%{}]xs|<.*>)`),
		),
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^\.[A-Za-z]{2}(\s|$)`),
		),
		rule.Or(
			rule.MatchingLanguages("PicoLisp"),
			regexp.MustCompile(`(?m)^\((de|class|rel|code|data|must)\s`),
		),
	},
	".lisp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("NewLisp"),
			regexp.MustCompile(`(?m)^\s*\(define `),
		),
	},
	".ls": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("LoomScript"),
			regexp.MustCompile(`(?m)^\s*package\s*[\w\.\/\*\s]*\s*{`),
		),
		rule.Always(
			rule.MatchingLanguages("LiveScript"),
		),
	},
	".lsp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Common Lisp"),
			regexp.MustCompile(`(?m)^\s*\((?i:defun|in-package|defpackage) `),
		),
		rule.Or(
			rule.MatchingLanguages("NewLisp"),
			regexp.MustCompile(`(?m)^\s*\(define `),
		),
	},
	".m": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Objective-C"),
			regexp.MustCompile(`(?m)^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])`),
		),
		rule.Or(
			rule.MatchingLanguages("Mercury"),
			regexp.MustCompile(`(?m):- module`),
		),
		rule.Or(
			rule.MatchingLanguages("MUF"),
			regexp.MustCompile(`(?m)^: `),
		),
		rule.Or(
			rule.MatchingLanguages("M"),
			regexp.MustCompile(`(?m)^\s*;`),
		),
		rule.Or(
			rule.MatchingLanguages("Mathematica"),
			regexp.MustCompile(`(?m)\*\)$`),
		),
		rule.Or(
			rule.MatchingLanguages("MATLAB"),
			regexp.MustCompile(`(?m)^\s*%`),
		),
		rule.Or(
			rule.MatchingLanguages("Limbo"),
			regexp.MustCompile(`(?m)^\w+\s*:\s*module\s*{`),
		),
	},
	".md": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Markdown"),
			regexp.MustCompile(`(?m)(^[-A-Za-z0-9=#!\*\[|>])|<\/|\A\z`),
		),
		rule.Or(
			rule.MatchingLanguages("GCC Machine Description"),
			regexp.MustCompile(`(?m)^(;;|\(define_)`),
		),
		rule.Always(
			rule.MatchingLanguages("Markdown"),
		),
	},
	".ml": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("OCaml"),
			regexp.MustCompile(`(?m)(^\s*module)|let rec |match\s+(\S+\s)+with`),
		),
		rule.Or(
			rule.MatchingLanguages("Standard ML"),
			regexp.MustCompile(`(?m)=> |case\s+(\S+\s)+of`),
		),
	},
	".mod": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)<!ENTITY `),
		),
		rule.Or(
			rule.MatchingLanguages("Modula-2"),
			regexp.MustCompile(`(?m)^\s*(?i:MODULE|END) [\w\.]+;`),
		),
		rule.Always(
			rule.MatchingLanguages("Linux Kernel Module", "AMPL"),
		),
	},
	".ms": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^[.'][A-Za-z]{2}(\s|$)`),
		),
		rule.And(
			rule.MatchingLanguages("Unix Assembly"),
			rule.Not(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)/\*`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^\s*\.(?:include\s|globa?l\s|[A-Za-z][_A-Za-z0-9]*:)`),
			),
		),
		rule.Always(
			rule.MatchingLanguages("MAXScript"),
		),
	},
	".n": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^[.']`),
		),
		rule.Or(
			rule.MatchingLanguages("Nemerle"),
			regexp.MustCompile(`(?m)^(module|namespace|using)\s`),
		),
	},
	".ncl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^\s*<\?xml\s+version`),
		),
		rule.Or(
			rule.MatchingLanguages("Text"),
			regexp.MustCompile(`(?m)THE_TITLE`),
		),
	},
	".nl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("NL"),
			regexp.MustCompile(`(?m)^(b|g)[0-9]+ `),
		),
		rule.Always(
			rule.MatchingLanguages("NewLisp"),
		),
	},
	".php": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Hack"),
			regexp.MustCompile(`(?m)<\?hh`),
		),
		rule.Or(
			rule.MatchingLanguages("PHP"),
			regexp.MustCompile(`(?m)<\?[^h]`),
		),
	},
	".pl": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Prolog"),
			regexp.MustCompile(`(?m)^[^#]*:-`),
		),
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Perl 6"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
	},
	".pm": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Perl 6"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("XPM"),
			regexp.MustCompile(`(?m)^\s*\/\* XPM \*\/`),
		),
	},
	".pod": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Pod 6"),
			regexp.MustCompile(`(?m)^[\s&&[^\n]]*=(comment|begin pod|begin para|item\d+)`),
		),
		rule.Always(
			rule.MatchingLanguages("Pod"),
		),
	},
	".pp": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Pascal"),
			regexp.MustCompile(`(?m)^\s*end[.;]`),
		),
		rule.Or(
			rule.MatchingLanguages("Puppet"),
			regexp.MustCompile(`(?m)^\s+\w+\s+=>\s`),
		),
	},
	".pro": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Prolog"),
			regexp.MustCompile(`(?m)^[^\[#]+:-`),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)last_client=`),
		),
		rule.And(
			rule.MatchingLanguages("QMake"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)HEADERS`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)SOURCES`),
			),
		),
		rule.Or(
			rule.MatchingLanguages("IDL"),
			regexp.MustCompile(`(?m)^\s*function[ \w,]+$`),
		),
	},
	".properties": &Heuristics{
		rule.And(
			rule.MatchingLanguages("INI"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[;\[]`),
			),
		),
		rule.And(
			rule.MatchingLanguages("Java Properties"),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
			),
			rule.Or(
				rule.MatchingLanguages(""),
				regexp.MustCompile(`(?m)^[#!]`),
			),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)^[^#!;][^=]*=`),
		),
		rule.Or(
			rule.MatchingLanguages("Java properties"),
			regexp.MustCompile(`(?m)^[^#!][^:]*:`),
		),
	},
	".props": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)^(\s*)(?i:<Project|<Import|<Property|<?xml|xmlns)`),
		),
		rule.Or(
			rule.MatchingLanguages("INI"),
			regexp.MustCompile(`(?m)(?i:\w+\s*=\s*)`),
		),
	},
	".q": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("q"),
			regexp.MustCompile(`(?m)((?i:[A-Z.][\w.]*:{)|(^|\n)\\(cd?|d|l|p|ts?) )`),
		),
		rule.Or(
			rule.MatchingLanguages("HiveQL"),
			regexp.MustCompile(`(?m)(?i:SELECT\s+[\w*,]+\s+FROM|(CREATE|ALTER|DROP)\s(DATABASE|SCHEMA|TABLE))`),
		),
	},
	".r": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Rebol"),
			regexp.MustCompile(`(?m)(?i:\bRebol\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("R"),
			regexp.MustCompile(`(?m)<-|^\s*#`),
		),
	},
	".rno": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("RUNOFF"),
			regexp.MustCompile(`(?m)(?i:^\.!|^\.end lit(?:eral)?\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("Roff"),
			regexp.MustCompile(`(?m)^\.\\" `),
		),
	},
	".rpy": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Python"),
			regexp.MustCompile(`(?m)(?m:^(import|from|class|def)\s)`),
		),
		rule.Always(
			rule.MatchingLanguages("Ren'Py"),
		),
	},
	".rs": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Rust"),
			regexp.MustCompile(`(?m)^(use |fn |mod |pub |macro_rules|impl|#!?\[)`),
		),
		rule.Or(
			rule.MatchingLanguages("RenderScript"),
			regexp.MustCompile(`(?m)#include|#pragma\s+(rs|version)|__attribute__`),
		),
	},
	".sc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("SuperCollider"),
			regexp.MustCompile(`(?m)(?i:\^(this|super)\.|^\s*~\w+\s*=\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Scala"),
			regexp.MustCompile(`(?m)(^\s*import (scala|java)\.|^\s*class\b)`),
		),
	},
	".sql": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("PLpgSQL"),
			regexp.MustCompile(`(?m)(?i:^\\i\b|AS \$\$|LANGUAGE '?plpgsql'?|SECURITY (DEFINER|INVOKER)|BEGIN( WORK| TRANSACTION)?;)`),
		),
		rule.Or(
			rule.MatchingLanguages("SQLPL"),
			regexp.MustCompile(`(?m)(?i:(alter module)|(language sql)|(begin( NOT)+ atomic)|signal SQLSTATE '[0-9]+')`),
		),
		rule.Or(
			rule.MatchingLanguages("PLSQL"),
			regexp.MustCompile(`(?m)(?i:\$\$PLSQL_|XMLTYPE|sysdate|systimestamp|\.nextval|connect by|AUTHID (DEFINER|CURRENT_USER)|constructor\W+function)`),
		),
		rule.Not(
			rule.MatchingLanguages("SQL"),
			regexp.MustCompile(`(?m)(?i:begin|boolean|package|exception)`),
		),
	},
	".srt": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("SubRip Text"),
			regexp.MustCompile(`(?m)^(\d{2}:\d{2}:\d{2},\d{3})\s*(-->)\s*(\d{2}:\d{2}:\d{2},\d{3})$`),
		),
	},
	".t": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("Perl"),
			regexp.MustCompile(`(?m)\buse\s+(?:strict\b|v?5\.)`),
		),
		rule.Or(
			rule.MatchingLanguages("Perl 6"),
			regexp.MustCompile(`(?m)^\s*(?:use\s+v6\b|\bmodule\b|\b(?:my\s+)?class\b)`),
		),
		rule.Or(
			rule.MatchingLanguages("Turing"),
			regexp.MustCompile(`(?m)^\s*%[ \t]+|^\s*var\s+\w+(\s*:\s*\w+)?\s*:=\s*\w+`),
		),
	},
	".toc": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("World of Warcraft Addon Data"),
			regexp.MustCompile(`(?m)^## |@no-lib-strip@`),
		),
		rule.Or(
			rule.MatchingLanguages("TeX"),
			regexp.MustCompile(`(?m)^\\(contentsline|defcounter|beamer|boolfalse)`),
		),
	},
	".ts": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)<TS\b`),
		),
		rule.Always(
			rule.MatchingLanguages("TypeScript"),
		),
	},
	".tst": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("GAP"),
			regexp.MustCompile(`(?m)gap> `),
		),
		rule.Always(
			rule.MatchingLanguages("Scilab"),
		),
	},
	".tsx": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("TypeScript"),
			regexp.MustCompile(`(?m)^\s*(import.+(from\s+|require\()['"]react|\/\/\/\s*<reference\s)`),
		),
		rule.Or(
			rule.MatchingLanguages("XML"),
			regexp.MustCompile(`(?m)(?i:^\s*<\?xml\s+version)`),
		),
	},
	".w": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("OpenEdge ABL"),
			regexp.MustCompile(`(?m)&ANALYZE-SUSPEND _UIB-CODE-BLOCK _CUSTOM _DEFINITIONS`),
		),
		rule.Or(
			rule.MatchingLanguages("CWeb"),
			regexp.MustCompile(`(?m)^@(<|\w+\.)`),
		),
	},
	".x": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("RPC"),
			regexp.MustCompile(`(?m)\b(program|version)\s+\w+\s*{|\bunion\s+\w+\s+switch\s*\(`),
		),
		rule.Or(
			rule.MatchingLanguages("Logos"),
			regexp.MustCompile(`(?m)^%(end|ctor|hook|group)\b`),
		),
		rule.Or(
			rule.MatchingLanguages("Linked Script"),
			regexp.MustCompile(`(?m)OUTPUT_ARCH\(|OUTPUT_FORMAT\(|SECTIONS`),
		),
	},
	".yy": &Heuristics{
		rule.Or(
			rule.MatchingLanguages("JSON"),
			regexp.MustCompile(`(?m)\"modelName\"\:\s*\"GM`),
		),
		rule.Always(
			rule.MatchingLanguages("Yacc"),
		),
	},
}
