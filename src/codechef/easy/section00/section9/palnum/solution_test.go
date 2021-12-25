package main

import "testing"

func runSample(t *testing.T, N int, M int, k int, S string, ops [][]int, expect string) {
	_, res := solve(N, M, k, S, ops)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, K := 5, 10, 10
	S := "75643"
	ops := [][]int{
		{0, 1, 1},
		{1, 2, 1},
		{2, 3, 1},
		{3, 4, 1},
		{4, 5, 1},
		{5, 6, 1},
		{6, 7, 1},
		{7, 8, 1},
		{8, 9, 1},
		{9, 0, 1},
	}
	expect := "95759"
	runSample(t, N, M, K, S, ops, expect)
}

func TestSample2(t *testing.T) {
	N, M, K := 4, 5, 1000000
	S := "1895"
	ops := [][]int{
		{2, 3, 10},
		{3, 1, 9},
		{1, 2, 5},
		{5, 6, 99},
		{8, 9, 45},
	}
	expect := ""
	runSample(t, N, M, K, S, ops, expect)
}

func TestSample3(t *testing.T) {
	N, M, K := 6, 4, 12
	S := "141746"
	ops := [][]int{
		{1, 5, 3},
		{5, 7, 2},
		{7, 9, 10},
		{6, 1, 5},
	}
	expect := "147741"
	runSample(t, N, M, K, S, ops, expect)
}
