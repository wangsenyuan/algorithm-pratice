package main

import "testing"

func runSample(t *testing.T, a []int, expect string) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 5, 1, 3}
	expect := "Petr"
	runSample(t, a, expect)
}
