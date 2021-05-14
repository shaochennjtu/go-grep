package main

import (
	"flag"
	"fmt"
	"os"
	"log"
		
	"github.com/shaochennjtu/go-grep/"
)

var h = flag.Bool("h", false, "go-grep -h")

Usage
go-grep [-e] [pattern] [file]


func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", internal.UsageText)
		flag.PrintDefaults()
	}
	log.SetLevel(log.InfoLevel)
	flag.Parse()
}

func main() {
	internal.Grep(
		[]byte(flag.Arg(0)),
		internal.ReadData(),
		*regexpFlag)
}
