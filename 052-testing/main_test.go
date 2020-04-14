package main

import "testing"

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
