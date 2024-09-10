package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, x int, y int, expect []int) {
	res := solve(n, x, y)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x, y := 5, 1, 3
	expect := []int{1, 3}
	runSample(t, n, x, y, expect)
}

func TestSample2(t *testing.T) {
	n, x, y := 6, 3, 4
	expect := []int{2, 6}
	runSample(t, n, x, y, expect)
}
