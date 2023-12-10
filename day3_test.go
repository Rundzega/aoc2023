package aoc2023

import (
	"testing"
)

func TestFixEngine(t *testing.T) {
    want := 4361
    actual := fixEngine("data/day3_test.txt")
    if want != actual {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}
