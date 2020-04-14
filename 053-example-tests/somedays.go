// Package somedays experiments with example testing
package somedays

// Message converts the list of arguments into a string
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
