package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"Пятак", "пятка", "тяпка", "птк", "листок", "слиток", "столиК", "листок"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input:    []string{"лиса", "ёж", "змея"},
			expected: map[string][]string{},
		},
		{
			input: []string{"кот", "змея", "мезя", "змея"},
			expected: map[string][]string{
				"змея": {"змея", "мезя"},
			},
		},
	}

	for _, tc := range testCases {
		result := findAnagramSets(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected: %v, got: %v", tc.expected, result)
		}
	}
}
