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
	a := []int{1, 1, 1, 1, 1}
	expect := []int{1, 2, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	expect := []int{0, 2}
	runSample(t, a, expect)
}
func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 2, 2}
	expect := []int{0, 1, 3}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{2, 2, 4, 4, 3, 1, 1, 2, 3, 2}
	expect := []int{0, 1, 2, 3, 5}
	runSample(t, a, expect)
}
