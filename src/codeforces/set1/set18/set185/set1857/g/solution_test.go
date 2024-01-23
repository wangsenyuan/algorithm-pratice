package main

import "testing"

func runSample(t *testing.T, n int, s int, edges [][]int, expect int) {
	res := solve(n, s, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, s := 2, 5
	edges := [][]int{
		{1, 2, 4},
	}
	expect := 1
	runSample(t, n, s, edges, expect)
}

func TestSample2(t *testing.T) {
	n, s := 4, 5
	edges := [][]int{
		{1, 2, 2},
		{2, 3, 4},
		{3, 4, 3},
	}
	expect := 8
	runSample(t, n, s, edges, expect)
}

func TestSample3(t *testing.T) {
	n, s := 5, 6
	edges := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{3, 4, 6},
		{3, 5, 1},
	}
	expect := 80
	runSample(t, n, s, edges, expect)
}

func TestSample4(t *testing.T) {
	n, s := 10, 200
	edges := [][]int{
		{1, 2, 3},
		{2, 3, 33},
		{3, 4, 200},
		{1, 5, 132},
		{5, 6, 1},
		{5, 7, 29},
		{7, 8, 187},
		{7, 9, 20},
		{7, 10, 4},
	}
	expect := 650867886
	runSample(t, n, s, edges, expect)
}
