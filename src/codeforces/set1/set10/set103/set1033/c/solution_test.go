package main

import "testing"

func runSample(t *testing.T, a []int, expect string) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 6, 5, 4, 2, 7, 1, 8}
	expect := "BAAAABAB"
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 11, 2, 5, 10, 9, 7, 13, 15, 8, 4, 12, 6, 1, 14}
	expect := "ABAAAABBBAABAAB"
	runSample(t, a, expect)
}
