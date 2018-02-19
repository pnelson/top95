// top95 tests strings against a popular wordlist.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pnelson/top95"
)

var (
	h = flag.Bool("h", false, "show this usage information")
	o = flag.String("o", "", "output filename")
	v = flag.String("v", "wordlist", "variable name")
	p = flag.String("p", "top95", "package name")
	n = flag.Int("n", 1, "minimum word length")
	s = flag.String("s", defaultSource, "wordlist source url or path")
	t = flag.Duration("t", 90*time.Second, "timeout to fetch remote wordlist")
)

// https://github.com/berzerk0/Probable-Wordlists
const defaultSource = "https://raw.githubusercontent.com/berzerk0/Probable-Wordlists/master/Real-Passwords/Top95Thousand-probable.txt"

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [WORD]...\n\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	if *o != "" {
		w := wordlist{
			filename: *o,
			pkgname:  *p,
			varname:  *v,
			min:      *n,
			source:   *s,
			timeout:  *t,
		}
		err := w.generate()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
		return
	}
	ok := include(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
}

func include(ss []string) bool {
	ok := true
	for _, s := range ss {
		if top95.Include(s) {
			fmt.Fprintln(os.Stderr, s)
			ok = false
		}
	}
	return ok
}
