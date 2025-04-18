package grep

import (
	"regexp"
)

type Options struct {
	Pattern string
	Ext string
	LineNumber bool
	IgnoreCase bool
}

type Match struct {
	FileName  string
	LineNum   int
	LineText  string
	MatchText string // optional: 색상 강조용
}

func CreateRegexp(options Options, pattern string) *regexp.Regexp {
	flags := ""

	if options.IgnoreCase {
		flags += "(?i)"
	}

	re := regexp.MustCompile(flags + pattern)
	return re
}
