package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s []int, expect []int) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := []int{0, 0, 0}
	expect := []int{3, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := []int{0, 1}
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := []int{0, 1, 1, 1, 10}
	expect := []int{1, 4, 3, 2, 5}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := []int{0, 0, 3, 3, 13, 3, 6, 34, 47, 12, 20, 6, 6, 21, 55}
	expect := []int{2, 1, 15, 10, 12, 3, 6, 13, 14, 8, 9, 5, 4, 7, 11}
	runSample(t, s, expect)
}
