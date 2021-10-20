package lcp45

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, position []int, terrain [][]int, obstacle [][]int, expect [][]int) {
	res := bicycleYard(position, terrain, obstacle)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	position := []int{0, 0}
	terrain := [][]int{{0, 0}, {0, 0}}
	obstacle := [][]int{{0, 0}, {0, 0}}
	expect := [][]int{{0, 1}, {1, 0}, {1, 1}}
	runSample(t, position, terrain, obstacle, expect)
}

func TestSample2(t *testing.T) {
	position := []int{1, 1}
	terrain := [][]int{{5, 0}, {0, 6}}
	obstacle := [][]int{{0, 6}, {7, 0}}
	expect := [][]int{{0, 1}}
	runSample(t, position, terrain, obstacle, expect)
}
