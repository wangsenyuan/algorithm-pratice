package p1623

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, tows [][]int, radius int, expect []int) {
	res := bestCoordinate(tows, radius)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	tows := [][]int{
		{1, 2, 5}, {2, 1, 7}, {3, 1, 9},
	}
	radius := 2
	expect := []int{2, 1}
	runSample(t, tows, radius, expect)
}

func TestSample2(t *testing.T) {
	tows := [][]int{
		{1, 2, 13}, {2, 1, 7}, {0, 1, 9},
	}
	radius := 2
	expect := []int{1, 2}
	runSample(t, tows, radius, expect)
}

func TestSample3(t *testing.T) {
	tows := [][]int{
		{28, 6, 30}, {23, 16, 0}, {21, 42, 22}, {50, 33, 34}, {14, 7, 50}, {40, 31, 4}, {39, 45, 17}, {46, 21, 12}, {45, 36, 45}, {35, 43, 43}, {29, 41, 48}, {22, 27, 5}, {42, 44, 45}, {10, 49, 50}, {47, 43, 26}, {40, 36, 25}, {10, 25, 6}, {27, 30, 30}, {50, 35, 20}, {11, 0, 44}, {34, 29, 28},
	}
	radius := 12
	expect := []int{42, 44}
	runSample(t, tows, radius, expect)
}
