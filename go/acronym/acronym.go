/*
Package acronym implement Acronym exercise solution in https://exercism.io/
*/
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate convert a phrase to its acronym
func Abbreviate(s string) string {
	s = strings.ToUpper(s)

	re := regexp.MustCompile(`[ _,-]`)
	words := re.Split(s, -1)

	abbr := ""
	for _, word := range words {
		if word != "" {
			abbr = abbr + string(word[0])
		}
	}
	return abbr
}
