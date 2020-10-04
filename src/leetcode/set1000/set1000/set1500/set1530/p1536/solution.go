package p1536

func minSwaps(grid [][]int) int {
	n := len(grid)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				pos[i] = j
			}
		}
	}

	var res int

	for i := 0; i < n; i++ {
		if pos[i] <= i {
			continue
		}
		j := i + 1
		for j < n && pos[j] > i {
			j++
		}
		if j == n {
			return -1
		}
		res += j - i
		for j > i {
			pos[j] = pos[j-1]
			j--
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

//type Pair struct {
//	first, second int
//}
//
//type Pairs []Pair
//
//func (ps Pairs) Len() int {
//	return len(ps)
//}
//
//func (ps Pairs) Less(i, j int) bool {
//	return ps[i].first < ps[j].first
//}
//
//func (ps Pairs) Swap(i, j int) {
//	ps[i], ps[j] = ps[j], ps[i]
//}
