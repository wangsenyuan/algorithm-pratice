package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, c []int, expect []int) {
	res := solve(c)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 4, 2, 4}
	expect := []int{1, 1, 0, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 3, 4, 2, 3, 2, 7}
	expect := []int{0, 1, 1, 0, 0, 0, 1}
	// 0, 1, 1, 0, 0, 0, 1
	// 0, 1, 1, 0, 0, 0, 1
	// 0, 1, 1, 0, 0, 0, 1
	// 0, 0, 1, 1, 0, 0, 1
	// 0, 0, 0, 1, 1, 0, 1
	// 0, 0, 0, 0, 1, 1, 1
	// 0, 0, 0, 0, 1, 1, 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 0, 0, 4}
	expect := []int{0, 0, 0, 1}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2, 3}
	expect := []int{1, 0, 1}
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{2, 2}
	expect := []int{1, 1}
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 4, 3, 6, 5, 6, 8, 16, 17, 9, 20, 10, 21, 21, 7, 21, 5, 4, 3, 21, 1}
	expect := []int{1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0}
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{1, 3, 5, 6, 3, 3, 3, 2, 1}
	expect := []int{1, 0, 1, 1, 0, 0, 0, 0, 0}
	runSample(t, a, expect)
}
