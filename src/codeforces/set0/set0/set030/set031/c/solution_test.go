package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, lessons [][]int, expect []int) {
	res := solve(lessons)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	lessons := [][]int{
		{3, 10},
		{20, 30},
		{1, 3},
	}
	expect := []int{1, 2, 3}
	runSample(t, lessons, expect)
}

func TestSample2(t *testing.T) {
	lessons := [][]int{
		{3, 10},
		{20, 30},
		{1, 3},
		{1, 39},
	}
	expect := []int{4}
	runSample(t, lessons, expect)
}

