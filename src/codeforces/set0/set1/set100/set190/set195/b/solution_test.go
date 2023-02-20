package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, m int, expect []int) {
	res := solve(n, m)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 3
	expect := []int{2, 1, 3, 2}
	runSample(t, n, m, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 1
	expect := []int{1, 1, 1}
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 2, 6
	expect := []int{3, 4}
	runSample(t, n, m, expect)
}
