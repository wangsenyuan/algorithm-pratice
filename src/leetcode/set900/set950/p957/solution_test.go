package p957

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, cells []int, N int, expect []int) {
	res := prisonAfterNDays(cells, N)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", cells, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cells := []int{0, 1, 0, 1, 1, 0, 0, 1}
	N := 7
	expect := []int{0, 0, 1, 1, 0, 0, 0, 0}
	runSample(t, cells, N, expect)
}

func TestSample2(t *testing.T) {
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
	N := 1000000000
	expect := []int{0, 0, 1, 1, 1, 1, 1, 0}
	runSample(t, cells, N, expect)
}
