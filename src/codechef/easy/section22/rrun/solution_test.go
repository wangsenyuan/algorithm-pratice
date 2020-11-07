package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{4, 2, 3, 1}
	expect := []int{7, 7, 7, 6}
	runSample(t, n, A, expect)
}
