package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jhphon0730/go-grep/internal/utils"
	"github.com/jhphon0730/go-grep/internal/grep"
)

func main() {
	g := grep.NewGrep("grep", grep.Options{})
	greps, err := g.GrepFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range greps {
		fmt.Printf("%s:%d:%s\n", m.FileName, m.LineNum, m.LineText)
	}

	return
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}

	pattern, file_list := utils.ParseArgs(args)
	fmt.Fprintln(os.Stderr, "Pattern", pattern)
	fmt.Fprintln(os.Stderr, "File list", file_list)
}
