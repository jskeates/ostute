package main

import "testing"

func TestNotYetImplementedMessageIsCorrect(t *testing.T) {
	expected := "Not yet implemented, sorry!"
	if NotYetImplementedMessage != expected {
		t.Fatalf("NotYetImplementedMessage has wrong text. Expected \"%s\" but was \"%s\".", expected, NotYetImplementedMessage)
	}
}
