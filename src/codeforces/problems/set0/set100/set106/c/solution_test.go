package main

import "testing"

func runSample(t *testing.T, n int, c0, d0 int, stuffs [][]int, expect int) {
	res := solve(n, c0, d0, stuffs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	c0, d0 := 2, 1
	stuffs := [][]int{
		{7, 3, 2, 100},
		{12, 3, 1, 10},
	}
	expect := 241
	runSample(t, n, c0, d0, stuffs, expect)
}

func TestSample2(t *testing.T) {
	n := 100
	c0, d0 := 25, 50
	stuffs := [][]int{
		{15, 5, 20, 10},
	}
	expect := 200
	runSample(t, n, c0, d0, stuffs, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	c0, d0 := 11, 5
	stuffs := [][]int{
		{100, 1, 3, 10},
		{100, 1, 2, 4},
	}
	expect := 30
	runSample(t, n, c0, d0, stuffs, expect)
}
