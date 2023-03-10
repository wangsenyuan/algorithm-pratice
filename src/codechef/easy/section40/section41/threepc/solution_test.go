package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, B []int, C []int, expect []int) {
	res := solve(A, B, C)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 5, 8, 1, 3, 7}
	B := []int{0, 8, 1, 1, 2, 0}
	C := []int{8, 8, 3, 9, 3, 3}
	expect := []int{4, 1, 6}
	runSample(t, A, B, C, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 3}
	B := []int{2, 2}
	C := []int{1, 1}
	expect := []int{2, 0, 0}
	runSample(t, A, B, C, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 1}
	B := []int{1, 1, 1}
	C := []int{1, 1, 1}
	expect := []int{3, 3, 3}
	runSample(t, A, B, C, expect)
}
