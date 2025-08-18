package main

import (
	"bytes"
	"strings"
	"testing"
)

func runGrepOutput(input string, opts options) string {
	var buf bytes.Buffer
	old := out
	out = &buf
	defer func() { out = old }()

	r := strings.NewReader(input)
	runGrep(r, opts)
	return buf.String()
}

func TestGrep(t *testing.T) {
	text := "alpha\nBeta\ngamma\nALPHA\nbeta\nGamma\n"

	tests := []struct {
		name   string
		opts   options
		expect string
	}{
		{"fixed match", options{pattern: "alpha", fixed: true}, "alpha\n"},
		{"ignore case", options{pattern: "alpha", fixed: true, ignore: true}, "alpha\nALPHA\n"},
		{"regex match", options{pattern: "B.*", fixed: false}, "Beta\n"},
		{"invert match", options{pattern: "alpha", fixed: true, invert: true}, "Beta\ngamma\nALPHA\nbeta\nGamma\n"},
		{"line numbers", options{pattern: "alpha", fixed: true, lineNum: true}, "1:alpha\n"},
		{"before context", options{pattern: "gamma", fixed: true, before: 1}, "Beta\ngamma\n"},
		{"after context", options{pattern: "Beta", fixed: true, after: 1}, "Beta\ngamma\n"},
		{"before and after context", options{pattern: "Beta", fixed: true, before: 1, after: 1}, "alpha\nBeta\ngamma\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := runGrepOutput(text, tt.opts)
			if out != tt.expect {
				t.Errorf("expected %q, got %q", tt.expect, out)
			}
		})
	}
}
