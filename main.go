package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jhphon0730/go-grep/internal/grep"
)

func main() {
	ext := flag.String("ext", "", "파일 확장자 필터 (예: .go, .txt)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}
	pattern := args[0]

	opts := grep.Options{
		Ext: *ext,
	}

	g := grep.NewGrep(pattern, opts)

	matches, err := g.GrepFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range matches {
		fmt.Printf("%s:%d: %s\n", m.FileName, m.LineNum, m.LineText)
	}
}
