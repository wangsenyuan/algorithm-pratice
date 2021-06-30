package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n, k int, expect []int) {
	res := solve(n, k)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	expect := []int{6, 12}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	expect := []int{3, 3}
	runSample(t, n, k, expect)
}

func TestSample3(t *testing.T) {
	n, k := 2, 3
	expect := []int{1, 1, 1}
	runSample(t, n, k, expect)
}

func TestSample4(t *testing.T) {
	expect := nCr(3, 3)
	res := big_nCr(3, 3)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}
