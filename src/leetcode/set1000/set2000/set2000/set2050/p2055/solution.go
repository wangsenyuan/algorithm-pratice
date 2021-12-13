package p2055

import "sort"

func platesBetweenCandles(s string, queries [][]int) []int {

	type plate struct {
		ind        int
		leftCandle int
	}

	n := len(s)

	type Q struct {
		left, right int
		ind         int
	}

	qs := make([]Q, len(queries))
	for i := 0; i < len(queries); i++ {
		qs[i] = Q{queries[i][0], queries[i][1], i}
	}

	sort.Slice(qs, func(i, j int) bool {
		return qs[i].right < qs[j].right
	})

	ans := make([]int, len(queries))

	var que []plate
	prevPlate := -1

	for i, k := 0, 0; i < n; i++ {
		if s[i] == '|' {
			if prevPlate >= 0 {
				for j := prevPlate + 1; j < i; j++ {
					que = append(que, plate{j, prevPlate})
				}
			}

			prevPlate = i
		}

		for k < len(qs) && qs[k].right == i {
			q := qs[k]
			x := sort.Search(len(que), func(x int) bool {
				return que[x].leftCandle >= q.left
			})
			ans[q.ind] = len(que) - x
			k++
		}

	}

	return ans
}
