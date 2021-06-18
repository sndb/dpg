package dwgen

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"
)

func r(max int) int64 {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return x.Int64()
}

// WordList holds a list of diceware words and the number of dice to be thrown
// for indexing the list
type WordList struct {
	List    map[int]string
	DiceNum int
}

// Config holds parameters for Generator
type Config struct {
	WordCount  int
	SourceList string
	Delimiter  string
}

// Generator generates passphrases
type Generator struct {
	*Config
	*WordList
}

// New creates a new Generator
func New(config *Config) (*Generator, error) {
	var l *WordList
	switch config.SourceList {
	case "l":
		l = &LongWordList
	case "s":
		l = &ShortWordList
	case "u3cp":
		l = &UniquePrefixShortWordList
	default:
		return nil, fmt.Errorf("%q is not a valid word list", config.SourceList)
	}
	g := Generator{
		Config:   config,
		WordList: l,
	}
	return &g, nil
}

// Generate generates a new passphrase
func (g *Generator) Generate() (string, error) {
	var p []string
	for i := 0; i < g.Config.WordCount; i++ {
		p = append(p, g.WordList.List[Dice(g.WordList.DiceNum)])
	}
	return strings.Join(p, g.Config.Delimiter), nil
}

// Dice throws n dice
func Dice(n int) int {
	var d int
	for i := 0; i < n; i++ {
		d *= 10
		x := r(6) + 1
		d += int(x)
	}
	return d
}
