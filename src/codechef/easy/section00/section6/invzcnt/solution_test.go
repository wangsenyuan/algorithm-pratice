package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, Q int, A []int, K []int, expect []int) {
	res := solve(n, Q, A, K)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v %v, expect %v, but got %v", n, Q, A, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, Q := 4, 4
	A := []int{3, 1, 2, 0}
	K := []int{0, 1, 2, 3}
	expect := []int{5, 3, 3, 1}
	runSample(t, n, Q, A, K, expect)
}
