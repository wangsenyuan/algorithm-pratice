package main

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, a []int, expect []int) {
	res := solve(a)

	sort.Ints(expect)
	sort.Ints(res)

	if !reflect.DeepEqual(expect, res) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 2, 2}
	expect := []int{-3, -3, -3, -3}
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0}
	expect := []int{0}
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-3, -3, 2}
	expect := []int{-3, -3, 2}
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-3, -3, -4}
	expect := []int{-3, -3, 3}
	// -3, 2, 4 = 24
	runSample(t, a, expect)
}
