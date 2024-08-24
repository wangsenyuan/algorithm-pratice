package main

import "testing"

func runSample(t *testing.T, a []int, expect string) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2}
	expect := "WIN"
	runSample(t, a, expect)
}
