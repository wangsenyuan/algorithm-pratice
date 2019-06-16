package p889

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, R int, C int, r0 int, c0 int, expect [][]int) {
	res := spiralMatrixIII(R, C, r0, c0)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %d, expect %v, but got %v", R, C, r0, c0, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 1, 4
	r0, c0 := 0, 0
	expect := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	runSample(t, R, C, r0, c0, expect)
}

func TestSample2(t *testing.T) {
	R, C := 5, 6
	r0, c0 := 1, 4
	expect := [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}

	runSample(t, R, C, r0, c0, expect)
}
