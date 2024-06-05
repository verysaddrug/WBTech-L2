package main

import (
	"reflect"
	"testing"
)

func TestSortByKey(t *testing.T) {
	testCases := []struct {
		input    []string
		key      int
		expected []string
	}{
		{[]string{"c b a", "3 2 1", "d e f"}, 1, []string{"3 2 1", "c b a", "d e f"}},
	}

	for _, tc := range testCases {
		*columnFlag = tc.key
		output := sortLines(tc.input)
		result := getField(output, tc.key)
		expected := getField(tc.expected, tc.key)
		if result != expected {
			t.Errorf("Expected sorted column %d to be %s, but got %s", tc.key, expected, result)
		}
	}
}

func TestSortByNumericValue(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"10", "3", "2", "1"}, []string{"1", "2", "3", "10"}},
		{[]string{"100", "20", "3", "1000"}, []string{"3", "20", "100", "1000"}},
		{[]string{"5", "10", "15", "20"}, []string{"5", "10", "15", "20"}},
	}

	for _, tc := range testCases {
		*numFlag = true
		output := sortLines(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Errorf("Expected sorted by numeric value to be %v, but got %v", tc.expected, output)
		}
	}
}

func TestSortInReverseOrder(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"1", "2", "3"}, []string{"3", "2", "1"}},
		{[]string{"world", "hello"}, []string{"hello", "world"}},
	}

	for _, tc := range testCases {
		reverse(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Expected sorted in reverse order to be %v, but got %v", tc.expected, tc.input)
		}
	}
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{"hello", "world", "hello", "world"}, []string{"hello", "world"}},
		{[]string{"1", "2", "3", "2", "1", "3"}, []string{"1", "2", "3"}},
		{[]string{"", "", "", ""}, []string{""}},
	}

	for _, tc := range testCases {
		result := unique(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected unique(%v) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
