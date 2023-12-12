package aoc2023

import (
	"testing"
)

func TestScratchcards(t *testing.T) {
    want := 13
    actual := scratchcards("data/day4_test.txt")
    if want != actual {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}

func TestMadness(t *testing.T) {
    want := 30
    actual := madness("data/day4_test.txt")
    if want != actual {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}
