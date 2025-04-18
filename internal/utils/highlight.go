package utils

import (
	"strings"
)

func Highlight(lineText, matchText string) string {
	const highlightStart = "\033[1;31m" // Red
	const highlightEnd = "\033[0m"       // Reset

	return strings.ReplaceAll(lineText, matchText, highlightStart+matchText+highlightEnd)
}

