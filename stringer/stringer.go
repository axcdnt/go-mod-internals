package stringer

import "strings"

// Tokenize returns an array of strings splitted by space
func Tokenize(data string) []string {
	return strings.Split(data, " ")
}
