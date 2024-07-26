package main

import "testing"

func runSample(t *testing.T, a []int, expect string) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 1, 3, 7, 4}
	expect := "111111"
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	expect := "10101"
	runSample(t, a, expect)
}
