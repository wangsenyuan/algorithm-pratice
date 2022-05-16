package p2271

import "sort"

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	// always start from the begining

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	n := len(tiles)

	all := make([]int, n)

	for i := 0; i < n; i++ {
		all[i] = tiles[i][1] - tiles[i][0] + 1
		if i > 0 {
			all[i] += all[i-1]
		}
	}

	var best int
	var j int
	// j will only increase or stay
	for i := 0; i < n; i++ {
		cur := tiles[i][0]
		for j+1 < n && tiles[j+1][0]-cur < carpetLen {
			j++
		}
		var tmp int
		if j > 0 {
			tmp = all[j-1]
		}
		if i > 0 {
			tmp -= all[i-1]
		}
		tmp += min(tiles[i][0]+carpetLen-1, tiles[j][1]) - tiles[j][0] + 1
		best = max(best, tmp)
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
