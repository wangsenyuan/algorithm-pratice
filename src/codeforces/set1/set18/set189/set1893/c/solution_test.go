package main

import "testing"

func runSample(t *testing.T, L []int, R []int, a [][]int, c [][]int, expect int) {
	res := solve(L, R, a, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	L := []int{5, 1, 4}
	R := []int{6, 3, 4}
	a := [][]int{
		{10, 11, 12},
		{12},
		{12, 13},
	}
	c := [][]int{
		{3, 3, 1},
		{4},
		{1, 5},
	}
	expect := 1
	runSample(t, L, R, a, c, expect)
}

func TestSample2(t *testing.T) {
	L := []int{1000}
	R := []int{1006}
	a := [][]int{
		{1000, 1001, 1002, 1003, 1004, 1005, 1006},
	}
	c := [][]int{
		{147, 145, 143, 143, 143, 143, 142},
	}
	expect := 139
	runSample(t, L, R, a, c, expect)
}

func TestSample3(t *testing.T) {
	L := []int{1, 1}
	R := []int{1, 1}
	a := [][]int{
		{1},
		{2},
	}
	c := [][]int{
		{1},
		{1},
	}
	expect := 1
	runSample(t, L, R, a, c, expect)
}
