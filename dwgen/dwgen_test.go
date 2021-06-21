package dwgen

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name    string
		in      *Config
		want    *Generator
		wantErr error
	}{
		{"nil config", nil, &Generator{&DefaultConfig, WordListShorthands["l"]}, nil},
		{"nonexistent source list", func() *Config { cfg := DefaultConfig; cfg.SourceList = "nonexistent"; return &cfg }(), nil, ErrInvalidWordList},
		{"manual config", &Config{8, "u3cp", "-"}, &Generator{&Config{8, "u3cp", "-"}, &UniquePrefixShortWordList}, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := New(c.in)
			if !errors.Is(err, c.wantErr) {
				t.Fatalf("New(%v) error == %v, want %v", c.in, err, c.wantErr)
			}
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("New(%v) == %v, want %v", c.in, got, c.want)
			}
		})
	}
}
