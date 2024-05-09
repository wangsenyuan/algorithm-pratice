package main

import "testing"

func runSample(t *testing.T, n int, a []int, offers [][]int, expect int) {
	res := solve(n, a, offers)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1, 3, 3}
	offers := [][]int{
		{2, 3, 5},
		{2, 1, 1},
	}
	expect := 5
	runSample(t, n, a, offers, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{1, 3, 3, 7}
	//offers := [][]int{
	//	{2, 3, 5},
	//	{2, 1, 1},
	//}
	expect := 16
	runSample(t, n, a, nil, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	a := []int{1, 2, 3, 4, 5}
	offers := [][]int{
		{1, 2, 8},
		{1, 3, 10},
		{1, 4, 7},
		{1, 5, 15},
	}
	expect := 18
	runSample(t, n, a, offers, expect)
}
