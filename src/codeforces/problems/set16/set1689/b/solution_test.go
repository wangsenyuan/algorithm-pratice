package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, expect []int) {
	res := solve(len(P), P)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 2, 3}
	expect := []int{2, 3, 1}
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{2, 3, 4, 5, 1}
	expect := []int{1, 2, 3, 4, 5}
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{2, 3, 1, 4}
	expect := []int{1, 2, 4, 3}
	runSample(t, P, expect)
}
