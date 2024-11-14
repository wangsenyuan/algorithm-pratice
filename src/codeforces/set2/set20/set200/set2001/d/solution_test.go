package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 1, 3}
	expect := []int{3, 2, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := []int{1}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 1, 3, 2, 1, 3, 2, 1}
	expect := []int{3, 1, 2}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{1, 2}
	expect := []int{1, 2}
	runSample(t, a, expect)
}

func TestSample5(t *testing.T) {
	a := []int{5, 2, 1, 7, 9, 7, 2, 5, 5, 2}
	expect := []int{5, 1, 9, 7, 2}
	runSample(t, a, expect)
}

func TestSample6(t *testing.T) {
	a := []int{2, 2, 8, 7, 7, 9, 8, 1, 9, 6}
	expect := []int{2, 7, 9, 8, 1, 6}
	runSample(t, a, expect)
}

func TestSample7(t *testing.T) {
	a := []int{4, 1, 4, 5, 4, 5, 10, 1, 5, 1}
	expect := []int{5, 4, 10, 1}
	runSample(t, a, expect)
}
