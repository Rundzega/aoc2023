package aoc2023

import (
	"testing"
)

func TestGame(t *testing.T) {
    want := 8
    actual := game(12, 13, 14, "data/day2_test.txt")
    if actual != want {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}

func TestMinimumCubes(t *testing.T) {
    want := 2286
    actual := minimumCubes("data/day2_test.txt")
    if actual != want {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}

