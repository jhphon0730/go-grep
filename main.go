package main
import (
	"flag"
	"fmt"
	"log"

	"github.com/jhphon0730/go-grep/internal/utils"
	"github.com/jhphon0730/go-grep/internal/grep"
)

func main() {
	ext := flag.String("ext", "", "파일 확장자 필터 (예: .go, .txt)")
	line_number := flag.Bool("line-number", false, "라인 번호 출력")
	ignoreCase := flag.Bool("ignore-case", false, "대소문자 무시")
	word_match := flag.Bool("word-match", false, "단어 단위로 매칭")

	flag.Parse()

	pattern := flag.Arg(0)
	if pattern == "" {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}

	g := grep.NewGrep(pattern, grep.Options{
		Ext: *ext,
		LineNumber: *line_number,
		IgnoreCase: *ignoreCase,
		WordMatch: *word_match,
	})

	matches, err := g.GrepFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		text := utils.Highlight(match.LineText, match.MatchText)
		filename := utils.Colorize(match.FileName, utils.ColorBlue)

		if g.GetOptions().LineNumber {
			linenum := utils.Colorize(utils.IntToStr(match.LineNum), utils.ColorGreen)
			fmt.Printf("%s:%s:%s\n", filename, linenum, text)
		} else {
			fmt.Printf("%s:%s\n", filename, text)
		}
	}
}
