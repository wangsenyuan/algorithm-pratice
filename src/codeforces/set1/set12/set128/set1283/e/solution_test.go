package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x []int, expect []int) {
	res := solve(x)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{1, 2, 4, 4}
	expect := []int{2, 4}
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1, 1, 8, 8, 8, 4, 4, 4, 4}
	expect := []int{3, 8}
	runSample(t, x, expect)
}

func TestSample3(t *testing.T) {
	x := []int{4, 3, 7, 1, 4, 3, 3}
	expect := []int{3, 6}
	runSample(t, x, expect)
}

func TestSample4(t *testing.T) {
	x := []int{1, 2}
	expect := []int{1, 2}
	runSample(t, x, expect)
}
