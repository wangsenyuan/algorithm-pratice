package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 3, 1}
	expect := 3
	runSample(t, a, expect)
}
func TestSample2(t *testing.T) {
	a := []int{5, 2, 3, 1, 5, 2}
	expect := 9
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{7, 2, 4, 9, 1, 4, 6, 3, 7, 4, 2, 3}
	expect := 35
	runSample(t, a, expect)
}
