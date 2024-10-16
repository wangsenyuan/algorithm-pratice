package main

import "testing"

func runSample(t *testing.T, s int, employees [][]int, expect int) {
	res := solve(s, employees)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := 26
	employees := [][]int{
		{10, 12},
		{1, 4},
		{10, 11},
	}
	expect := 11
	runSample(t, s, employees, expect)
}

func TestSample2(t *testing.T) {
	s := 1337
	employees := [][]int{
		{1, 1000000000},
	}
	expect := 1337
	runSample(t, s, employees, expect)
}

func TestSample3(t *testing.T) {
	s := 26
	employees := [][]int{
		{4, 4},
		{2, 4},
		{6, 8},
		{5, 6},
		{2, 7},
	}
	expect := 6
	runSample(t, s, employees, expect)
}
