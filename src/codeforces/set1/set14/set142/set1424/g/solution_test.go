package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, records [][]int, expect []int) {
	res := solve(records)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	records := [][]int{
		{1, 5},
		{2, 4},
		{5, 6},
	}
	expect := []int{2, 2}
	runSample(t, records, expect)
}

func TestSample2(t *testing.T) {
	records := [][]int{
		{3, 4},
		{4, 5},
		{4, 6},
		{8, 10},
	}
	expect := []int{4, 2}
	runSample(t, records, expect)
}
