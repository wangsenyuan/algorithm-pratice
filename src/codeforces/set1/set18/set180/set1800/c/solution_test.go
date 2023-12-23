package main

import "testing"

func runSample(t *testing.T, cards []int, expect int) {
	res := solve(cards)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cards := []int{3, 3, 3, 0, 0}
	expect := 6
	runSample(t, cards, expect)
}
