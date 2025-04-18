package grep

import (
	"regexp"
)

type Options struct {
	Ext string
}

type Match struct {
	FileName  string
	LineNum   int
	LineText  string
	MatchText string // optional: 색상 강조용
}

func CreateRegexp(options Options, pattern string) *regexp.Regexp {
	flags := ""

	re := regexp.MustCompile(flags + pattern)
	return re
}
