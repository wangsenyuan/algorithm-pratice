package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 1, 2}
	expect := []int{1, 0, -1, -1, 2}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 6, 6, 4, 2, 1}
	expect := []int{1, 0, 1, -1, -1, 2}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 2, 2, 3, 2, 1, 4}
	expect := []int{0, 1, 2, -1, -1, -1, 2}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 2, 4, 2, 4, 1, 1, 4, 4}
	expect := []int{0, 1, 1, 3, -1, -1, -1, -1, 6}
	runSample(t, A, expect)
}
