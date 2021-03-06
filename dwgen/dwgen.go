package dwgen

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// DefaultConfig is the default Config used by Generator
var DefaultConfig = &Config{
	WordCount: 6,
	Delimiter: " ",
}

// Config holds the settings for Generator
type Config struct {
	// WordCount is the amount of words used to generate passphrase
	WordCount int

	// Delimiter is the delimiter between separate words
	Delimiter string
}

// Generator uses Config to generate passphrases
type Generator struct {
	*Config
}

// New returns a new Generator
func New(config *Config) *Generator {
	if config == nil {
		config = DefaultConfig
	}
	g := Generator{
		Config: config,
	}
	return &g
}

// Generate generates a new passphrase
func (g *Generator) Generate() string {
	var p []string
	for i := 0; i < g.Config.WordCount; i++ {
		p = append(p, wordList[rollDice(diceNum)])
	}
	return strings.Join(p, g.Config.Delimiter)
}

func rollDice(n int) int {
	d := 0
	for i := 0; i < n; i++ {
		d *= 10
		d += randInt(6) + 1
	}
	return d
}

func randInt(max int) int {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic("dwgen: can't get a random number: " + err.Error())
	}
	return int(x.Int64())
}
