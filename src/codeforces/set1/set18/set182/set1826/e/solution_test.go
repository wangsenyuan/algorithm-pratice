package main

import "testing"

func runSample(t *testing.T, n int, m int, rates [][]int, p []int, expect int) {

	var it int

	f := func() []int {
		res := rates[it]
		it++
		return res
	}

	res := solve(n, m, p, f)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 5
	p := []int{10, 10, 10, 10, 10}
	rates := [][]int{
		{1, 2, 3, 4, 5},
		{1, 5, 2, 3, 4},
		{2, 3, 4, 5, 1},
	}
	expect := 30
	runSample(t, n, m, rates, p, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 5
	p := []int{10, 10, 10, 10, 50}
	rates := [][]int{
		{1, 2, 3, 4, 5},
		{1, 5, 2, 3, 4},
		{2, 3, 4, 5, 1},
	}
	expect := 50
	runSample(t, n, m, rates, p, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 5
	p := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	rates := [][]int{
		{5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
	}
	expect := 5000000000
	runSample(t, n, m, rates, p, expect)
}

func TestSample4(t *testing.T) {
	n, m := 1, 3
	p := []int{1, 2, 3}
	rates := [][]int{
		{3, 3, 3},
	}
	expect := 3
	runSample(t, n, m, rates, p, expect)
}
