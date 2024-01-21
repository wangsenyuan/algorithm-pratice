package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, x []int, expect []int) {
	res := solve(x)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{1, 4, 3}
	expect := []int{8, 7, 6}
	runSample(t, x, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1, 2, 5, 7, 1}
	expect := []int{16, 15, 18, 24, 16}
	runSample(t, x, expect)
}

func TestSample3(t *testing.T) {
	x := []int{1, 10, 100, 1000}
	expect := []int{1111, 1093, 1093, 2893}
	runSample(t, x, expect)
}
