package main

import "testing"

func runSample(t *testing.T, n int, x int, y int, choosen []int, expect int) {
	res := solve(n, x, y, choosen)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 7, 3, 1
	choosen := []int{6, 4, 3}
	expect := 5
	runSample(t, n, x, y, choosen, expect)
}
