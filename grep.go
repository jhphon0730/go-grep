package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jhphon0730/go-grep/internal/utils"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Usage: gitgrep-lite [options] <pattern> [-- pathspecs...]")
	}

	pattern, file_list := utils.ParseArgs(args)
	fmt.Println("Pattern:", pattern)
	fmt.Println("FileList", file_list)
}
