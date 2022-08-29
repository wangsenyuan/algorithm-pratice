package p2392

import "testing"

func runSample(t *testing.T, k int, rowConditions [][]int, colConditions [][]int, expect bool) {
	res := buildMatrix(k, rowConditions, colConditions)

	if expect != (len(res) > 0) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	row := make([]int, k+1)
	col := make([]int, k+1)
	cnt := make([]int, k+1)
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			if res[i][j] > 0 {
				x := res[i][j]
				cnt[x]++
				row[res[i][j]] = i
				col[res[i][j]] = j
			}
		}
	}

	for i := 1; i <= k; i++ {
		if cnt[i] != 1 {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	for _, con := range rowConditions {
		a, b := con[0], con[1]
		if row[a] >= row[b] {
			t.Fatalf("Sample result %v, not correct, expect %d above %d, but not", res, a, b)
		}
	}

	for _, con := range colConditions {
		l, r := con[0], con[1]
		if col[l] >= col[r] {
			t.Fatalf("Sample result %v, not correct, expect %d left %d, but not", res, l, r)
		}
	}
}

func TestSample1(t *testing.T) {
	k := 3
	row := [][]int{
		{1, 2}, {3, 2},
	}
	col := [][]int{
		{2, 1}, {3, 2},
	}
	expect := true
	runSample(t, k, row, col, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	row := [][]int{
		{1, 2}, {2, 3}, {3, 1}, {2, 3},
	}
	col := [][]int{
		{2, 1},
	}
	expect := false
	runSample(t, k, row, col, expect)
}

func TestSample3(t *testing.T) {
	k := 8
	row := [][]int{
		{1, 2}, {7, 3}, {4, 3}, {5, 8}, {7, 8}, {8, 2}, {5, 8}, {3, 2}, {1, 3}, {7, 6}, {4, 3}, {7, 4}, {4, 8}, {7, 3}, {7, 5},
	}
	col := [][]int{
		{5, 7}, {2, 7}, {4, 3}, {6, 7}, {4, 3}, {2, 3}, {6, 2},
	}
	expect := true
	runSample(t, k, row, col, expect)
}
