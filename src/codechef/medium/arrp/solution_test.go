package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 4, 2, 3, 5}
	expect := []int{1, 0, 1, 0, 0}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 1, 1, 1}
	expect := []int{1, 1, 0, 1}
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 1, 2, 2}
	expect := []int{1, 0, 1, 0}
	runSample(t, n, A, expect)
}
