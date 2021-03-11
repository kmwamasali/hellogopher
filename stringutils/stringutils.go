package stringutils

import "strings"

// Upper returns the uppercase of a given string argument
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower returns the lowercase of a given string argument
func Lower(s string) string {
	return strings.ToLower(s)
}
