package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 5, 2}
	expect := []int{1, 2, 5}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 6, 2}
	expect := []int{2, 6, 2}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{8, 7, 16, 9}
	expect := []int{7, 8, 9, 16}
	runSample(t, A, expect)
}
