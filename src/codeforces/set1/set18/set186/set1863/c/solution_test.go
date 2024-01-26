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
	a := []int{1}
	k := 2
	expect := []int{1}
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 3}
	k := 1
	expect := []int{2, 0, 1}
	runSample(t, a, k, expect)
}
func TestSample3(t *testing.T) {
	a := []int{0, 2}
	k := 2
	expect := []int{2, 1}
	runSample(t, a, k, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 3, 0, 4, 2, 1, 6, 9, 10, 8}
	k := 100
	expect := []int{7, 5, 3, 0, 4, 2, 1, 6, 9, 10}
	runSample(t, a, k, expect)
}
