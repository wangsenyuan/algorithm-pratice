package main

import "testing"

func runSample(t *testing.T, n int, s string, ops [][]int64, ques []int64, expect string) {
	res := string(solve(n, s, ops, ques))

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	s := "mark"
	ops := [][]int64{
		{1, 4},
		{5, 7},
		{3, 8},
	}
	ques := []int64{1, 10, 12}
	expect := "mar"
	runSample(t, n, s, ops, ques, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	s := "creamii"
	ops := [][]int64{
		{2, 3},
		{3, 4},
		{2, 9},
	}
	ques := []int64{9, 11, 12}
	expect := "ear"
	runSample(t, n, s, ops, ques, expect)
}
