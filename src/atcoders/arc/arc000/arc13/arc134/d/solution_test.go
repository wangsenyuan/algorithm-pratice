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
	a := []int{2, 1, 3, 1, 2, 2}
	expect := []int{1, 2}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{38, 38, 80, 62, 62, 67, 38, 78, 74, 52, 53, 77, 59, 83, 74, 63, 80, 61, 68, 55}
	expect := []int{38, 38, 38, 52, 53, 77, 80, 55}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{52, 73, 49, 63, 55, 74, 35, 68, 22, 22, 74, 50, 71, 60, 52, 62, 65, 54, 70, 59, 65, 54, 60, 52}
	expect := []int{22, 22, 50, 65, 54, 52}
	runSample(t, a, expect)
}
