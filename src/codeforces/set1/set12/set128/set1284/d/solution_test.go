package main

import "testing"

func runSample(t *testing.T, lectures [][]int, expect bool) {
	res := solve(lectures)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lectures := [][]int{
		{1, 2, 3, 6},
		{3, 4, 7, 8},
	}
	expect := true
	runSample(t, lectures, expect)
}

func TestSample2(t *testing.T) {
	lectures := [][]int{
		{1, 3, 2, 4},
		{4, 5, 6, 7},
		{3, 4, 5, 5},
	}
	expect := false
	runSample(t, lectures, expect)
}

func TestSample3(t *testing.T) {
	lectures := [][]int{
		{1, 5, 2, 9},
		{2, 4, 5, 8},
		{3, 6, 7, 11},
		{7, 10, 12, 16},
		{8, 11, 13, 17},
		{9, 12, 14, 18},
	}
	expect := true
	runSample(t, lectures, expect)
}
