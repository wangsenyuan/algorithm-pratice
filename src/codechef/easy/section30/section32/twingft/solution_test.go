package main

import "testing"

func runSample(t *testing.T, k int, A []int, expect int) {
	res := int(solve(k, A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	A := []int{1, 3, 2}
	expect := 3
	runSample(t, k, A, expect)
}

func TestSample2(t *testing.T) {
	k := 1
	A := []int{1, 3, 3}
	expect := 4
	runSample(t, k, A, expect)
}
func TestSample3(t *testing.T) {
	k := 2
	A := []int{5, 1, 3, 2, 4}
	expect := 8
	runSample(t, k, A, expect)
}
