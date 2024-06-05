package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
	}

	for _, test := range tests {
		result, err := UnpackString(test.input)
		if (err != nil) != test.err {
			t.Errorf("Input: %v. Expected error: %v, got: %v", test.input, test.err, err)
		}
		if result != test.expected {
			t.Errorf("Input: %v. Expected: %v, got: %v", test.input, test.expected, result)
		}
	}
}
