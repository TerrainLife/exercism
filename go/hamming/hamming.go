/*
Package hamming implement Hamming exercise solution in https://exercism.io/
*/
package hamming

import (
	"errors"
)

// Distance calc Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	if (a == "" && b != "") || (a != "" && b == "") {
		return 0, errors.New("Error - empty stand(s)!")
	}
	if len(a) != len(b) {
		return 0, errors.New("Error - stands have different len!")
	}

	difs := 0
	for i := range a {
		if a[i] != b[i] {
			difs++
		}
	}
	return difs, nil
}
