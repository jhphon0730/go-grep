package grep

import (
	"regexp"
)

type Options struct {
	Ext string
	LineNumber bool
	IgnoreCase bool
	WordMatch bool
}

type Match struct {
	FileName  string
	LineNum   int
	LineText  string
	MatchText string // optional: 색상 강조용
}

func CreateRegexp(options Options, pattern string) *regexp.Regexp {
	if options.WordMatch {
		pattern = "\\b" + pattern + "\\b"
	}
	if options.IgnoreCase {
		pattern = "(?i)" + pattern
	}
	return regexp.MustCompile(pattern)
}
