package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, p int, A []int, expect []int) {
	res := solve(n, p, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, p := 3, 2
	A := []int{3, 4, 5}
	expect := []int{3}
	runSample(t, n, p, A, expect)
}

func TestSample2(t *testing.T) {
	n, p := 4, 3
	A := []int{2, 3, 5, 6}
	expect := []int{3, 4}
	runSample(t, n, p, A, expect)
}

func TestSample3(t *testing.T) {
	n, p := 4, 3
	A := []int{9, 1, 1, 1}
	runSample(t, n, p, A, nil)
}
