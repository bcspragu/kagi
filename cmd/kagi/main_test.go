package main

import (
	"bytes"
	"testing"
)

func TestStreamAndRemoveDoubleNewlines(t *testing.T) {
	inp := `Hello

I am a string with

some double newlines


And a triple newline!`

	var buf bytes.Buffer
	if err := streamAndRemoveDoubleNewlines(inp, &buf); err != nil {
		t.Fatalf("streamAndRemoveDoubleNewlines: %v", err)
	}

	want := `Hello
I am a string with
some double newlines
And a triple newline!`

	got := buf.String()

	if got != want {
		t.Errorf("unexpected output = %q, want %q", got, want)
	}
}
