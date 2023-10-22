package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, workers []string, holidays []int, projects [][]int, expect []int) {
	res := solve1(workers, holidays, projects)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	workers := []string{
		"Saturday Sunday",
		"Tuesday Thursday",
		"Monday Wednesday Friday Saturday",
	}
	holidays := []int{4, 7, 13, 14, 15}
	projects := [][]int{
		{1, 1, 3, 3, 2},
		{2, 3, 2},
		{3, 3, 3, 1, 1},
		{3, 3, 3, 3, 3, 3, 3, 3},
	}
	expect := []int{25, 9, 27, 27}
	runSample(t, workers, holidays, projects, expect)
}
