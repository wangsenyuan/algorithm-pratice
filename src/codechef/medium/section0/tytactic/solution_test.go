package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, skills []int, edges [][]int, queries [][]byte, expect []int) {
	res := solve(n, m, skills, edges, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", n, m, skills, edges, queries, expect)
	}
}

func TestSample1(t *testing.T) {
	n, m := 5, 8
	skills := []int{7, 2, 0, 5, 8}
	edges := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{1, 5},
	}

	queries := [][]byte{
		[]byte("Q 1"),
		[]byte("Q 2"),
		[]byte("U 2 4"),
		[]byte("Q 1"),
		[]byte("Q 2"),
		[]byte("U 5 3"),
		[]byte("Q 1"),
		[]byte("Q 2"),
	}

	expect := []int{22, 7, 24, 9, 19, 9}

	runSample(t, n, m, skills, edges, queries, expect)
}
