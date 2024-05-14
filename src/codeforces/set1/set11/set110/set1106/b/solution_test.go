package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, a []int, c []int, customers [][]int, expect []int) {
	res := solve(a, c, customers)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{8, 6, 2, 1, 4, 5, 7, 5}
	c := []int{6, 3, 3, 2, 6, 2, 3, 2}
	customers := [][]int{
		{2, 8},
		{1, 4},
		{4, 7},
		{3, 4},
		{6, 10},
	}
	expect := []int{22, 24, 14, 10, 39}
	runSample(t, a, c, customers, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 6, 6, 6, 6, 6}
	c := []int{6, 66, 666, 6666, 66666, 666666}
	customers := [][]int{
		{1, 6},
		{2, 6},
		{3, 6},
		{4, 6},
		{5, 6},
		{6, 66},
	}
	expect := []int{36, 396, 3996, 39996, 399996, 0}
	runSample(t, a, c, customers, expect)
}
