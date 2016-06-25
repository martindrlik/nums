// nums provides conversion between standard positional numeral systems.
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	dst = flag.Int("target-base", 16, "")
	src = flag.Int("source-base", 10, "")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
	}
	for _, arg := range args {
		d := ParseDecimal(arg, *src)
		s := d.Conv(*dst)
		fmt.Println(prnt(arg, *src), "=", prnt(s, *dst))
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: nums [-target-base b1] [-source-base b2] n1 [n2 ...]\n")
	fmt.Fprintf(os.Stderr, "  n1, n2, etc.: number of base b2 numeral system\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  english alphabeth is used to represent\n")
	fmt.Fprintf(os.Stderr, "  numbers in numeral base greater than 10\n")
	os.Exit(2)
}

func prnt(s string, base int) string {
	if base == 10 {
		return s
	}
	return fmt.Sprintf("(%s)%d", s, base)
}
