package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B []int, expect []int) {
	res := solve(A, B)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 3, 5}
	B := []int{6}
	expect := []int{6}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 4, 7}
	B := []int{2, 3, 7}
	expect := []int{2, 7, 3}
	runSample(t, A, B, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 7, 4}
	B := []int{2, 7, 3}
	runSample(t, A, B, nil)
}
