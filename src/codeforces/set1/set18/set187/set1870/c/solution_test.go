package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, k int, a []int, expect []int) {
	res := solve(k, a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := []int{1, 1}
	expect := []int{4}
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{1, 2}
	expect := []int{4, 2}
	runSample(t, k, a, expect)
}

func TestSample3(t *testing.T) {
	k := 5
	a := []int{3, 2, 4}
	expect := []int{0, 6, 6, 2, 0}
	runSample(t, k, a, expect)
}

func TestSample4(t *testing.T) {
	k := 3
	a := []int{1, 2, 3, 2, 1}
	expect := []int{10, 6, 2}
	runSample(t, k, a, expect)
}
