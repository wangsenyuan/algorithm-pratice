package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, k int, expect []int) {
	res := solve(A, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 6, 3, 2, 5, 1}
	k := 2
	expect := []int{1, 2, 3, 4, 5, 6}
	runSample(t, A, k, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 6, 3, 2, 5, 1}
	k := 3
	expect := []int{1, 4, 3, 2, 5, 6}
	runSample(t, A, k, expect)
}

func TestSample3(t *testing.T) {
	A := []int{4, 6, 3, 2, 5, 1}
	k := 4
	expect := []int{1, 6, 3, 2, 5, 4}
	runSample(t, A, k, expect)
}

func TestSample5(t *testing.T) {
	A := []int{4, 6, 3, 2, 5, 1}
	k := 5
	expect := []int{4, 6, 3, 2, 5, 1}
	runSample(t, A, k, expect)
}

func TestSample6(t *testing.T) {
	A := []int{4, 6, 3, 2, 5, 1}
	k := 6
	expect := []int{4, 6, 3, 2, 5, 1}
	runSample(t, A, k, expect)
}