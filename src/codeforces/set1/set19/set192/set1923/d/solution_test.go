package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 4, 2}
	expect := []int{2, 1, 2, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := []int{1, 1, -1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 2, 3, 1, 1}
	expect := []int{2, 1, -1, 1, 2}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, 2, 3, 6, 1, 1, 8}
	expect := []int{2, 1, 1, 3, 1, 1, 4}
	runSample(t, a, expect)
}
