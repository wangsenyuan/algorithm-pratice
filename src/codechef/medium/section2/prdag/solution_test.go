package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, mat []string, expect [][]int) {
	ms := make([][]byte, n)
	for i := 0; i < n; i++ {
		ms[i] = []byte(mat[i])
	}
	res := solve(n, ms)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, mat, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	mat := []string{
		"010",
		"000",
		"000",
	}

	expect := [][]int{{1, 3}, {2, 3}}

	runSample(t, n, mat, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	mat := []string{
		"000",
		"000",
		"100",
	}

	expect := [][]int{{1, 2}, {3, 2}}

	runSample(t, n, mat, expect)
}
