package main

import "testing"

func runSample(t *testing.T, n int, E [][]int, X []int, Y []int, S []string, expect string) {
	res := solve(n, E, X, Y, S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	E := [][]int{
		{2, 3},
		{2, 1},
	}
	X := []int{2, 2}
	Y := []int{1, 3}
	S := []string{"ab", "bc"}
	expect := "abc"
	runSample(t, n, E, X, Y, S, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	E := [][]int{
		{2, 3},
		{2, 1},
	}
	X := []int{2, 2}
	Y := []int{1, 3}
	S := []string{"ab", "cd"}
	expect := ""
	runSample(t, n, E, X, Y, S, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	E := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{1, 8},
		{1, 9},
		{1, 10},
	}
	X := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 10}
	Y := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2}
	S := []string{"ab", "ab", "ab", "ab", "ab", "ab", "ab", "ab", "ab", "aba"}
	expect := "baaaaaaaaa"
	runSample(t, n, E, X, Y, S, expect)
}
