package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	a := []int{2, 4, 1, 3}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 9, 1, 9, 6}
	expect := 12
	runSample(t, a, expect)
}
