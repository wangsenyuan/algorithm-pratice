package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N, M, K int, A []int, L []int, R []int, expect []int) {
	res := solve(N, M, K, A, L, R)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v %v, expect %v, but got %v", N, M, K, A, L, R, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, K := 7, 7, 5
	A := []int{4, 5, 6, 6, 5, 7, 4}
	L := []int{6, 5, 3, 3, 1}
	R := []int{6, 6, 5, 7, 7}
	expect := []int{0, 0, 1, 1, 6}
	runSample(t, N, M, K, A, L, R, expect)
}
