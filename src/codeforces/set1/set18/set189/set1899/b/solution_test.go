package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{1, 2}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 2, 3, 6, 1, 3}
	expect := 9
	runSample(t, a, expect)
}
