package main

import "testing"

func runSample(t *testing.T, v int, boats [][]int, expect int) {
	res, _ := solve(v, boats)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	v := 2
	boats := [][]int{
		{1, 2},
		{2, 7},
		{1, 3},
	}
	expect := 7
	runSample(t, v, boats, expect)
}
