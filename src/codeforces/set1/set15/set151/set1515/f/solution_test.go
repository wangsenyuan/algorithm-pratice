package main

import "testing"

func runSample(t *testing.T, n int, x int, edges [][]int, asphalt []int, expect bool) {
	ok, res := solve(n, x, edges, asphalt)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t, %v", expect, ok, res)
	}

}

func TestSample1(t *testing.T) {
	asphalt := []int{0, 0, 0, 4, 0}
	n, x := 5, 1
	edges := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}
	expect := true
	runSample(t, n, x, edges, asphalt, expect)
}

func TestSample2(t *testing.T) {
	asphalt := []int{1, 1}
	n, x := 2, 2
	edges := [][]int{
		{1, 2},
	}
	expect := true
	runSample(t, n, x, edges, asphalt, expect)
}

func TestSample3(t *testing.T) {
	asphalt := []int{0, 1}
	n, x := 2, 2
	edges := [][]int{
		{1, 2},
	}
	expect := false
	runSample(t, n, x, edges, asphalt, expect)
}

func TestSample4(t *testing.T) {
	asphalt := []int{0, 9, 4, 0, 10}
	n, x := 5, 5
	edges := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
		{1, 4},
		{4, 5},
	}
	expect := true
	runSample(t, n, x, edges, asphalt, expect)
}
