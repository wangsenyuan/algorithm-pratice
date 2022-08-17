package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, k int, G []string, Q [][]int, expect []int) {
	res := solve(n, m, k, G, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, k := 5, 5, 4
	G := []string{
		"H.H..",
		"..H..",
		"H...H",
		"HHHH.",
		"HH..H",
	}
	Q := [][]int{
		{1, 1},
		{2, 3},
	}
	expect := []int{0, 4}
	runSample(t, n, m, k, G, Q, expect)
}
