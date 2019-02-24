package p1000

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, N int, lamps [][]int, queries [][]int, expect []int) {
	res := gridIllumination(N, lamps, queries)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %v %v, expect %v, but got %v", N, lamps, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N := 5
	lamps := [][]int{{0, 0}, {4, 4}}
	quries := [][]int{{1, 1}, {1, 0}}
	expect := []int{1, 0}
	runSample(t, N, lamps, quries, expect)
}

func TestSample2(t *testing.T) {
	N := 10
	lamps := [][]int{{3, 4}, {6, 6}, {1, 8}, {4, 5}, {8, 7}, {0, 6}, {5, 2}, {1, 9}}
	quries := [][]int{{7, 9}, {2, 8}, {8, 6}, {6, 8}, {2, 8}}
	expect := []int{1, 1, 1, 1, 1}
	runSample(t, N, lamps, quries, expect)
}
