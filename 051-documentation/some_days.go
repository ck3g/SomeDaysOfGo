// Copyright 2020 Author etc.

// Package somedays experiments with the documentation
package somedays

// To a package documentation to godoc.org, one need to input the link to the package source code
// at the search field

// Message converts a list of string arguments into a string
// The documentation of a function should start with a function name
// as you can see on the first line
func Message(m ...string) string {
	var message string
	for i, substr := range m {
		if i == 0 {
			message += substr
		} else {
			message += " " + substr
		}
	}

	return message
}
