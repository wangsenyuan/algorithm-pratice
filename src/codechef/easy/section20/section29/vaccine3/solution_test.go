package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, G int, P int, X []int, expect []int) {
	res := solve(G, P, X)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	G, P := 5, 2
	X := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	expect := []int{6, 6}
	runSample(t, G, P, X, expect)
}

func TestSample2(t *testing.T) {
	G, P := 5, 2
	X := []int{2, 2, 2, 2, 2, 3, 2, 2, 2, 2}
	expect := []int{6, 7}
	runSample(t, G, P, X, expect)
}

func TestSample3(t *testing.T) {
	G, P := 9, 4
	X := []int{2, 2, 2, 2, 3, 2, 2, 2, 2, 2}
	expect := []int{1, 1}
	runSample(t, G, P, X, expect)
}
