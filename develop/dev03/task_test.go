package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
		{"\\", "", true},
	}

	for _, test := range tests {
		result, err := Parse(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("expected error for input %q, but got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect error for input %q, but got %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("expected %q for input %q, but got %q", test.expected, test.input, result)
			}
		}
	}
}
