package grep

import (
	"bufio"
	"os"
	"regexp"
)

type Options struct {
}

type Match struct {
	FileName  string
	LineNum   int
	LineText  string
	MatchText string // optional: 색상 강조용
}

type Grepper interface {
	GrepFile(filename string, pattern string) ([]Match, error)
	GrepFiles(files []string, pattern string) ([]Match, error)
}

type grepper struct {
	Options Options
}

func NewGrep(options Options) Grepper {
	return grepper{
		Options: options,
	}
}

func (g grepper) GrepFile(filename string, pattern string) ([]Match, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matches []Match
	scanner := bufio.NewScanner(file)
	flags := ""

	re := regexp.MustCompile(flags + pattern)
	line_number := 0

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			matches = append(matches, Match{
				FileName:  filename,
				LineNum:   line_number,
				LineText:  line,
				MatchText: re.FindString(line),
			})
		}
		line_number ++
	}

	return matches, nil
}

func (g grepper) GrepFiles(files []string, pattern string) ([]Match, error) {}
