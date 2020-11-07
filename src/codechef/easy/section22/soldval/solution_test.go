package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{6, 5, 5, 5, 2}
	expect := []int{6, 5, 4, 3, 2}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	expect := []int{1, 2, 3, 4, 5}
	runSample(t, n, A, expect)
}
