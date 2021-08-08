package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sndb/dpg/dwgen"
)

var (
	c int
	d string
	n int
)

func main() {
	flag.StringVar(&d, "d", " ", "delimiter string")
	flag.IntVar(&n, "n", 1, "number of passphrases to generate")
	flag.Parse()
	tail := flag.Args()
	switch len(tail) {
	case 0:
		c = 6
	case 1:
		i, err := strconv.Atoi(tail[0])
		if err != nil {
			flag.CommandLine.Usage()
		}
		c = i
	default:
		flag.CommandLine.Usage()
	}

	config := dwgen.Config{
		WordCount: c,
		Delimiter: d,
	}
	g := dwgen.New(&config)

	for i := 0; i < n; i++ {
		p := g.Generate()
		fmt.Println(p)
	}
}
