package slinguist

// CODE GENERATED AUTOMATICALLY WITH github.com/src-d/simple-linguist/cli/slinguist-generate
// THIS FILE SHOULD NOT BE EDITED BY HAND
// Extracted from github/linguist commit: 1a9ff6ed2e16e9af9e9d3bc6817ac46eb3421e00

import (
	"path/filepath"
	"regexp"
	"strings"
)

func GetLanguageByContent(filename string, content []byte) (lang string, safe bool) {
	ext := strings.ToLower(filepath.Ext(filename))
	if fnMatcher, ok := matchers[ext]; ok {
		lang, safe = fnMatcher(content)
		return
	}

	return GetLanguageByExtension(filename)
}

type languageMatcher func([]byte) (string, bool)

var matchers = map[string]languageMatcher{
	".asc": func(i []byte) (string, bool) {
		if asc_PublicKey_Matcher_0.Match(i) {
			return "Public Key", true
		} else if asc_AsciiDoc_Matcher_0.Match(i) {
			return "AsciiDoc", true
		} else if asc_AGSScript_Matcher_0.Match(i) {
			return "AGS Script", true
		}

		return OtherLanguage, false
	},
	".bb": func(i []byte) (string, bool) {
		if bb_BlitzBasic_Matcher_0.Match(i) || bb_BlitzBasic_Matcher_1.Match(i) {
			return "BlitzBasic", true
		} else if bb_BitBake_Matcher_0.Match(i) {
			return "BitBake", true
		}

		return OtherLanguage, false
	},
	".builds": func(i []byte) (string, bool) {
		if builds_XML_Matcher_0.Match(i) {
			return "XML", true
		}

		return "Text", false
	},
	".ch": func(i []byte) (string, bool) {
		if ch_xBase_Matcher_0.Match(i) {
			return "xBase", true
		}

		return OtherLanguage, false
	},
	".cl": func(i []byte) (string, bool) {
		if cl_CommonLisp_Matcher_0.Match(i) {
			return "Common Lisp", true
		} else if cl_Cool_Matcher_0.Match(i) {
			return "Cool", true
		} else if cl_OpenCL_Matcher_0.Match(i) {
			return "OpenCL", true
		}

		return OtherLanguage, false
	},
	".cls": func(i []byte) (string, bool) {
		if cls_TeX_Matcher_0.Match(i) {
			return "TeX", true
		}

		return OtherLanguage, false
	},
	".cs": func(i []byte) (string, bool) {
		if cs_Smalltalk_Matcher_0.Match(i) {
			return "Smalltalk", true
		} else if cs_CSharp_Matcher_0.Match(i) || cs_CSharp_Matcher_1.Match(i) {
			return "C#", true
		}

		return OtherLanguage, false
	},
	".d": func(i []byte) (string, bool) {
		if d_D_Matcher_0.Match(i) {
			return "D", true
		} else if d_DTrace_Matcher_0.Match(i) {
			return "DTrace", true
		} else if d_Makefile_Matcher_0.Match(i) {
			return "Makefile", true
		}

		return OtherLanguage, false
	},
	".ecl": func(i []byte) (string, bool) {
		if ecl_ECLiPSe_Matcher_0.Match(i) {
			return "ECLiPSe", true
		} else if ecl_ECL_Matcher_0.Match(i) {
			return "ECL", true
		}

		return OtherLanguage, false
	},
	".es": func(i []byte) (string, bool) {
		if es_Erlang_Matcher_0.Match(i) {
			return "Erlang", true
		}

		return OtherLanguage, false
	},
	".f": func(i []byte) (string, bool) {
		if f_Forth_Matcher_0.Match(i) {
			return "Forth", true
		} else if f_FilebenchWML_Matcher_0.Match(i) {
			return "Filebench WML", true
		} else if f_FORTRAN_Matcher_0.Match(i) {
			return "FORTRAN", true
		}

		return OtherLanguage, false
	},
	".for": func(i []byte) (string, bool) {
		if for_Forth_Matcher_0.Match(i) {
			return "Forth", true
		} else if for_FORTRAN_Matcher_0.Match(i) {
			return "FORTRAN", true
		}

		return OtherLanguage, false
	},
	".fr": func(i []byte) (string, bool) {
		if fr_Forth_Matcher_0.Match(i) {
			return "Forth", true
		} else if fr_Frege_Matcher_0.Match(i) {
			return "Frege", true
		}

		return "Text", false
	},
	".fs": func(i []byte) (string, bool) {
		if fs_Forth_Matcher_0.Match(i) {
			return "Forth", true
		} else if fs_FSharp_Matcher_0.Match(i) {
			return "F#", true
		} else if fs_GLSL_Matcher_0.Match(i) {
			return "GLSL", true
		} else if fs_Filterscript_Matcher_0.Match(i) {
			return "Filterscript", true
		}

		return OtherLanguage, false
	},
	".gs": func(i []byte) (string, bool) {
		if gs_Gosu_Matcher_0.Match(i) {
			return "Gosu", true
		}

		return OtherLanguage, false
	},
	".h": func(i []byte) (string, bool) {
		if h_ObjectiveDashC_Matcher_0.Match(i) {
			return "Objective-C", true
		} else if h_CPlusPlus_Matcher_0.Match(i) || h_CPlusPlus_Matcher_1.Match(i) || h_CPlusPlus_Matcher_2.Match(i) || h_CPlusPlus_Matcher_3.Match(i) || h_CPlusPlus_Matcher_4.Match(i) || h_CPlusPlus_Matcher_5.Match(i) || h_CPlusPlus_Matcher_6.Match(i) {
			return "C++", true
		}

		return OtherLanguage, false
	},
	".inc": func(i []byte) (string, bool) {
		if inc_PHP_Matcher_0.Match(i) {
			return "PHP", true
		} else if inc_POVDashRaySDL_Matcher_0.Match(i) {
			return "POV-Ray SDL", true
		}

		return OtherLanguage, false
	},
	".l": func(i []byte) (string, bool) {
		if l_CommonLisp_Matcher_0.Match(i) {
			return "Common Lisp", true
		} else if l_Lex_Matcher_0.Match(i) {
			return "Lex", true
		} else if l_Groff_Matcher_0.Match(i) {
			return "Groff", true
		} else if l_PicoLisp_Matcher_0.Match(i) {
			return "PicoLisp", true
		}

		return OtherLanguage, false
	},
	".ls": func(i []byte) (string, bool) {
		if ls_LoomScript_Matcher_0.Match(i) {
			return "LoomScript", true
		}

		return "LiveScript", false
	},
	".lsp": func(i []byte) (string, bool) {
		if lsp_CommonLisp_Matcher_0.Match(i) {
			return "Common Lisp", true
		} else if lsp_NewLisp_Matcher_0.Match(i) {
			return "NewLisp", true
		}

		return OtherLanguage, false
	},
	".lisp": func(i []byte) (string, bool) {
		if lisp_CommonLisp_Matcher_0.Match(i) {
			return "Common Lisp", true
		} else if lisp_NewLisp_Matcher_0.Match(i) {
			return "NewLisp", true
		}

		return OtherLanguage, false
	},
	".m": func(i []byte) (string, bool) {
		if m_ObjectiveDashC_Matcher_0.Match(i) {
			return "Objective-C", true
		} else if m_Mercury_Matcher_0.Match(i) {
			return "Mercury", true
		} else if m_MUF_Matcher_0.Match(i) {
			return "MUF", true
		} else if m_M_Matcher_0.Match(i) {
			return "M", true
		} else if m_Mathematica_Matcher_0.Match(i) {
			return "Mathematica", true
		} else if m_Matlab_Matcher_0.Match(i) {
			return "Matlab", true
		} else if m_Limbo_Matcher_0.Match(i) {
			return "Limbo", true
		}

		return OtherLanguage, false
	},
	".md": func(i []byte) (string, bool) {
		if md_Markdown_Matcher_0.Match(i) || md_Markdown_Matcher_1.Match(i) {
			return "Markdown", true
		} else if md_GCCmachinedescription_Matcher_0.Match(i) {
			return "GCC machine description", true
		}

		return "Markdown", false
	},
	".ml": func(i []byte) (string, bool) {
		if ml_OCaml_Matcher_0.Match(i) {
			return "OCaml", true
		} else if ml_StandardML_Matcher_0.Match(i) {
			return "Standard ML", true
		}

		return OtherLanguage, false
	},
	".mod": func(i []byte) (string, bool) {
		if mod_XML_Matcher_0.Match(i) {
			return "XML", true
		} else if mod_ModulaDash2_Matcher_0.Match(i) || mod_ModulaDash2_Matcher_1.Match(i) {
			return "Modula-2", true
		}

		return "Linux Kernel Module", false
	},
	".ms": func(i []byte) (string, bool) {
		if ms_Groff_Matcher_0.Match(i) {
			return "Groff", true
		}

		return "MAXScript", false
	},
	".n": func(i []byte) (string, bool) {
		if n_Groff_Matcher_0.Match(i) {
			return "Groff", true
		} else if n_Nemerle_Matcher_0.Match(i) {
			return "Nemerle", true
		}

		return OtherLanguage, false
	},
	".ncl": func(i []byte) (string, bool) {
		if ncl_Text_Matcher_0.Match(i) {
			return "Text", true
		}

		return OtherLanguage, false
	},
	".nl": func(i []byte) (string, bool) {
		if nl_NL_Matcher_0.Match(i) {
			return "NL", true
		}

		return "NewLisp", false
	},
	".php": func(i []byte) (string, bool) {
		if php_Hack_Matcher_0.Match(i) {
			return "Hack", true
		} else if php_PHP_Matcher_0.Match(i) {
			return "PHP", true
		}

		return OtherLanguage, false
	},
	".pl": func(i []byte) (string, bool) {
		if pl_Prolog_Matcher_0.Match(i) {
			return "Prolog", true
		} else if pl_Perl_Matcher_0.Match(i) {
			return "Perl", true
		} else if pl_Perl6_Matcher_0.Match(i) {
			return "Perl6", true
		}

		return OtherLanguage, false
	},
	".pm": func(i []byte) (string, bool) {
		if pm_Perl_Matcher_0.Match(i) {
			return "Perl", true
		} else if pm_Perl6_Matcher_0.Match(i) {
			return "Perl6", true
		}

		return OtherLanguage, false
	},
	".t": func(i []byte) (string, bool) {
		if t_Perl_Matcher_0.Match(i) {
			return "Perl", true
		} else if t_Perl6_Matcher_0.Match(i) {
			return "Perl6", true
		}

		return OtherLanguage, false
	},
	".pod": func(i []byte) (string, bool) {
		if pod_Pod_Matcher_0.Match(i) {
			return "Pod", true
		}

		return "Perl", false
	},
	".pro": func(i []byte) (string, bool) {
		if pro_Prolog_Matcher_0.Match(i) {
			return "Prolog", true
		} else if pro_INI_Matcher_0.Match(i) {
			return "INI", true
		} else if pro_QMake_Matcher_0.Match(i) && pro_QMake_Matcher_1.Match(i) {
			return "QMake", true
		} else if pro_IDL_Matcher_0.Match(i) {
			return "IDL", true
		}

		return OtherLanguage, false
	},
	".props": func(i []byte) (string, bool) {
		if props_XML_Matcher_0.Match(i) {
			return "XML", true
		} else if props_INI_Matcher_0.Match(i) {
			return "INI", true
		}

		return OtherLanguage, false
	},
	".r": func(i []byte) (string, bool) {
		if r_Rebol_Matcher_0.Match(i) {
			return "Rebol", true
		} else if r_R_Matcher_0.Match(i) {
			return "R", true
		}

		return OtherLanguage, false
	},
	".rno": func(i []byte) (string, bool) {
		if rno_RUNOFF_Matcher_0.Match(i) {
			return "RUNOFF", true
		} else if rno_Groff_Matcher_0.Match(i) {
			return "Groff", true
		}

		return OtherLanguage, false
	},
	".rpy": func(i []byte) (string, bool) {
		if rpy_Python_Matcher_0.Match(i) {
			return "Python", true
		}

		return "Ren'Py", false
	},
	".rs": func(i []byte) (string, bool) {
		if rs_Rust_Matcher_0.Match(i) {
			return "Rust", true
		} else if rs_RenderScript_Matcher_0.Match(i) {
			return "RenderScript", true
		}

		return OtherLanguage, false
	},
	".sc": func(i []byte) (string, bool) {
		if sc_SuperCollider_Matcher_0.Match(i) || sc_SuperCollider_Matcher_1.Match(i) || sc_SuperCollider_Matcher_2.Match(i) {
			return "SuperCollider", true
		} else if sc_Scala_Matcher_0.Match(i) || sc_Scala_Matcher_1.Match(i) || sc_Scala_Matcher_2.Match(i) {
			return "Scala", true
		}

		return OtherLanguage, false
	},
	".sql": func(i []byte) (string, bool) {
		if sql_PLpgSQL_Matcher_0.Match(i) || sql_PLpgSQL_Matcher_1.Match(i) || sql_PLpgSQL_Matcher_2.Match(i) {
			return "PLpgSQL", true
		} else if sql_SQLPL_Matcher_0.Match(i) || sql_SQLPL_Matcher_1.Match(i) {
			return "SQLPL", true
		} else if sql_PLSQL_Matcher_0.Match(i) || sql_PLSQL_Matcher_1.Match(i) {
			return "PLSQL", true
		} else if sql_SQL_Matcher_0.Match(i) {
			return "SQL", true
		}

		return OtherLanguage, false
	},
	".srt": func(i []byte) (string, bool) {
		if srt_SubRipText_Matcher_0.Match(i) {
			return "SubRip Text", true
		}

		return OtherLanguage, false
	},
	".toc": func(i []byte) (string, bool) {
		if toc_WorldofWarcraftAddonData_Matcher_0.Match(i) {
			return "World of Warcraft Addon Data", true
		} else if toc_TeX_Matcher_0.Match(i) {
			return "TeX", true
		}

		return OtherLanguage, false
	},
	".ts": func(i []byte) (string, bool) {
		if ts_XML_Matcher_0.Match(i) {
			return "XML", true
		}

		return "TypeScript", false
	},
	".tst": func(i []byte) (string, bool) {
		if tst_GAP_Matcher_0.Match(i) {
			return "GAP", true
		}

		return "Scilab", false
	},
	".tsx": func(i []byte) (string, bool) {
		if tsx_TypeScript_Matcher_0.Match(i) {
			return "TypeScript", true
		} else if tsx_XML_Matcher_0.Match(i) {
			return "XML", true
		}

		return OtherLanguage, false
	},
}

