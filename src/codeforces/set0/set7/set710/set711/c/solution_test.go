package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, c []int, p [][]int, expect int) {
	res := solve(n, m, k, c, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 3, 2, 2
	c := []int{0, 0, 0}
	p := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	expect := 10
	runSample(t, n, m, k, c, p, expect)
}

func TestSample2(t *testing.T) {
	n, m, k := 3, 2, 2
	c := []int{2, 1, 2}
	p := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
	}
	expect := -1
	runSample(t, n, m, k, c, p, expect)
}

func TestSample3(t *testing.T) {
	n, m, k := 3, 2, 2
	c := []int{2, 0, 0}
	p := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
	}
	expect := 5
	runSample(t, n, m, k, c, p, expect)
}

func TestSample4(t *testing.T) {
	n, m, k := 3, 2, 3
	c := []int{2, 1, 2}
	p := [][]int{
		{1, 3},
		{2, 4},
		{3, 5},
	}
	expect := 0
	runSample(t, n, m, k, c, p, expect)
}

func TestSample5(t *testing.T) {
	n, m, k := 4, 2, 1
	c := []int{0, 0, 0, 2}
	p := [][]int{
		{10, 30000},
		{20000, 1000000000},
		{1000000000, 50000},
		{55, 55},
	}
	expect := 1000080000
	runSample(t, n, m, k, c, p, expect)
}

func TestSample6(t *testing.T) {
	n, m, k := 2, 1, 2
	c := []int{0, 1}
	p := [][]int{
		{10},
		{10},
	}
	expect := -1
	runSample(t, n, m, k, c, p, expect)
}
