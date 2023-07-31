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
	A := []int{1, 9, 17, 6, 14, 3}
	expect := []int{19, 8}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 2, 2}
	expect := []int{-1}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{7, 3, 4}
	expect := []int{-1}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 2, 4}
	expect := []int{-1}
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{0, 1000000000, 0, 1000000000, 0}
	expect := []int{2000000000, 1000000000}
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{1, 1}
	expect := []int{0}
	runSample(t, A, expect)
}
