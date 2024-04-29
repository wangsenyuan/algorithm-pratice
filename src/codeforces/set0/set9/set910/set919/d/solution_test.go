package main

import "testing"

func runSample(t *testing.T, n int, edges [][]int, s string, expect int) {
	res := solve(n, edges, s)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	s := "abaca"
	edges := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 5},
	}
	expect := 3
	runSample(t, n, edges, s, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	s := "xzyabc"
	edges := [][]int{
		{1, 2},
		{3, 1},
		{2, 3},
		{5, 4},
		{4, 3},
		{6, 4},
	}
	expect := -1
	runSample(t, n, edges, s, expect)
}

func TestSample3(t *testing.T) {
	n := 10
	s := "xzyzyzyzqx"
	edges := [][]int{
		{1, 2},
		{2, 4},
		{3, 5},
		{4, 5},
		{2, 6},
		{6, 8},
		{6, 5},
		{2, 10},
		{3, 9},
		{10, 9},
		{4, 6},
		{1, 10},
		{2, 8},
		{3, 7},
	}
	expect := 4
	runSample(t, n, edges, s, expect)
}
