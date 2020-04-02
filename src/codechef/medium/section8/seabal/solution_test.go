package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, M int, A []int, L []int, R []int, K int, X []int, expect []int) {
	res := solve(N, M, A, L, R, K, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M := 5, 3
	A := []int{1, 1, 1, 1, 1}
	L := []int{5, 2, 1}
	R := []int{5, 2, 3}
	K := 5
	X := []int{4, 2, 0, 2, 3}
	expect := []int{0, 1, 1, 2, 3}
	runSample(t, N, M, A, L, R, K, X, expect)
}
