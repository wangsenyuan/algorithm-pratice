package main

import "testing"

func runSample(t *testing.T, people [][]int, S int, expect int) {
	res := solve(people, S)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := 12
	people := [][]int{
		{3, 5, 7},
		{4, 6, 7},
		{5, 9, 5},
	}
	expect := 84
	runSample(t, people, S, expect)
}

func TestSample2(t *testing.T) {
	S := 10
	people := [][]int{
		{7, 4, 7},
		{5, 8, 8},
		{12, 5, 8},
		{6, 11, 6},
		{3, 3, 7},
		{5, 9, 6},
	}
	expect := 314
	runSample(t, people, S, expect)
}

func TestSample3(t *testing.T) {
	S := 100
	people := [][]int{
		{97065, 97644, 98402},
	}
	expect := 9551390130
	runSample(t, people, S, expect)
}
