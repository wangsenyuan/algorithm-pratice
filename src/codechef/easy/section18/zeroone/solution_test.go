package main

import "testing"

func runSample(t *testing.T, C []int, expect int) {
	res := int(solve(C))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := []int{2, 3, 4, 3}
	expect := 30
	runSample(t, C, expect)
}
