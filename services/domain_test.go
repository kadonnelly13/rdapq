package services

import (
	"testing"
)

func TestParseDomain(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"simple domain", "google.com", "com"},
		{"subdomain", "www.google.com", "com"},
		{"url", "https://www.google.com", "com"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			r := parseDomain(test.input)
			if r != test.output {
				t.Errorf(`parseDomain(%q) = %q, want %q`, test.input, r, test.output)
			}
		})
	}
}
