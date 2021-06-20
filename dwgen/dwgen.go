package dwgen

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
)

// WordListShorthands is the map of the list shorthands to the list pointers
var WordListShorthands = map[string]*WordList{
	"l":    &LongWordList,
	"s":    &ShortWordList,
	"u3cp": &UniquePrefixShortWordList,
}

// DefaultConfig is the default Config used by Generator
var DefaultConfig = Config{
	WordCount:  6,
	SourceList: "l",
	Delimiter:  " ",
}

// ErrInvalidWordList is returned by New to indicate that an invalid word list
// shorthand is used
var ErrInvalidWordList = errors.New("not a valid word list")

// WordList holds the list of diceware words and the number of dice to be
// thrown for indexing the list
type WordList struct {
	List       map[int]string
	DiceNumber int
}

// Config holds the settings for Generator
type Config struct {
	// WordCount is the amount of words used to generate passphrase
	WordCount int

	// SourceList is the word list used to get words from
	SourceList string

	// Delimiter is the delimiter between separate words
	Delimiter string
}

// Generator uses Config and WordList to generate passphrases
type Generator struct {
	*Config
	*WordList
}

// New returns a new Generator
func New(config *Config) (*Generator, error) {
	if config == nil {
		config = &DefaultConfig
	}
	wordList, ok := WordListShorthands[config.SourceList]
	if !ok {
		return nil, fmt.Errorf("%q: %w", config.SourceList, ErrInvalidWordList)
	}
	g := Generator{
		Config:   config,
		WordList: wordList,
	}
	return &g, nil
}

// Generate generates a new passphrase
func (g *Generator) Generate() (string, error) {
	var p []string
	for i := 0; i < g.Config.WordCount; i++ {
		p = append(p, g.WordList.List[Dice(g.WordList.DiceNumber)])
	}
	return strings.Join(p, g.Config.Delimiter), nil
}

// Dice throws n dice
func Dice(n int) int {
	d := 0
	for i := 0; i < n; i++ {
		d *= 10
		x := r(6) + 1
		d += int(x)
	}
	return d
}

func r(max int) int64 {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		log.Fatal(err)
	}
	return x.Int64()
}
