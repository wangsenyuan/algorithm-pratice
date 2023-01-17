package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 5, 5, 5}
	expect := []int{5, 5, 5, 5}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 2, 3, 1}
	expect := []int{4, 2, 1, 1}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{6, 10, 15}
	expect := []int{15, 5, 1}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 6, 3, 9, 24}
	expect := []int{24, 3, 3, 3, 1}
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{3, 11, 7, 2, 5}
	expect := []int{11, 1, 1, 1, 1}
	runSample(t, A, expect)
}
