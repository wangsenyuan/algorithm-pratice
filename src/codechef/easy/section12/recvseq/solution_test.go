package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int) {
	res := solve(A)

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 3, 10, 7}
	expect := []int{1, 3, 5, 7}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{-10, -5, 0, 5, 10}
	expect := []int{-10, -5, 0, 5, 10}
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 2, 10}
	expect := []int{2, 2, 2, 2}
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 0, 0, -1}
	expect := []int{2, 1, 0, -1}
	runSample(t, A, expect)
}
