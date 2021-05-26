package stringutils

import "testing"

func TestTableUpper(t *testing.T) {
	var stringTests = []struct {
		input string
		expected string
	}{
			{"abc", "ABC"},
			{"cde", "CDE"},
			{"def", "DEF"},
			{"ghi", "GHI"},
	}

	for _, test := range stringTests {
		if output := Upper(test.input);output != test.expected {
			t.Error("expected {} to equal {}", test.input, test.expected)
		}
	}
}

func TestTableLower(t *testing.T) {
	var stringTests = []struct {
		input string
		expected string
	}{
			{"ABC", "abc"},
			{"CDE", "cde"},
			{"DEF", "def"},
			{"GHI", "ghi"},
	}

	for _, test := range stringTests {
		if output := Lower(test.input);output != test.expected {
			t.Error("expected {} to equal {}", test.input, test.expected)
		}
	}
}