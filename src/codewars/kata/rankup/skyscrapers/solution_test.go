package skyscrapers

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, clues []int, expect [][]int) {
	res := SolvePuzzle(clues)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	clues := []int{0, 0, 1, 2,
		0, 2, 0, 0,
		0, 3, 0, 0,
		0, 1, 0, 0}
	expect := [][]int{
		{2, 1, 4, 3},
		{3, 4, 1, 2},
		{4, 2, 3, 1},
		{1, 3, 2, 4},
	}
	runSample(t, clues, expect)
}
