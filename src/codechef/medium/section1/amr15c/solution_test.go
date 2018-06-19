package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, expect []int) {
	res := solve(n, k)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 2, 2
	runSample(t, n, k, nil)
}

func TestSample2(t *testing.T) {
	n, k := 3, 0
	runSample(t, n, k, []int{1, 2, 3})
}

func TestSample3(t *testing.T) {
	n, k := 3, 1
	runSample(t, n, k, []int{2, 3, 1})
}

func TestSample4(t *testing.T) {
	n, k := 6, 3
	expect := []int{4, 5, 6, 1, 2, 3}
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 7, 3
	expect := []int{4, 5, 6, 7, 1, 2, 3}
	runSample(t, n, k, expect)
}

func TestSample6(t *testing.T) {
	n, k := 8, 3
	expect := []int{4, 5, 6, 7, 8, 1, 2, 3}
	runSample(t, n, k, expect)
}

func TestSample7(t *testing.T) {
	n, k := 9, 3
	expect := []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	runSample(t, n, k, expect)
}

func TestSample8(t *testing.T) {
	n, k := 10, 3
	expect := []int{4, 5, 6, 1, 8, 9, 10, 2, 3, 7}
	runSample(t, n, k, expect)
}

func TestSample9(t *testing.T) {
	n, k := 11, 3
	expect := []int{4, 5, 6, 1, 2, 9, 10, 11, 3, 7, 8}
	runSample(t, n, k, expect)
}

func TestSample10(t *testing.T) {
	n, k := 12, 3
	expect := []int{4, 5, 6, 1, 2, 3, 10, 11, 12, 7, 8, 9}
	runSample(t, n, k, expect)
}
