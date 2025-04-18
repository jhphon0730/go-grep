package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jhphon0730/go-grep/internal/grep"
)

func main() {
	ext := flag.String("ext", "", "파일 확장자 필터 (예: .go, .txt)")
	line_number := flag.Bool("line-number", false, "라인 번호 출력")
	ignoreCase := flag.Bool("ignore-case", false, "대소문자 무시")

	flag.Parse()

	pattern := flag.Arg(0)
	if pattern == "" {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}

	g := grep.NewGrep(pattern, grep.Options{
		Ext: *ext,
		LineNumber: *line_number,
		IgnoreCase: *ignoreCase,
	})

	matches, err := g.GrepFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		if g.GetOptions().LineNumber {
			fmt.Printf("%s:%d:%s\n", match.FileName, match.LineNum, match.LineText)
		} else {
			fmt.Printf("%s:%s\n", match.FileName, match.LineText)
		}
	}
}
