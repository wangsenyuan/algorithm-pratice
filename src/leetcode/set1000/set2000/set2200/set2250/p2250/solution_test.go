package p2250

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, recs [][]int, points [][]int, expect []int) {
	res := countRectangles(recs, points)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	recs := [][]int{{1, 1}, {2, 2}, {3, 3}}
	points := [][]int{{1, 3}, {1, 1}}
	expect := []int{1, 3}
	runSample(t, recs, points, expect)
}

func TestSample2(t *testing.T) {
	recs := [][]int{{1, 2}, {2, 3}, {2, 5}}
	points := [][]int{{2, 1}, {1, 4}}
	expect := []int{2, 1}
	runSample(t, recs, points, expect)
}

func TestSample3(t *testing.T) {
	recs := [][]int{{7, 1}, {2, 6}, {1, 4}, {5, 2}, {10, 3}, {2, 4}, {5, 9}}
	points := [][]int{{10, 3}, {8, 10}, {2, 3}, {5, 4}, {8, 5}, {7, 10}, {6, 6}, {3, 6}}
	expect := []int{1, 0, 4, 1, 0, 0, 0, 1}
	runSample(t, recs, points, expect)
}
