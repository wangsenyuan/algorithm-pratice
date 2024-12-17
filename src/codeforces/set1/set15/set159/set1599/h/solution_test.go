package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, rect []int) {
	ask := func(x int, y int) int {
		var res int
		if x <= rect[0] {
			res += rect[0] - x
		} else if x >= rect[2] {
			res += x - rect[2]
		}
		if y <= rect[1] {
			res += rect[1] - y
		} else if y >= rect[3] {
			res += y - rect[3]
		}
		return res
	}

	res := solve(ask)

	if !reflect.DeepEqual(res, rect) {
		t.Fatalf("Sample result %v, not correct, expecting %v", res, rect)
	}
}

func TestSample1(t *testing.T) {
	rect := []int{2, 3, 4, 5}
	runSample(t, rect)
}

func TestSample2(t *testing.T) {
	rect := []int{3, 3, n - 10, n - 100}
	runSample(t, rect)
}
