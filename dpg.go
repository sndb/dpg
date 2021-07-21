package main

import (
	"flag"
	"fmt"

	"github.com/sndb/dpg/dwgen"
)

func main() {
	count := flag.Int("c", 6, "word count")
	delimiter := flag.String("d", " ", "delimiter string")
	number := flag.Int("n", 1, "number of passphrases to generate")
	flag.Parse()

	config := dwgen.Config{
		WordCount: *count,
		Delimiter: *delimiter,
	}
	g := dwgen.New(&config)

	for i := 0; i < *number; i++ {
		p := g.Generate()
		fmt.Println(p)
	}
}
