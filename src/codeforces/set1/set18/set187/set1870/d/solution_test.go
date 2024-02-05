package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, k int, expect []int) {
	res := solve(a, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	k := 5
	expect := []int{5, 0, 0}
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4}
	k := 7
	expect := []int{2, 1}
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 1}
	k := 2
	expect := []int{2, 2, 2}
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{10, 6, 4, 6, 3, 4}
	k := 7
	expect := []int{2, 2, 2, 2, 2, 1}
	runSample(t, a, k, expect)
}
