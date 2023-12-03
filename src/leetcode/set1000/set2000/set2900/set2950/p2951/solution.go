package p2951

import "sort"

func countCompleteSubstrings(word string, k int) int {
	n := len(word)
	pos := make([]int, 26)
	id := make([]int, 26)
	for i := 0; i < 26; i++ {
		pos[i] = -1
		id[i] = i
	}

	var res int
	// L 是最靠左的满足条件相邻字符 <= 2的位置
	L := 0
	cnt := make([][]int, n+1)
	cnt[0] = make([]int, 26)
	for i := 0; i < n; i++ {
		if i > 0 && distance(word[i], word[i-1]) > 2 {
			L = i
		}
		cnt[i+1] = make([]int, 26)
		copy(cnt[i+1], cnt[i])
		x := int(word[i] - 'a')
		cnt[i+1][x]++
		pos[x] = i
		sort.Slice(id, func(u, v int) bool {
			return pos[id[u]] > pos[id[v]]
		})

		for j := 0; j < 26 && pos[id[j]] >= L; j++ {
			// j + 1 is cnt
			ln := (j + 1) * k
			if ln > i-L+1 {
				break
			}
			y := id[j]
			if cnt[i+1][y]-cnt[i+1-ln][y] > k {
				// it has too many
				break
			}
			// 因为前面的计数都 <= k, 总数 = k * (j + 1), 所以，肯定在某个位置可以得到
			if cnt[i+1][y]-cnt[i+1-ln][y] == k && (j == 25 || pos[id[j+1]] < i+1-ln) {
				res++
			}
		}
	}

	return res
}

func distance(a, b byte) int {
	x := int(a - 'a')
	y := int(b - 'a')
	return abs(x - y)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
