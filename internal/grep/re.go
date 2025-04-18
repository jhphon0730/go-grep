package grep

import (
	"regexp"
)

func CreateRegexp(options Options, pattern string) *regexp.Regexp {
	flags := ""

	re := regexp.MustCompile(flags + pattern)
	return re
}
