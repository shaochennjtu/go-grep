package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	
	"github.com/shaochennjtu/go-grep/"
)

var regexpFlag = flag.Bool("e", false, "Interpret pattern as an regular expression")

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
