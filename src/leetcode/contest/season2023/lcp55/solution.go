package lcp55

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	var res int

	for _, cur := range fruits {
		i, num := cur[0], cur[1]
		j := (num + limit - 1) / limit
		res += time[i] * j
	}

	return res
}
