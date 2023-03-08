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
	A := []int{2, 3}
	var expect []int
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 5}
	var expect []int = []int{5, 4}
	runSample(t, A, expect)
}
func TestSample3(t *testing.T) {
	A := []int{2, 4, 6, 9}
	var expect []int = []int{2, 1}
	runSample(t, A, expect)
}
