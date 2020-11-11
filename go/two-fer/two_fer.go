/*
Package twofer implement Two-fer exercise solution in https://exercism.io/
*/
package twofer

// ShareWith shares with name
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
