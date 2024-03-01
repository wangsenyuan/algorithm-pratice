package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{20, 5, 1, 4, 2}
	expect := []int{4, 3, 0, 3, 1}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1434, 7, 1442}
	expect := []int{1, 0, 2}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1}
	expect := []int{0}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{999999999, 999999999, 999999999, 1000000000, 1000000000}
	expect := []int{4, 4, 4, 4, 4}
	runSample(t, a, expect)
}
