package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, B []string, expect int) {
	res := solve(B)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	B := "WWWBWWBW\nBBBBBBBB\nWWWBWWBW\nWWWBWWBW\nWWWBWWBW\nWWWBWWBW\nWWWBWWBW\nWWWBWWBW"
	board := strings.Split(B, "\n")
	expect := 3
	runSample(t, board, expect)
}
