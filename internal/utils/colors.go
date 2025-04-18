package utils

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[1;31m"
	ColorGreen  = "\033[1;32m"
	ColorBlue   = "\033[1;34m"
)

func Colorize(text, color string) string {
	return color + text + ColorReset
}
