package p822

import "testing"

func runSample(t *testing.T, fronts []int, backs []int, expect int) {
	res := flipgame(fronts, backs)
	if res != expect {
		t.Errorf("sample %v %v, expect %d, but got %d", fronts, backs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	fronts := []int{1, 2, 4, 4, 7}
	backs := []int{1, 3, 4, 1, 3}
	runSample(t, fronts, backs, 2)
}

func TestSample2(t *testing.T) {
	fronts := []int{1, 2}
	backs := []int{1, 1}
	runSample(t, fronts, backs, 2)
}
