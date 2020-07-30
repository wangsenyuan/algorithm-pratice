package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)
	if expect != (len(res) > 0) {
		t.Errorf("Sample %d %v, expect %t, but got %v", n, A, expect, res)
		return
	}

	if expect {
		a := calMex(n, res)

		if !reflect.DeepEqual(a, A) {
			t.Errorf("Sample %d %v, result %v, not correct with %v", n, A, res, a)
		}
	}
}

func calMex(n int, A []int) []int {
	res := make([]int, n)
	mem := make(map[int]bool)
	var mex int

	for i := 0; i < n; i++ {
		mem[A[i]] = true
		for mem[mex] {
			mex++
		}
		res[i] = mex
	}

	return res
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	runSample(t, n, A, true)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{0, 0, 0, 2}
	runSample(t, n, A, true)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{1, 1, 3}
	runSample(t, n, A, true)
}
