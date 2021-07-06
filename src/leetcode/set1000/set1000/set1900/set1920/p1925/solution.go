package p1925

const N = 101

func longestCommomSubsequence(arrays [][]int) []int {
	n := len(arrays)

	pos := make([]int, n)
	back := make([]int, n)
	var res []int

	for num := 1; num < N; num++ {
		// num after (including) pos[..]
		ok := true
		for i := 0; i < n; i++ {
			arr := arrays[i]
			l, r := pos[i], len(arr)
			for l < r {
				mid := (l + r) / 2
				if arr[mid] >= num {
					r = mid
				} else {
					l = mid + 1
				}
			}
			if r == len(arr) || arr[r] > num {
				ok = false
				break
			}
			back[i] = r
		}
		if !ok {
			continue
		}
		res = append(res, num)
		copy(pos, back)
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
