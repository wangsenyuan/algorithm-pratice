package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, X int, Y int, expect []int) {
	res := solve(A, X, Y)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{9}
	X := 6
	Y := 99
	expect := []int{15}
	runSample(t, A, X, Y, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 3}
	X := 7
	Y := 44
	expect := []int{3, 5, 6}
	runSample(t, A, X, Y, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 10, 15, 20, 25}
	X := 20
	Y := 6
	expect := []int{5, 20, 25, 27, 30}
	runSample(t, A, X, Y, expect)
}
