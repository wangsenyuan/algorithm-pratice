package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, B []int, m int, L []int, R []int, X []int, Y []int, expect []int) {
	res := solve(n, A, B, m, L, R, X, Y)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %d %v %v %v %v, expect %v, but got %v", n, A, B, m, L, R, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{5, 4, 4, 2, 1}
	B := []int{7, 7, 4, 7, 7}
	m := 3
	L := []int{1, 2, 3}
	R := []int{2, 5, 4}
	X := []int{2, 1, 2}
	Y := []int{0, 1, 2}
	expect := []int{1, 2, 0, 3, -1}
	runSample(t, n, A, B, m, L, R, X, Y, expect)
}
