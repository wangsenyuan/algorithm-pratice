package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, b []int) {
	res := solve(b)
	n := len(b)
	x := make([]int, n)
	for i := 1; i <= n; i++ {
		x[i-1] = i / res[i-1]
	}

	if !reflect.DeepEqual(x, b) {
		t.Fatalf("Sample result %v, not correct, it leads to %v", res, x)
	}
}

func TestSample1(t *testing.T) {
	b := []int{0, 2, 0, 1}
	runSample(t, b)
}

func TestSample2(t *testing.T) {
	b := []int{1, 1}
	runSample(t, b)
}

func TestSample3(t *testing.T) {
	b := []int{0, 0, 1, 4, 1}
	runSample(t, b)
}

func TestSample4(t *testing.T) {
	b := []int{0, 1, 3}
	runSample(t, b)
}
