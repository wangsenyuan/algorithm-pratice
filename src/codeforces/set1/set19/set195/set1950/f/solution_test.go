package main

import "testing"

func runSample(t *testing.T, a, b, c int) {
	expect := solve1(a, b, c)
	res := solve(a, b, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b, c := 8, 17, 9
	runSample(t, a, b, c)
}
