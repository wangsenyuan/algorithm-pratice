package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x int, n int, m int, expect []int) {
	res := solve(x, n, m)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, n, m := 12, 1, 2
	expect := []int{1, 2}
	runSample(t, x, n, m, expect)
}

func TestSample2(t *testing.T) {
	x, n, m := 12, 1, 1
	expect := []int{3, 3}
	runSample(t, x, n, m, expect)
}

func TestSample3(t *testing.T) {
	x, n, m := 12, 0, 0
	expect := []int{12, 12}
	runSample(t, x, n, m, expect)
}

func TestSample4(t *testing.T) {
	x, n, m := 706636307, 0, 3
	expect := []int{88329539, 88329539}
	runSample(t, x, n, m, expect)
}

func TestSample5(t *testing.T) {
	x, n, m := 12, 1000000000, 1000000000
	expect := []int{0, 0}
	runSample(t, x, n, m, expect)
}
