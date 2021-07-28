package main

import (
	"flag"
	"fmt"

	"github.com/sndb/dpg/dwgen"
)

func main() {
	c := flag.Int("c", 6, "word count")
	d := flag.String("d", " ", "delimiter string")
	n := flag.Int("n", 1, "number of passphrases to generate")
	flag.Parse()

	config := dwgen.Config{
		WordCount: *c,
		Delimiter: *d,
	}
	g := dwgen.New(&config)

	for i := 0; i < *n; i++ {
		p := g.Generate()
		fmt.Println(p)
	}
}
