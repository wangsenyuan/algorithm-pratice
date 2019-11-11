package p1253

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, upper, lower int, cols []int, expect [][]int) {
	res := reconstructMatrix(upper, lower, cols)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %v, expect %v, but got %v", upper, lower, cols, expect, res)
	}
}

func TestSample1(t *testing.T) {
	upper := 1
	lower := 4
	cols := []int{2, 1, 2, 0, 0, 2}

	runSample(t, upper, lower, cols, nil)
}
