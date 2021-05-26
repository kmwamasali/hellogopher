package stringutils

import "testing"

func TestUpper(t *testing.T) {
	if Upper("abc") != "ABC" {
		t.Error("expected 'abc' to equal 'ABC")
	}
}

func TestLower(t *testing.T) {
	if Lower("ABC") != "abc" {
		t.Error("expected 'ABC' to equal 'abc'")
	}
}