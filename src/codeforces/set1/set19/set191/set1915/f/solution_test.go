package main

import "testing"

func runSample(t *testing.T, people [][]int, expect int) {
	res := solve(people)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	people := [][]int{
		{2, 3},
		{1, 4},
	}
	expect := 1
	runSample(t, people, expect)
}

func TestSample2(t *testing.T) {
	people := [][]int{
		{2, 6},
		{3, 9},
		{4, 5},
		{1, 8},
		{7, 10},
		{-2, 100},
	}
	expect := 9
	runSample(t, people, expect)
}

func TestSample3(t *testing.T) {
	people := [][]int{
		{-10, 10},
		{-5, 5},
		{-12, 12},
		{-13, 13},
	}
	expect := 6
	runSample(t, people, expect)
}
