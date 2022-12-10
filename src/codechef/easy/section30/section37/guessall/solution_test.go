package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, expect []int64) {
	k := len(A)
	n := len(B)

	var sum int64

	for i := 0; i < k; i++ {
		sum += int64(A[i])
	}

	ask := func(i int) int64 {
		j := i % (k + 1)
		if j < k {
			return int64(A[j])
		}
		return -sum
	}

	res := solve(k, n, B, ask)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 4, 3}
	B := []int{0, 1, 3}
	expect := []int64{3, 4, -10}
	runSample(t, A, B, expect)
}
