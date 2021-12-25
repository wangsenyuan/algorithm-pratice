package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, L []int, R []int, m int, P []int, expect []int) {
	res := solve(n, L, R, m, P)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v %d %v, expect %v, but got %v", n, L, R, m, P, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	L := []int{5, 9, 2, 20}
	R := []int{7, 10, 3, 30}
	P := []int{5, 6, 7, 35, 1}
	expect := []int{0, 0, 2, -1, 1}
	runSample(t, n, L, R, m, P, expect)
}
