package p1203

import "testing"

func checkValid(res []int, beforeItems [][]int) bool {
	set := make(map[int]bool)

	for i := 0; i < len(res); i++ {
		set[res[i]] = true

		before := beforeItems[res[i]]

		for j := 0; j < len(before); j++ {
			if !set[before[j]] {
				return false
			}
		}
	}

	return len(set) == len(res)
}

func runSample(t *testing.T, n, m int, group []int, beforeItems [][]int, noResult bool) {
	res := sortItems(n, m, group, beforeItems)

	if noResult && len(res) > 0 {
		t.Errorf("sample %d %d %v %v, result %v, not valid", n, m, group, beforeItems, res)
	}
	if !noResult && !checkValid(res, beforeItems) {
		t.Errorf("sample %d %d %v %v, result %v, not valid", n, m, group, beforeItems, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 8, 2
	groups := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}
	runSample(t, n, m, groups, beforeItems, false)
}

func TestSample2(t *testing.T) {
	n, m := 8, 2
	groups := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	beforeItems := [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}
	runSample(t, n, m, groups, beforeItems, true)
}
