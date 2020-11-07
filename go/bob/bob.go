/*
Package bob implement Bob exercise solution in https://exercism.io/
*/
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
	isEmpty := remark == ""

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

// Hey is a Bob`s answer on remark
func Hey(remark string) string {
	return getAnswers()[categorizeRemark(remark)]
}
