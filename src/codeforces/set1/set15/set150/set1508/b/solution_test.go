package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int64, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	var k int64 = 3
	expect := []int{2, 1, 3}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	var k int64 = 5
	expect := []int{1, 2, 4, 3, 5, 6}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	var k int64 = 4
	expect := []int{3, 2, 1}
	runSample(t, n, k, expect)
}
