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
	runSample(t, 4, 1, []int{2, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, []int{0, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 2, []int{0, 1})
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 2, []int{1, 3})
}

func TestSample5(t *testing.T) {
	runSample(t, 9, 2, []int{6, 3})
}

func TestSample6(t *testing.T) {
	runSample(t, 1000000000, 1, []int{249999999500000000, 1})
}
