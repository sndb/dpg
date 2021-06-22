package dwgen

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name string
		in   *Config
		want *Generator
	}{
		{"nil config", nil, &Generator{&DefaultConfig}},
		{"manual config", &Config{8, "-"}, &Generator{&Config{8, "-"}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := New(c.in)
			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("New(%v) == %v, want %v", c.in, got, c.want)
			}
		})
	}
}
