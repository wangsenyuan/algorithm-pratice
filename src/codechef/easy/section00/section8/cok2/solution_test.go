package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, Q, K, L, R int, C []int, P []int, expect []int) {
	res := solve(N, Q, K, L, R, C, P)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d %d %v %v, expect %v, but got %v", N, Q, K, L, R, C, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, Q, K, L, R := 3, 2, 5, 4, 6
	C := []int{1, 2, 8}
	P := []int{6, 8, 10}
	expect := []int{-1, 8}
	runSample(t, N, Q, K, L, R, C, P, expect)
}

func TestSample2(t *testing.T) {
	N, Q, K, L, R := 3, 5, 10, 20, 30
	C := []int{21, 22, 23}
	P := []int{20, 22, 23}
	expect := []int{20, 22, 23, -1, -1}
	runSample(t, N, Q, K, L, R, C, P, expect)
}
