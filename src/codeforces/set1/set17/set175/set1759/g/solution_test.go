package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, b []int, expect []int) {
	res := solve(b)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	b := []int{4, 3, 6}
	expect := []int{1, 4, 2, 3, 5, 6}
	runSample(t, b, expect)
}

func TestSample2(t *testing.T) {
	b := []int{2, 4}
	expect := []int{1, 2, 3, 4}
	runSample(t, b, expect)
}

func TestSample3(t *testing.T) {
	b := []int{8, 7, 2, 3}
	//expect := []int{1, 2, 3, 4}
	runSample(t, b, nil)
}

func TestSample4(t *testing.T) {
	b := []int{6, 4, 2}
	expect := []int{5, 6, 3, 4, 1, 2}
	runSample(t, b, expect)
}
