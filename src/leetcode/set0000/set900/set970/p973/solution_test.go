package p973

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, points [][]int, K int, expect [][]int) {
	res := kClosest(points, K)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %d, expect %v, but got %v", points, K, expect, res)
	}
}

func TestSample1(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	K := 1
	expect := [][]int{{-2, 2}}
	runSample(t, points, K, expect)
}

func TestSample2(t *testing.T) {
	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	K := 2
	expect := [][]int{{3, 3}, {-2, 4}}
	runSample(t, points, K, expect)
}
