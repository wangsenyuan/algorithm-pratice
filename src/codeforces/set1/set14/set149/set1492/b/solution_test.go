package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, P []int, expect []int) {
	res := solve(P)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	P := []int{1, 2, 3, 4}
	expect := []int{4, 3, 2, 1}
	runSample(t, P, expect)
}

func TestSample2(t *testing.T) {
	P := []int{1, 5, 2, 4, 3}
	expect := []int{5, 2, 4, 3, 1}
	runSample(t, P, expect)
}

func TestSample3(t *testing.T) {
	P := []int{4, 2, 5, 3, 6, 1}
	expect := []int{6, 1, 5, 3, 4, 2}
	runSample(t, P, expect)
}

func TestSample4(t *testing.T) {
	P := []int{1}
	expect := []int{1}
	runSample(t, P, expect)
}
