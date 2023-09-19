package main

import "testing"

func runSample(t *testing.T, n int, m int, cards []int, expect bool) {
	res := solve(n, m, len(cards), cards)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	m := 3
	cards := []int{3, 6, 4, 1, 2, 5}
	expect := true
	runSample(t, n, m, cards, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	m := 3
	cards := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := false
	runSample(t, n, m, cards, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	m := 4
	cards := []int{10, 4, 9, 3, 5, 6, 8, 2, 7, 1}
	expect := true
	runSample(t, n, m, cards, expect)
}
