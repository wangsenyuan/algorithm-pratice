package p1222

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queens [][]int, king []int, expect [][]int) {
	res := queensAttacktheKing(queens, king)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", queens, king, expect, res)
	}
}

func TestSample1(t *testing.T) {
	queens := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
	king := []int{0, 0}
	expect := [][]int{{0, 1}, {1, 0}, {3, 3}}
	runSample(t, queens, king, expect)
}

func TestSample2(t *testing.T) {
	queens := [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}
	king := []int{3, 3}
	expect := [][]int{{2, 2}, {3, 4}, {4, 4}}
	runSample(t, queens, king, expect)
}

func TestSample3(t *testing.T) {
	queens := [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}}
	king := []int{3, 4}
	expect := [][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}}
	runSample(t, queens, king, expect)
}
