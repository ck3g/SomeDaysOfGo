package main

import "testing"

// To check test coverage in HTML type run the following commands:
// $ go test --coverprofile c.out
// go tool cover -html=c.out

func TestMessageHasASingleArgument(t *testing.T) {
	actual := Message("message")
	expected := "message"

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMessageHasManyArguments(t *testing.T) {
	actual := Message("Here", "is", "the", "message")
	expected := "Here is the message"

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestMessageWithTableTests(t *testing.T) {
	type test struct {
		data     []string
		expected string
	}

	tests := []test{
		test{[]string{""}, ""},
		test{[]string{"message"}, "message"},
		test{[]string{"Here", "is", "the", "message"}, "Here is the message"},
	}

	for _, tst := range tests {
		actual := Message(tst.data...)
		if tst.expected != actual {
			t.Error("Expected", tst.expected, "got", actual)
		}
	}
}
