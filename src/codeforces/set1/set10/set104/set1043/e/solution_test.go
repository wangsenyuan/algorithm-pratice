package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, people [][]int, avoids [][]int, expect []int) {
	res := solve(people, avoids)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	people := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	avoids := [][]int{
		{1, 2},
		{2, 3},
	}

	expect := []int{3, 0, 3}

	runSample(t, people, avoids, expect)
}

func TestSample2(t *testing.T) {
	people := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}
	avoids := [][]int{
		{1, 2},
		{2, 3},
		{1, 3},
	}

	expect := []int{0, 0, 0}

	runSample(t, people, avoids, expect)
}

func TestSample3(t *testing.T) {
	people := [][]int{
		{-1, 3},
		{2, 4},
		{1, 1},
		{3, 5},
		{2, 2},
	}
	avoids := [][]int{
		{1, 4},
		{2, 3},
		{3, 5},
	}

	expect := []int{4, 14, 4, 16, 10}

	runSample(t, people, avoids, expect)
}
