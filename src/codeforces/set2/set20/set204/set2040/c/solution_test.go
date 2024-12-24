package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	expect := []int{1, 3, 2}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 3
	expect := []int{2, 3, 1}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 4, 11
	runSample(t, n, k, nil)
}

func TestSample4(t *testing.T) {
	n, k := 4, 6
	expect := []int{2, 4, 3, 1}
	runSample(t, n, k, expect)
}

func TestSample5(t *testing.T) {
	n, k := 7, 34
	expect := []int{2, 3, 4, 5, 7, 6, 1}
	runSample(t, n, k, expect)
}
