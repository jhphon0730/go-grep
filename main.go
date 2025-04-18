package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jhphon0730/go-grep/internal/grep"
)

func main() {
	ext := flag.String("ext", "", "파일 확장자 필터 (예: .go, .txt)")
	line := flag.Bool("line", false, "라인 번호 출력")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}
	pattern := args[0]

	g := grep.NewGrep(pattern, grep.Options{
		Ext: *ext,
		Line: *line,
	})

	matches, err := g.GrepFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range matches {
		if g.GetOptions().Line {
			fmt.Println(match.FileName + ":", match.LineNum, match.LineText)
		} else {
			fmt.Println(match.FileName, match.LineText)
		}
	}
}
