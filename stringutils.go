package stringutils

import "strings"

// Upper returns the uppercase of the given string argument.
func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
  return strings.ToLower(s)
}