var (
	asc_PublicKey_Matcher_0                = regexp.MustCompile(`/^(----[- ]BEGIN|ssh-(rsa|dss)) /`)
	asc_AsciiDoc_Matcher_0                 = regexp.MustCompile(`/^[=-]+(\s|\n)|{{[A-Za-z]/`)
	asc_AGSScript_Matcher_0                = regexp.MustCompile(`/^(\/\/.+|((import|export)\s+)?(function|int|float|char)\s+((room|repeatedly|on|game)_)?([A-Za-z]+[A-Za-z_0-9]+)\s*[;\(])/`)
	bb_BlitzBasic_Matcher_0                = regexp.MustCompile(`/^\s*; /`)
	bb_BlitzBasic_Matcher_1                = regexp.MustCompile(`End Function`)
	bb_BitBake_Matcher_0                   = regexp.MustCompile(`/^\s*(# |include|require)\b/`)
	builds_XML_Matcher_0                   = regexp.MustCompile(`/^(\s*)(<Project|<Import|<Property|<?xml|xmlns)/i`)
	ch_xBase_Matcher_0                     = regexp.MustCompile(`/^\s*#\s*(if|ifdef|ifndef|define|command|xcommand|translate|xtranslate|include|pragma|undef)\b/i`)
	cl_CommonLisp_Matcher_0                = regexp.MustCompile(`/^\s*\((defun|in-package|defpackage) /i`)
	cl_Cool_Matcher_0                      = regexp.MustCompile(`/^class/x`)
	cl_OpenCL_Matcher_0                    = regexp.MustCompile(`/\/\* |\/\/ |^\}/`)
	cls_TeX_Matcher_0                      = regexp.MustCompile(`/\\\w+{/`)
	cs_Smalltalk_Matcher_0                 = regexp.MustCompile(`/![\w\s]+methodsFor: /`)
	cs_CSharp_Matcher_0                    = regexp.MustCompile(`/^\s*namespace\s*[\w\.]+\s*{/`)
	cs_CSharp_Matcher_1                    = regexp.MustCompile(`/^\s*\/\//`)
	d_D_Matcher_0                          = regexp.MustCompile(`/^module /`)
	d_DTrace_Matcher_0                     = regexp.MustCompile(`/^((dtrace:::)?BEGIN|provider |#pragma (D (option|attributes)|ident)\s)/`)
	d_Makefile_Matcher_0                   = regexp.MustCompile(`/(\/.*:( .* \\)$| : \\$|^ : |: \\$)/`)
	ecl_ECLiPSe_Matcher_0                  = regexp.MustCompile(`/^[^#]+:-/`)
	ecl_ECL_Matcher_0                      = regexp.MustCompile(`:=`)
	es_Erlang_Matcher_0                    = regexp.MustCompile(`/^\s*(?:%%|main\s*\(.*?\)\s*->)/`)
	f_Forth_Matcher_0                      = regexp.MustCompile(`/^: /`)
	f_FilebenchWML_Matcher_0               = regexp.MustCompile(`flowop`)
	f_FORTRAN_Matcher_0                    = regexp.MustCompile(`/^([c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)/i`)
	for_Forth_Matcher_0                    = regexp.MustCompile(`/^: /`)
	for_FORTRAN_Matcher_0                  = regexp.MustCompile(`/^([c*][^abd-z]|      (subroutine|program|end|data)\s|\s*!)/i`)
	fr_Forth_Matcher_0                     = regexp.MustCompile(`/^(: |also |new-device|previous )/`)
	fr_Frege_Matcher_0                     = regexp.MustCompile(`/^\s*(import|module|package|data|type) /`)
	fs_Forth_Matcher_0                     = regexp.MustCompile(`/^(: |new-device)/`)
	fs_FSharp_Matcher_0                    = regexp.MustCompile(`/^\s*(#light|import|let|module|namespace|open|type)/`)
	fs_GLSL_Matcher_0                      = regexp.MustCompile(`/^\s*(#version|precision|uniform|varying|vec[234])/`)
	fs_Filterscript_Matcher_0              = regexp.MustCompile(`/#include|#pragma\s+(rs|version)|__attribute__/`)
	gs_Gosu_Matcher_0                      = regexp.MustCompile(`/^uses java\./`)
	h_ObjectiveDashC_Matcher_0             = regexp.MustCompile(`/^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])/`)
	h_CPlusPlus_Matcher_0                  = regexp.MustCompile(`/^\s*#\s*include <(cstdint|string|vector|map|list|array|bitset|queue|stack|forward_list|unordered_map|unordered_set|(i|o|io)stream)>/`)
	h_CPlusPlus_Matcher_1                  = regexp.MustCompile(`/^\s*template\s*</`)
	h_CPlusPlus_Matcher_2                  = regexp.MustCompile(`/^[ \t]*try/`)
	h_CPlusPlus_Matcher_3                  = regexp.MustCompile(`/^[ \t]*catch\s*\(/`)
	h_CPlusPlus_Matcher_4                  = regexp.MustCompile(`/^[ \t]*(class|(using[ \t]+)?namespace)\s+\w+/`)
	h_CPlusPlus_Matcher_5                  = regexp.MustCompile(`/^[ \t]*(private|public|protected):$/`)
	h_CPlusPlus_Matcher_6                  = regexp.MustCompile(`/std::\w+/`)
	inc_PHP_Matcher_0                      = regexp.MustCompile(`/^<\?(?:php)?/`)
	inc_POVDashRaySDL_Matcher_0            = regexp.MustCompile(`/^\s*#(declare|local|macro|while)\s/`)
	l_CommonLisp_Matcher_0                 = regexp.MustCompile(`/\(def(un|macro)\s/`)
	l_Lex_Matcher_0                        = regexp.MustCompile(`/^(%[%{}]xs|<.*>)/`)
	l_Groff_Matcher_0                      = regexp.MustCompile(`/^\.[a-z][a-z](\s|$)/i`)
	l_PicoLisp_Matcher_0                   = regexp.MustCompile(`/^\((de|class|rel|code|data|must)\s/`)
	ls_LoomScript_Matcher_0                = regexp.MustCompile(`/^\s*package\s*[\w\.\/\*\s]*\s*{/`)
	lsp_CommonLisp_Matcher_0               = regexp.MustCompile(`/^\s*\((defun|in-package|defpackage) /i`)
	lsp_NewLisp_Matcher_0                  = regexp.MustCompile(`/^\s*\(define /`)
	lisp_CommonLisp_Matcher_0              = regexp.MustCompile(`/^\s*\((defun|in-package|defpackage) /i`)
	lisp_NewLisp_Matcher_0                 = regexp.MustCompile(`/^\s*\(define /`)
	m_ObjectiveDashC_Matcher_0             = regexp.MustCompile(`/^\s*(@(interface|class|protocol|property|end|synchronised|selector|implementation)\b|#import\s+.+\.h[">])/`)
	m_Mercury_Matcher_0                    = regexp.MustCompile(`:- module`)
	m_MUF_Matcher_0                        = regexp.MustCompile(`/^: /`)
	m_M_Matcher_0                          = regexp.MustCompile(`/^\s*;/`)
	m_Mathematica_Matcher_0                = regexp.MustCompile(`/\*\)$/`)
	m_Matlab_Matcher_0                     = regexp.MustCompile(`/^\s*%/`)
	m_Limbo_Matcher_0                      = regexp.MustCompile(`/^\w+\s*:\s*module\s*{/`)
	md_Markdown_Matcher_0                  = regexp.MustCompile(`/(^[-a-z0-9=#!\*\[|>])|<\//i`)
	md_Markdown_Matcher_1                  = regexp.MustCompile(`/^$/`)
	md_GCCmachinedescription_Matcher_0     = regexp.MustCompile(`/^(;;|\(define_)/`)
	ml_OCaml_Matcher_0                     = regexp.MustCompile(`/(^\s*module)|let rec |match\s+(\S+\s)+with/`)
	ml_StandardML_Matcher_0                = regexp.MustCompile(`/=> |case\s+(\S+\s)+of/`)
	mod_XML_Matcher_0                      = regexp.MustCompile(`<!ENTITY `)
	mod_ModulaDash2_Matcher_0              = regexp.MustCompile(`/^\s*MODULE [\w\.]+;/i`)
	mod_ModulaDash2_Matcher_1              = regexp.MustCompile(`/^\s*END [\w\.]+;/i`)
	ms_Groff_Matcher_0                     = regexp.MustCompile(`/^[.'][a-z][a-z](\s|$)/i`)
	n_Groff_Matcher_0                      = regexp.MustCompile(`/^[.']/`)
	n_Nemerle_Matcher_0                    = regexp.MustCompile(`/^(module|namespace|using)\s/`)
	ncl_Text_Matcher_0                     = regexp.MustCompile(`THE_TITLE`)
	nl_NL_Matcher_0                        = regexp.MustCompile(`/^(b|g)[0-9]+ /`)
	php_Hack_Matcher_0                     = regexp.MustCompile(`<\?hh`)
	php_PHP_Matcher_0                      = regexp.MustCompile(`/<?[^h]/`)
	pl_Prolog_Matcher_0                    = regexp.MustCompile(`/^[^#]*:-/`)
	pl_Perl_Matcher_0                      = regexp.MustCompile(`/use strict|use\s+v?5\./`)
	pl_Perl6_Matcher_0                     = regexp.MustCompile(`/^(use v6|(my )?class|module)/`)
	pm_Perl_Matcher_0                      = regexp.MustCompile(`/use strict|use\s+v?5\./`)
	pm_Perl6_Matcher_0                     = regexp.MustCompile(`/^(use v6|(my )?class|module)/`)
	t_Perl_Matcher_0                       = regexp.MustCompile(`/use strict|use\s+v?5\./`)
	t_Perl6_Matcher_0                      = regexp.MustCompile(`/^(use v6|(my )?class|module)/`)
	pod_Pod_Matcher_0                      = regexp.MustCompile(`/^=\w+\b/`)
	pro_Prolog_Matcher_0                   = regexp.MustCompile(`/^[^#]+:-/`)
	pro_INI_Matcher_0                      = regexp.MustCompile(`last_client=`)
	pro_QMake_Matcher_0                    = regexp.MustCompile(`HEADERS`)
	pro_QMake_Matcher_1                    = regexp.MustCompile(`SOURCES`)
	pro_IDL_Matcher_0                      = regexp.MustCompile(`/^\s*function[ \w,]+$/`)
	props_XML_Matcher_0                    = regexp.MustCompile(`/^(\s*)(<Project|<Import|<Property|<?xml|xmlns)/i`)
	props_INI_Matcher_0                    = regexp.MustCompile(`/\w+\s*=\s*/i`)
	r_Rebol_Matcher_0                      = regexp.MustCompile(`/\bRebol\b/i`)
	r_R_Matcher_0                          = regexp.MustCompile(`/<-|^\s*#/`)
	rno_RUNOFF_Matcher_0                   = regexp.MustCompile(`/^\.!|^\.end lit(?:eral)?\b/i`)
	rno_Groff_Matcher_0                    = regexp.MustCompile(`/^\.\\" /`)
	rpy_Python_Matcher_0                   = regexp.MustCompile(`/(^(import|from|class|def)\s)/m`)
	rs_Rust_Matcher_0                      = regexp.MustCompile(`/^(use |fn |mod |pub |macro_rules|impl|#!?\[)/`)
	rs_RenderScript_Matcher_0              = regexp.MustCompile(`/#include|#pragma\s+(rs|version)|__attribute__/`)
	sc_SuperCollider_Matcher_0             = regexp.MustCompile(`/\^(this|super)\./`)
	sc_SuperCollider_Matcher_1             = regexp.MustCompile(`/^\s*(\+|\*)\s*\w+\s*{/`)
	sc_SuperCollider_Matcher_2             = regexp.MustCompile(`/^\s*~\w+\s*=\./`)
	sc_Scala_Matcher_0                     = regexp.MustCompile(`/^\s*import (scala|java)\./`)
	sc_Scala_Matcher_1                     = regexp.MustCompile(`/^\s*val\s+\w+\s*=/`)
	sc_Scala_Matcher_2                     = regexp.MustCompile(`/^\s*class\b/`)
	sql_PLpgSQL_Matcher_0                  = regexp.MustCompile(`/^\\i\b|AS \$\$|LANGUAGE '?plpgsql'?/i`)
	sql_PLpgSQL_Matcher_1                  = regexp.MustCompile(`/SECURITY (DEFINER|INVOKER)/i`)
	sql_PLpgSQL_Matcher_2                  = regexp.MustCompile(`/BEGIN( WORK| TRANSACTION)?;/i`)
	sql_SQLPL_Matcher_0                    = regexp.MustCompile(`/(alter module)|(language sql)|(begin( NOT)+ atomic)/i`)
	sql_SQLPL_Matcher_1                    = regexp.MustCompile(`/signal SQLSTATE '[0-9]+'/i`)
	sql_PLSQL_Matcher_0                    = regexp.MustCompile(`/\$\$PLSQL_|XMLTYPE|sysdate|systimestamp|\.nextval|connect by|AUTHID (DEFINER|CURRENT_USER)/i`)
	sql_PLSQL_Matcher_1                    = regexp.MustCompile(`/constructor\W+function/i`)
	sql_SQL_Matcher_0                      = regexp.MustCompile(`! /begin|boolean|package|exception/i`)
	srt_SubRipText_Matcher_0               = regexp.MustCompile(`/^(\d{2}:\d{2}:\d{2},\d{3})\s*(-->)\s*(\d{2}:\d{2}:\d{2},\d{3})$/`)
	toc_WorldofWarcraftAddonData_Matcher_0 = regexp.MustCompile(`/^## |@no-lib-strip@/`)
	toc_TeX_Matcher_0                      = regexp.MustCompile(`/^\\(contentsline|defcounter|beamer|boolfalse)/`)
	ts_XML_Matcher_0                       = regexp.MustCompile(`<TS`)
	tst_GAP_Matcher_0                      = regexp.MustCompile(`gap> `)
	tsx_TypeScript_Matcher_0               = regexp.MustCompile(`/^\s*(import.+(from\s+|require\()['"]react|\/\/\/\s*<reference\s)/`)
	tsx_XML_Matcher_0                      = regexp.MustCompile(`/^\s*<\?xml\s+version/i`)
)
