package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, heights []int, Q [][]int, expect string) {
	solver := NewSolver(n, E, heights)

	var buf bytes.Buffer

	for _, q := range Q {
		res := solver.Query(q[0], q[1])
		if res {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}

	if buf.String() != expect {
		t.Errorf("Sample expect %s, but got %s", expect, buf.String())
	}
}

func TestSample1(t *testing.T) {
	n := 5
	E := [][]int{
		{1, 5},
		{1, 2},
		{4, 2},
		{2, 3},
	}
	heights := []int{4, 6, 2, 1, 5}
	Q := [][]int{
		{1, 5},
		{3, 4},
		{3, 2},
		{5, 4},
	}
	expect := "1110"
	runSample(t, n, E, heights, Q, expect)
}
