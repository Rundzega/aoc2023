package aoc2023

import (
	"testing"
)

func TestCalibrate(t *testing.T) {
    want := 142
    actual := calibrate("data/day1_test.txt")
    if want != actual {
        t.Fatalf("Wanted %v and got %v", want, actual)
    }
}

