package p850

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, richer [][]int, quiet []int, expect []int) {
	res := loudAndRich(richer, quiet)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", richer, quiet, expect, res)
	}
}

func TestSample1(t *testing.T) {
	richer := [][]int{
		{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3},
	}

	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}

	expect := []int{5, 5, 2, 5, 4, 5, 6, 7}

	runSample(t, richer, quiet, expect)
}
