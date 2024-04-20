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
		{11, 6},
		{10, 4},
		{01, 3},
		{00, 3},
		{00, 7},
		{00, 9},
	}
	expect := 22
	runSample(t, people, expect)
}

func TestSample2(t *testing.T) {
	people := [][]int{
		{11, 1},
		{01, 1},
		{00, 100},
		{10, 1},
		{01, 1},
	}
	expect := 103
	runSample(t, people, expect)
}

func TestSample3(t *testing.T) {
	people := [][]int{
		{11, 19},
		{10, 22},
		{00, 18},
		{00, 29},
		{11, 29},
		{10, 28},
	}
	expect := 105
	runSample(t, people, expect)
}

func TestSample4(t *testing.T) {
	people := [][]int{
		{00, 5000},
		{00, 5000},
		{00, 5000},
	}
	expect := 0
	runSample(t, people, expect)
}

func TestSample5(t *testing.T) {
	people := [][]int{
		{01, 13},
		{11, 3},
	}
	expect := 16
	runSample(t, people, expect)
}
