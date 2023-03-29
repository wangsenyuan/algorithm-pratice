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
	A := []int{1}
	expect := []int{1}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 4, 2}
	expect := []int{1, 1, 2}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 12, 4, 8, 18, 3, 6}
	expect := []int{0, 1, 1, 1, 2, 2, 2}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{3, 1, 3, 2, 3, 1, 1, 3, 2, 1, 1, 3, 3, 3, 3, 3, 2, 1, 2, 1, 3, 2, 2, 1, 3, 1, 3, 1, 2, 3}
	expect := []int{0, 1, 1, 1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 8, 8, 8, 9, 10, 10, 11, 11, 12, 12, 12}
	runSample(t, A, expect)
}
