package main

import (
	"bytes"
	"testing"
)

func TestCat(t *testing.T) {
	const expected = `Roses are red,
Go is awesome,
Its you I love.
`
	input := bytes.NewBufferString("Roses are red,\nGo is awesome,\nIts you I love.")

	actual := cat(input, false)

	if actual != expected {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestCatWithLineNumber(t *testing.T) {
	const expected = `1  Roses are red,
2  Go is awesome,
3  Its you I love.
`
	input := bytes.NewBufferString("Roses are red,\nGo is awesome,\nIts you I love.")

	actual := cat(input, true)

	if actual != expected {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
