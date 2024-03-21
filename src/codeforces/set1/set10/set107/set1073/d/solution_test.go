package main

import "testing"

func runSample(t *testing.T, a []int, tot int, expect int) {
	res := solve(a, tot)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 2, 5}
	tot := 38
	expect := 10
	runSample(t, a, tot, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 100, 2, 6}
	tot := 21
	expect := 6
	runSample(t, a, tot, expect)
}
