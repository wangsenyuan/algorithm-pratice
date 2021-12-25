package main

import "testing"

func runSample(t *testing.T, n int, row []int, expect int) {
	res := solve(n, row)
	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, row, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	row := []int{0, 1, 1, 1, 0, 1, 1, 0, 1, 1}
	expect := 7
	runSample(t, n, row, expect)
}
