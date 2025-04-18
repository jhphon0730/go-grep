package grep

import (
	"bufio"
	"path/filepath"
	"os"
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
	GetGreps() []Match
	GrepFile(path string, f os.FileInfo, err error) (error)
	GrepFiles() ([]Match, error)
}

type grepper struct {
	pattern string
	greps []Match
	Options Options
}

func NewGrep(pattern string, options Options) Grepper {
	return &grepper{
		pattern: pattern,
		greps: []Match{},
		Options: options,
	}
}

func (g *grepper) GetGreps() []Match {
	return g.greps
}

func (g *grepper) GrepFile(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_number := 0
	re := CreateRegexp(g.Options, g.pattern)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			g.greps = append(g.greps, Match{
				FileName:  path,
				LineNum:   line_number,
				LineText:  line,
				MatchText: re.FindString(line),
			})
		}
		line_number++
	}

	return nil
}

func (g *grepper) GrepFiles() ([]Match, error) {
	g.greps = []Match{}
	if err := filepath.Walk(".", g.GrepFile); err != nil {
		return nil, err
	}

	return g.greps, nil
}
