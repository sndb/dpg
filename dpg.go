package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sndb/dpg/dwgen"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	wordCount := flag.Int("c", 6, "word count")
	sourceList := flag.String("l", "l", "source word list ('l': long word list (for use with five dice), 's': general short word list (for use with four dice), 'u3cp': short word list (with words that have unique three-character prefixes)")
	delimiter := flag.String("d", " ", "delimiter string")
	number := flag.Int("n", 1, "number of passphrases to generate")
	flag.Parse()

	config := dwgen.Config{
		WordCount:  *wordCount,
		SourceList: *sourceList,
		Delimiter:  *delimiter,
	}
	g, err := dwgen.New(&config)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *number; i++ {
		p, err := g.Generate()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}
