package aoc2023

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
    want := 8
    actual := game(12, 13, 14, "data/day2_test.txt")
    if actual != want {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}

func TestRunGame(t *testing.T) {
    result := game(12, 13, 14, "data/day2_puzzle.txt")
    fmt.Printf("%v \n", result)
}
