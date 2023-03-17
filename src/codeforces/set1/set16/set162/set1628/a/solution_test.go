package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 0, 2, 0, 3}
	expect := []int{4}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 3, 4, 0, 1, 2, 0}
	expect := []int{5, 1}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1}
	expect := []int{0}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 1, 2, 3, 4}
	expect := []int{5}
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{0, 1, 1, 0}
	expect := []int{2, 2}
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{0, 0, 2, 1, 1, 1, 0, 0, 1, 1}
	expect := []int{3, 2, 2, 0}
	runSample(t, A, expect)
}
