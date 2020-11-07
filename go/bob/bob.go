// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

const (
	regularQuestion = iota // "How are you?"
	yell                   // YELL AT HIM (in all capitals)
	yellQuestion           // yell a question
	nothing                // address him without actually saying anything
	other                  // anything else
)

func getAnswers() map[int]string {
	return map[int]string{
		regularQuestion: "Sure.",
		yell:            "Whoa, chill out!",
		yellQuestion:    "Calm down, I know what I'm doing!",
		nothing:         "Fine. Be that way!",
		other:           "Whatever.",
	}
}

func categorizeRemark(remark string) int {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isYell := (strings.ToUpper(remark) == remark) && (strings.ToUpper(remark) != strings.ToLower(remark))
	isEmpty := len(remark) == 0

	if isEmpty {
		return nothing
	}
	if isQuestion && isYell {
		return yellQuestion
	}
	if isQuestion {
		return regularQuestion
	}
	if isYell {
		return yell
	}

	return other
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	return getAnswers()[categorizeRemark(remark)]
}
