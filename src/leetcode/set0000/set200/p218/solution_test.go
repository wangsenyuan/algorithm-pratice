package p218

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, buildings [][]int, expect [][]int) {
	res := getSkyline(buildings)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	buildings := [][]int{
		{0, 2, 3},
		{2, 5, 3},
	}
	expect := [][]int{
		{0, 3},
		{5, 0},
	}
	runSample(t, buildings, expect)
}
