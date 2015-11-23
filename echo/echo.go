package main

import (
	"flag" // command line option parser
	"os"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse() // scans the arg list and set up flag
	s := ""
	// for循环是Go中唯一的循环语句,循环体无花括号
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if !*omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
