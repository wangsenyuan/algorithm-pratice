package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, A []int, expect []int) {
	res := solve(n, A)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v, expect %v, but got %v", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{4, 3, 1, 2}
	expect := []int{4, 3, -1, 2}
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{1, 2, 2, 1, 3, 1}
	expect := []int{-1, 2, 2, -1, 3, -1}
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{10, 1, 2, 10, 5}
	expect := []int{10, -1, 2, 10, -5}
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []int{1, 2, 1, 2}
	expect := []int{1, 2, -1, 2}
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	A := []int{1, 3, 1, 2}
	expect := []int{-1, 3, -1, 2}
	runSample(t, n, A, expect)
}

func TestSample6(t *testing.T) {
	n := 5
	A := []int{1, 1, 2, 1, 1}
	expect := []int{1, 1, 2, 1, 1}
	runSample(t, n, A, expect)
}

func TestSample7(t *testing.T) {
	n := 3
	A := []int{5, 10, 6}
	expect := []int{5, 10, -6}
	runSample(t, n, A, expect)
}

func TestSample8(t *testing.T) {
	n := 11
	A := []int{1, 5, 1, 1, 4, 5, 3, 5, 3, 5, 4}
	expect := []int{-1, 5, 1, 1, 4, 5, -3, 5, 3, 5, -4}
	runSample(t, n, A, expect)
}
