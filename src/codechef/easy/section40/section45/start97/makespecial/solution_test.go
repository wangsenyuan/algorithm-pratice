package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, L []int, R []int, expect []int) {
	res := solve(n, L, R)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	L := []int{1, 3, 4}
	R := []int{3, 6, 14}
	expect := []int{1, 2, 5}
	runSample(t, n, L, R, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	L := []int{1, 3, 4}
	R := []int{3, 6, 14}
	expect := []int{0, 0, 2}
	runSample(t, n, L, R, expect)
}
