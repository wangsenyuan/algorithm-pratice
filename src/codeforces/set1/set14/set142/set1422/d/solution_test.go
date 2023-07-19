package main

import "testing"

func runSample(t *testing.T, m int, s []int, f []int, locations [][]int, expect int) {
	res := solve(m, s, f, locations)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	s := []int{1, 1}
	f := []int{5, 5}
	special := [][]int{
		{1, 2},
		{4, 1},
		{3, 3},
	}
	expect := 5
	runSample(t, m, s, f, special, expect)
}

func TestSample2(t *testing.T) {
	m := 5
	s := []int{67, 59}
	f := []int{41, 2}
	special := [][]int{
		{39, 56},
		{7, 2},
		{15, 3},
		{74, 18},
		{22, 7},
	}
	expect := 42
	runSample(t, m, s, f, special, expect)
}
