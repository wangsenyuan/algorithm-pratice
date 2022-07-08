package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, m int, maps [][]string, expect []int) {
	res := solve(m, maps)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m := 3
	maps := [][]string{
		{
			"1 L 2",
			"1 R 3",
			"3 L 4",
			"3 R 5",
		},
		{
			"1 R 2",
			"1 L 3",
			"2 L 4",
		},
		{
			"1 L 2",
			"1 R 3",
		},
	}
	expect := []int{2, 2, 1}
	runSample(t, m, maps, expect)
}

func TestSample2(t *testing.T) {
	m := 2
	maps := [][]string{
		{
			"1 L 2",
			"2 L 3",
			"3 L 4",
			"4 L 5",
		},
		{
			"1 L 2",
			"2 R 3",
			"3 L 4",
			"4 R 5",
			"5 L 6",
		},
	}
	expect := []int{5, 1}
	runSample(t, m, maps, expect)
}
