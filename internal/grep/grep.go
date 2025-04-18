package grep

type Options struct {
	WordMatch  bool
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

func (g grepper) GrepFile(filename string, pattern string) ([]Match, error) { }

func (g grepper) GrepFiles(files []string, pattern string) ([]Match, error) { }
