package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{2, 2, 3}
	expect := []int{0, 6, 1}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{1, 2}
	expect := []int{2, 1}
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{1, 2, 2}
	expect := []int{3, 4, 0}
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	A := []int{1, 1, 1}
	expect := []int{7, 0, 0}
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 3
	A := []int{1, 2, 3}
	// 1, 1 2, 1 2, 1, 2, 3
	// 2  2, 3
	// 3
	expect := []int{4, 2, 1}
	runSample(t, n, A, expect)
}

func TestSample6(t *testing.T) {
	n := 4
	A := []int{1, 2, 2, 3}
	// [1], [1, 2], [1, 2], [1, 3], [1, 2, 3], [1, 2, 3]
	// [2], [2], [2, 2], [2, 3], [2, 3], [1, 2, 2], [2, 2, 3], [1, 2, 2, 3]
	// 3
	expect := []int{6, 8, 1, 0}
	runSample(t, n, A, expect)
}
