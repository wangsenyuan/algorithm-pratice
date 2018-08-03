package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, G [][]int, expect []int) {
	res := solve(n, G)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	G := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expect := []int{3, 2, 1}
	runSample(t, n, G, expect)
}
