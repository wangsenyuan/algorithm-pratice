package main

import "testing"

func runSample(t *testing.T, a, b int, expect [][]int) {
	res := solve(a, b)

	s1 := check(a, b, expect)
	s2 := check(a, b, res)

	if s1 != s2 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func check(a, b int, arr [][]int) int {
	var s1 int
	for _, x := range arr[0] {
		if x <= 0 || x > a {
			return -1
		}
		s1 += x
	}
	if s1 > a {
		return -1
	}

	var s2 int

	for _, x := range arr[1] {
		if x <= 0 || x > b {
			return -1
		}
		s2 += x
	}

	if s2 > b {
		return -1
	}

	return s1 + s2
}

func TestSample1(t *testing.T) {
	a, b := 9, 12
	expect := [][]int{
		{3, 6},
		{1, 2, 4, 5},
	}
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 3, 3
	expect := [][]int{
		{3},
		{1, 2},
	}
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 0, 0
	expect := [][]int{
		{},
		{},
	}
	runSample(t, a, b, expect)
}
