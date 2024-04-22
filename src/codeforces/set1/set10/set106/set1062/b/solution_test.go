package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 20
	expect := []int{10, 2}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 5184
	expect := []int{6, 4}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 1000000
	expect := []int{10, 4}
	runSample(t, n, expect)
}
