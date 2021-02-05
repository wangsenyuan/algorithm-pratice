package main

import "testing"

func runSample(tc *testing.T, n, m int, s, t int, edges []Edge, expect int64) {
	res := solve(n, m, s, t, edges)

	if res != expect {
		tc.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(tc *testing.T) {
	n, m, s, t := 6, 6, 1, 6
	edges := []Edge{
		{1, 2, 2, 'a'},
		{1, 2, 3, 'b'},
		{2, 3, 1, 'b'},
		{3, 4, 1, 'a'},
		{4, 5, 1, 'a'},
		{5, 6, 1, 'a'},
	}
	var expect int64 = 10
	runSample(tc, n, m, s, t, edges, expect)
}

func TestSample2(tc *testing.T) {
	n, m, s, t := 3, 2, 1, 3
	edges := []Edge{
		{1, 2, 1, 'a'},
		{2, 3, 2, 'b'},
	}
	var expect int64 = -1
	runSample(tc, n, m, s, t, edges, expect)
}

func TestSample3(tc *testing.T) {
	n, m, s, t := 3, 2, 1, 2
	edges := []Edge{
		{1, 2, 1, 'a'},
		{2, 3, 2, 'b'},
	}
	var expect int64 = 1
	runSample(tc, n, m, s, t, edges, expect)
}
