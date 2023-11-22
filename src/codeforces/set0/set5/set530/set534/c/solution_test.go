package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A int, d []int, expect []int) {
	res := solve(A, d)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := 8
	d := []int{4, 4}
	expect := []int{3, 3}
	runSample(t, A, d, expect)
}

func TestSample2(t *testing.T) {
	A := 3
	d := []int{5}
	expect := []int{4}
	runSample(t, A, d, expect)
}

func TestSample3(t *testing.T) {
	A := 3
	d := []int{2, 3}
	expect := []int{0, 1}
	runSample(t, A, d, expect)
}
