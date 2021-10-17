package lcp39

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m := len(source)
	n := len(source[0])

	cnt := make(map[int]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt[source[i][j]]++
			cnt[target[i][j]]--
		}
	}

	var res int
	for _, v := range cnt {
		if v > 0 {
			res += v
		}
	}

	return res
}
